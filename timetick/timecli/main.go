package main

import (
	"context"
	cfg "github.com/nj-eka/gobe1/timetick/timecli/config"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	d := net.Dialer{
		Timeout:   cfg.Timeout,
		KeepAlive: cfg.KeepAlive,
	}
	conn, err := d.DialContext(ctx, cfg.Protocol, cfg.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("client [pid: %d] connected to time server: %s -> %s ", os.Getpid(), conn.LocalAddr(), conn.RemoteAddr())
	go func(){ // force conn closing
		<-ctx.Done()
		if err := conn.Close(); err != nil{
			log.Println("closing conn failed: ", err)
		} else {
			log.Println("conn closed")
		}
	}()
	log.Println(io.Copy(os.Stdout, conn))
}
