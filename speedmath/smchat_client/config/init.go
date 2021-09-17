package config

import (
	"flag"
	. "github.com/nj-eka/gobe1/env"
	"time"
)

var (
	Address   string
	Protocol string
	Timeout   time.Duration
	KeepAlive time.Duration
)

func init(){
	flag.StringVar(&Address, "a", GetEnv("SMCHAT_ADDRESS", ":8000"), "address")
	flag.StringVar(&Protocol, "p", GetEnv("SMCHAT_PROTOCOL", "tcp"), "protocol")
	flag.DurationVar(&Timeout, "timeout", 1*time.Second, "response timeout")
	flag.DurationVar(&KeepAlive, "keepAlive", 1 * time.Minute, "keep alive duration")
	flag.Parse()
}
