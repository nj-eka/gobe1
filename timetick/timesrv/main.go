package main

import (
	"context"
	"fmt"
	cfg "github.com/nj-eka/gobe1/timetick/timesrv/config"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	netCfg := net.ListenConfig{
		KeepAlive: cfg.KeepAlive,
	}
	listener, err := netCfg.Listen(ctx, cfg.Protocol, cfg.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TimeTick server [pid: %d] started at %s://%s\n", os.Getpid(), listener.Addr().Network(), listener.Addr())
	wg := &sync.WaitGroup{}
	defer func(){
		wg.Wait()
		if err := listener.Close(); err != nil{
			log.Printf("closing listener [%s] err: %v\n", listener.Addr(), err)
		} else {
			log.Printf("listener [%s] closed\n", listener.Addr())
		}
	}()
	go broadcaster()
	go messageProducer()
	for {
		select {
		case <-ctx.Done():
			return
		case conn, ok := <- accept(listener):
			if ok {
				wg.Add(1)
				go handleConn(ctx, wg, conn)
			}
		}
	}
}

func accept(listener net.Listener) <-chan net.Conn{
	ch := make(chan net.Conn)
	go func(){
		defer close(ch)
		if conn, err := listener.Accept(); err == nil{
			ch <- conn
			log.Printf("new conn established: %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
		} else {
			log.Printf("listener accept failed: %v", err)
		}
	}()
	return ch
}

func handleConn(ctx context.Context, wg *sync.WaitGroup, conn net.Conn) {
	ch := make(chan string)
	entering <- ch
	defer func(){
		fmt.Printf("leaving <- %v\n", ch)
		leaving <- ch
		if err := conn.Close(); err != nil{
			log.Printf("closing conn [%s] err: %v\n", conn.LocalAddr(), err)
		} else{
			log.Printf("conn [%s] closed\n", conn.LocalAddr())
		}
		wg.Done()
	}()
	tck := time.NewTicker(cfg.TickInterval)
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ch:
			if err := sendMsg(conn, fmt.Sprintf("msg: %s\n", msg)); err != nil{
				return
			}
		case t := <-tck.C:
			if err := sendMsg(conn, fmt.Sprintf("now: %s\n", t)); err != nil{
				return
			}
		}
	}
}

func sendMsg(conn net.Conn, msg string) (err error){
	_, err = fmt.Fprint(conn, msg)
	return
}

func broadcaster(){
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				fmt.Printf("msg [%s] -> %v\n", msg, cli)
				cli <- msg
			}
		case cli := <-entering:
			fmt.Printf("open %v := <-entering\n", cli)
			clients[cli] = true
		case cli := <-leaving:
			fmt.Printf("close %v <-leaving\n", cli)
			delete(clients, cli)
			close(cli)
		}
	}
}

func messageProducer(){
	for {
		var msg string
		// bufio.NewReader(os.Stdin).ReadString('\n')
		if _, err := fmt.Scanf("%s\n", &msg); err == nil{
			log.Printf("got new msg: %s\n", msg)
			messages <- msg
		} else {
			log.Printf("scanf err: %v\n", err)
			return
		}
	}
}
