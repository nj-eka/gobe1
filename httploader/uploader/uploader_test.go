package uploader

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const ( // -> config
	logLevel         = logrus.DebugLevel
	rootSubDir          = "uploads_test"
	uploadAddr       = "localhost:8008"
	uploadEntry		 = "/upload"
	downloadAddr     = "localhost:8009"
	maxUploadSize	= int64(10 << 24)
)

var (
	wd, _ = os.Getwd()
	rootDir = filepath.Join(wd, rootSubDir)
	tempDir = os.TempDir()
	testFileTxtName = "test.txt"
	testFileTxtPath = filepath.Join(rootDir, testFileTxtName)
	testFileTxtContent = []byte("text content")
	testFileTxt *os.File
	testFileInfo FileInfo
)

func init(){
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
	if err := os.MkdirAll(rootDir, os.ModePerm); err != nil{
		logrus.Fatalf("making root dir [%s] failed: %v", rootDir, err)
	}
	var err error
	testFileTxt, err = os.Create(testFileTxtPath)
	if err != nil{
		logrus.Fatalf("creating test file [%s] failed: %v", testFileTxtPath, err)
	}
	if n, err := testFileTxt.Write(testFileTxtContent); err != nil || n != len(testFileTxtContent){
		testFileTxt.Close()
		logrus.Fatalf("writing test file [%s] failed: %v", testFileTxtPath, err)
	}
	if stat, err := testFileTxt.Stat(); err != nil{
		testFileTxt.Close()
		logrus.Fatalf("writing test file [%s] failed: %v", testFileTxtPath, err)
	} else {
		mts := "text/plain; charset=utf-8"
		//if mt := mimetype.Detect(testFileTxtContent); mt == nil {
		//	mts = mt.String()
		//}
		testFileInfo = FileInfo{
			RelPath:  strings.TrimPrefix(testFileTxtPath, rootDir),
			Size:     stat.Size(),
			ModTime:  JSONTime(stat.ModTime()),
			MimeType: mts,
		}
		testFileTxt.Close()
	}
}

func TestUploadHandler_WalkDir(t *testing.T) {
	// defer os.RemoveAll(rootDir)
	type args struct {
		pattern string
	}
	h := &UploadHandler{
		hostAddr:      uploadAddr,
		rootDir:       rootDir,
		tempDir:       tempDir,
		maxUploadSize: maxUploadSize,
		downloadAddr:  downloadAddr,
	}
	tests := []struct {
		name       string
		args       args
		wantResult FileInfos
		wantErr    error
	}{
		{"WalkDir_Pattern_Empty",
			args{""},
			FileInfos{testFileInfo},
			nil,
		},
		{"WalkDir_Pattern_All",
			args{"*"},
			FileInfos{testFileInfo},
			nil,
		},
		{"WalkDir_Pattern_Ext",
			args{"*.txt"},
			FileInfos{testFileInfo},
			nil,
		},
		{"WalkDir_Pattern_Nomatch",
			args{"______"},
			FileInfos{},
			nil,
		},
		{"WalkDir_Pattern_Bad",
			args{"\\"},
			FileInfos{},
			filepath.ErrBadPattern,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := h.WalkDir(tt.args.pattern)
			if (err != nil) && (err != tt.wantErr) {
				t.Errorf("WalkDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (len(gotResult) != len(tt.wantResult)) && len(gotResult) != 0 && !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("WalkDir() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestRun(t *testing.T) {
	defer os.RemoveAll(rootDir)

	file, _ := os.Open(testFileTxtPath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, uploadEntry, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()
	// создаем заглушку файлового сервера. Для прохождения тестов
	// нам достаточно чтобы он возвращал 200 статус
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r	*http.Request) {
		fmt.Fprintln(w, "ok!")
	}))
	defer ts.Close()
	uploadHandler := &UploadHandler{
		hostAddr:      uploadAddr,
		rootDir:       rootDir,
		tempDir:       tempDir,
		maxUploadSize: maxUploadSize,
		downloadAddr:  downloadAddr,
	}
	uploadHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := strings.TrimSuffix(testFileTxtName, filepath.Ext(testFileTxtName))
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
