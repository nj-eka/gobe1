package downloader

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

func Run(ctx context.Context, wg *sync.WaitGroup, serveAddr, rootDir string, operationTimeout time.Duration) {
	defer wg.Done()
	server := &http.Server{
		Addr:         serveAddr,
		Handler:      http.FileServer(http.Dir(rootDir)),
		ReadTimeout:  operationTimeout,
		WriteTimeout: operationTimeout,
	}
	go func() {
		select {
		case <-ctx.Done():
			err := server.Close()
			logrus.Infof("downloader closed on [%s] with root [%s]: %v", serveAddr, rootDir, err)
		}
	}()
	logrus.Infof("start downloader on [%s] with root [%s]", serveAddr, rootDir)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Errorf("downloader [%s] failed: %v", serveAddr, err)
	}
}
