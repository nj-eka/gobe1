package uploader

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

type UploadHandler struct {
	hostAddr string
	rootDir string
	tempDir string
	maxUploadSize int64
	downloadAddr string
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006/01/02 15:04:05"))
	return []byte(stamp), nil
}

func (t JSONTime) Less(than JSONTime) bool{
	return time.Time(t).After(time.Time(than))
}

type FileInfo struct {
	RelPath  string   `json:"rpath"`
	Size     int64    `json:"size"`
	ModTime  JSONTime `json:"mdf"`
	MimeType string   `json:"mime"`
}

type FileInfos []FileInfo

func (fis FileInfos) SortByModTime(){
	sort.SliceStable(fis, func(i,j int) bool{
		return fis[i].ModTime.Less(fis[j].ModTime)
	})
}

func (h *UploadHandler) WalkDir(pattern string) (result FileInfos, err error) {
	if pattern == "" {
		pattern = "*"
	}
	err = filepath.WalkDir(h.rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			logrus.Errorf("walking %s - %v err: %v", path, d, err)
		}
		if !d.IsDir() {
			matched, err := filepath.Match(pattern, d.Name())
			if matched {
				if fi, err := d.Info(); err == nil {
					mts := "unknown"
					if mt, err := mimetype.DetectFile(path); err == nil {
						mts = mt.String()
					}
					result = append(
						result,
						FileInfo{
							strings.TrimPrefix(path, h.rootDir),
							fi.Size(),
							JSONTime(fi.ModTime()),
							mts,
						},
					)
					result.SortByModTime()
				} else {
					logrus.Errorf("file stat %s err: %v", path, err)
				}
			}
			return err // nil || ErrBadPattern
		}
		return nil
	})
	if err != nil {
		logrus.Errorf("walking dir [%s] with pattern [%s] failed: %v", h.rootDir, pattern, err)
	}
	return
}

func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pattern := r.FormValue("pattern")
		if foundFiles , err := h.WalkDir(pattern); err == nil {
			if err := json.NewEncoder(w).Encode(foundFiles); err != nil {
				http.Error(w, "bad", http.StatusInternalServerError)
				logrus.Error(err)
			}
		} else {
			http.Error(w, "bad pattern", http.StatusBadRequest)
		}
	case http.MethodPost:
		// https://freshman.tech/file-upload-golang/
		r.Body = http.MaxBytesReader(w, r.Body, h.maxUploadSize)
		defer r.Body.Close()
		if err := r.ParseMultipartForm(h.maxUploadSize); err != nil {
			http.Error(w, "The uploaded file is too big. Please choose an file that's less than 16MB in size", http.StatusBadRequest)
			return
		}
		uploadFile, uploadFileHeader, err := r.FormFile("file")
		logrus.Debugf(
			"uploading file with name [%s] size [%d] headers [%v]: %v",
			uploadFileHeader.Filename,
			uploadFileHeader.Size,
			uploadFileHeader.Header,
			err,
		)
		if err != nil {
			http.Error(w, "bad file", http.StatusBadRequest)
			return
		}
		defer func() {
			if err := uploadFile.Close(); err != nil {
				logrus.Errorf("closing uploaded file [%s - %d] failed: %v", uploadFileHeader.Filename, uploadFileHeader.Size, err)
			}
		}()
		tempFile, err := os.CreateTemp(h.tempDir, fmt.Sprintf("%s_*", filepath.Base(uploadFileHeader.Filename)))
		if err != nil {
			http.Error(w, "bad operation", http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
		defer func(){
			if tempFile != nil {
				if err := os.Remove(tempFile.Name()); err != nil{
					logrus.Errorf("removing temp file [%s] for [%s - %d] failed: %v", tempFile.Name(), uploadFileHeader.Filename, uploadFileHeader.Size, err)
				}
			}
		}()
		if written, err := io.Copy(tempFile, uploadFile); err != nil{
			http.Error(w, "bad file", http.StatusBadRequest)
			logrus.Errorf("copy file [%d/%d] failed: %v", written, uploadFileHeader.Size, err)
			return
		}
		fileName := filepath.Base(uploadFileHeader.Filename)  // uploadFileHeader.Filename
		ext := filepath.Ext(fileName)
		filePath := filepath.Join(h.rootDir, fmt.Sprintf("%s_%d%s", strings.TrimSuffix(fileName, ext), time.Now().UnixNano(), ext))
		if err := os.Rename(tempFile.Name(), filePath); err != nil{
			// add duplicate handling if needed
			http.Error(w, "bad operation", http.StatusInternalServerError)
			logrus.Errorf("renaming uploaded file [%s -> %s] failed: %v", tempFile.Name(), filePath, err)
			return
		} else {
			tempFile = nil
			logrus.Infof("uploaded file [%s -> %s] saved - ok", uploadFileHeader.Filename, filePath)
		}
		_, _ = fmt.Fprintf(w, "File [%s] is uploaded and available on link: %s/%s\n ", uploadFileHeader.Filename, h.downloadAddr, strings.TrimPrefix(filePath, h.rootDir))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func Run(ctx context.Context, wg *sync.WaitGroup, serveAddr, serveEntry, rootDir, tempDir, downloadAddr string, operationTimeout time.Duration, maxUploadSize int64) {
	defer wg.Done()
	tempSessionDir, err := os.MkdirTemp(tempDir, "uploads_*");
	if err != nil{
		logrus.Errorf("creating temp dir in [%s] failed: %v", os.TempDir(), err)
		tempSessionDir = tempDir
	}
	logrus.Infoln("Uploader session temp dir:", tempSessionDir)
	defer func(){
		if err := os.RemoveAll(tempSessionDir); err != nil{
			logrus.Errorf("removing temp dir [%s] failed: %v", tempSessionDir, err)
		} else {
			logrus.Infof("removing temp dir [%s] - ok", tempSessionDir)
		}
	}()
	uploadHandler := &UploadHandler{
		hostAddr: serveAddr,
		rootDir:  rootDir,
		tempDir: tempSessionDir,
		maxUploadSize: maxUploadSize,
		downloadAddr: downloadAddr,
	}
	http.Handle(serveEntry, uploadHandler)
	server := &http.Server{
		Addr:         uploadHandler.hostAddr,
		ReadTimeout:  operationTimeout,
		WriteTimeout: operationTimeout,
	}
	go func() {
		select {
		case <-ctx.Done():
			err := server.Close()
			logrus.Infof("uploader closed on [%s/%s]: %v", uploadHandler.hostAddr, serveEntry, err)
		}
	}()
	logrus.Infof("start uploader on [%s%s] with root [%s]", uploadHandler.hostAddr, serveEntry, uploadHandler.rootDir)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error("uploader stopped with err: ", err)
	} else {
		logrus.Info("uploader stopped - ok")
	}
}
