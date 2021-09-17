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
	defer cancel()
	d := net.Dialer{
		Timeout:   cfg.Timeout,
		KeepAlive: cfg.KeepAlive,
	}
	conn, err := d.DialContext(ctx, cfg.Protocol, cfg.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("client [pid: %d] connected to chat server: %s -> %s ", os.Getpid(), conn.LocalAddr(), conn.RemoteAddr())
	go func() {
		// if conn.Closed => Copy exit with error - ok
		io.Copy(os.Stdout, conn)
	}()
	select {
	case <-ctx.Done():
		if err := conn.Close(); err != nil {
			log.Println("closing conn failed: ", err)
		} else {
			log.Println("conn closed")
		}
	case <-func() <-chan struct{} {
		done := make(chan struct{})
		go func() {
			// doc: Copy copies from src to dst until either EOF is reached on src or an error occurs.
			// os.Stdin never ends with EOF here
			_, _ = io.Copy(conn, os.Stdin)
			close(done)
		}()
		return done
		}():
	}
}
