package main

import (
	"context"
	cfg "github.com/nj-eka/gobe1/chat/chatcli/config"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()
	d := net.Dialer{
		Timeout:   cfg.Timeout,
		KeepAlive: cfg.KeepAlive,
	}
	conn, err := d.DialContext(ctx, cfg.Protocol, cfg.Address)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("client [pid: %d] connected to chat server: %s -> %s ", os.Getpid(), conn.LocalAddr(), conn.RemoteAddr())

	go func() {
		// if conn.Closed => cancel context
		if _, err := io.Copy(os.Stdout, conn); err != nil{
			cancel()
		}
	}()
	go func() {
		// doc: Copy copies from src to dst until either EOF is reached on src or an error occurs.
		// os.Stdin never ends with EOF here ... so it will be closed with process exit
		_, _ = io.Copy(conn, os.Stdin)
	}()

	<-ctx.Done()
	if err := conn.Close(); err != nil {
		log.Println("closing conn failed: ", err)
	} else {
		log.Println("conn closed")
	}
}
