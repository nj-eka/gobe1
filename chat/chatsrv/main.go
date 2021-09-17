package main

import (
	"bufio"
	"context"
	"fmt"
	cfg "github.com/nj-eka/gobe1/chat/chatsrv/config"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

const (
	csReset = "\033[0m"
	csBold = "\033[1m"
	csUnderline = "\033[4m"

	csRed   = "\033[31m"
	csGreen = "\033[32m"
	csYellow = "\033[33m"
	csBlue   = "\033[34m"
	csPurple = "\033[35m"
	csCyan   = "\033[36m"
)

var colors = [...]string{csRed, csGreen, csYellow, csBlue, csPurple, csCyan}

type cid struct{
	id int
	name string
	addr string
	cs   string
}

type client struct{
	ref *cid
	ch chan <- message
}

type message struct{
	from *cid
	text string
}

var (
	bot = &cid{id: 0, name: "bot", cs: csUnderline}
	//admin = &cid{id:1, name:"admin", cs: csBold}

	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan message, 1024)
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
	log.Printf("Chat server [pid: %d] started at %s://%s\n", os.Getpid(), listener.Addr().Network(), listener.Addr())
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
	ch := make(chan message)
	cli := client{
		ref: &cid{
			name: conn.LocalAddr().String(),
			addr: conn.RemoteAddr().String(),
		},
		ch: ch,
	}
	defer func(){
		log.Println("conn [%s -> %s] closed", conn.RemoteAddr(), conn.LocalAddr())
		_ = conn.Close()
		wg.Done()
	}()
	// set name
	if _, err := fmt.Fprintln(conn, "enter your nick:"); err != nil{
		return
	}
	input := bufio.NewScanner(conn)
	if input.Scan(){
		cli.ref.name = input.Text()		
	} else {
		return
	}
	// set color
	if _, err := fmt.Fprintln(conn, "choose your color [1-7]:"); err != nil{
		return
	}
	if input.Scan(){
		if i, err := strconv.ParseInt(input.Text(), 0 , 0); err == nil{
			if (i >= 0) && (i < int64(len(colors))) {
				cli.ref.cs = colors[i]
			}
		}
	} else {
		return
	}
	// welcome
	entering <- cli // update id

	// chat -> client
	go clientWriter(conn, cli.ref, ch)

	// client -> chat
	clientWrote := clientReader(input)
	for {
		select{
		case <-ctx.Done():
			_, _ = fmt.Fprintln(conn, "see you later")
			return
		case text, more := <- clientWrote:
			if !more{
				leaving <- cli
				return
			}
			messages <- message{
				from: cli.ref,
				text: text,
			}
		}
	}
}

func clientReader(scanner *bufio.Scanner) <- chan string{
	ch := make(chan string)
	go func(){
		defer close(ch)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()
	return ch
}

func clientWriter(conn net.Conn, reader *cid, ch <- chan message ) {
	for msg := range ch {
		_, _ = fmt.Fprint(conn, msg.from.cs)
		if msg.from.id == reader.id{
			_, _ = fmt.Fprint(conn, "you: ", msg.text)
		} else {
			_, _ = fmt.Fprintf(conn, "%s(%d): %s",  msg.from.name, msg.from.id, msg.text)
		}
		_, _ = fmt.Fprint(conn, csReset, "\n")
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	count := 2
	for {
		select {
		case cli := <-entering:
			count++
			cli.ref.id = count
			clients[cli] = true
			messages <- message{
				from: bot,
				text: fmt.Sprintf("user [%s:%d] from [%s] has joined", cli.ref.name, cli.ref.id, cli.ref.addr),
			}
		case cli := <-leaving:
			messages <- message{
				from: bot,
				text: fmt.Sprintf("user [%s:%d] from [%s] has left", cli.ref.name, cli.ref.id, cli.ref.addr),
			}
			delete(clients, cli)
			close(cli.ch)
		case msg := <- messages:  // read out all messages
			for {
				for cli := range clients {
					cli.ch <- msg
				}
				if len(messages) == 0{
					break
				}
				msg = <- messages
			}
		}
	}
}
