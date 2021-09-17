package main

import (
	"context"
	"github.com/nj-eka/gobe1/httpfileserver/downloader"
	"github.com/nj-eka/gobe1/httpfileserver/uploader"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"time"
)

const ( // -> config
	logLevel         = logrus.DebugLevel
	rootDir          = "uploads"
	uploadAddr       = "localhost:8008"
	uploadEntry		 = "/upload"
	downloadAddr     = "localhost:8009"
	operationTimeout = 10 * time.Second
	maxUploadSize	= int64(10 << 24)
)

var tempDir = os.TempDir()

func init(){
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
	if err := os.MkdirAll(rootDir, os.ModePerm); err != nil{
		logrus.Fatalf("making root dir [%s] failed: %v", rootDir, err)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go uploader.Run(ctx, &wg, uploadAddr, uploadEntry, rootDir, tempDir, downloadAddr, operationTimeout, maxUploadSize)
	go downloader.Run(ctx, &wg, downloadAddr, rootDir, operationTimeout)
	wg.Wait()
}
