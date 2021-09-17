package config

import (
	"flag"
	. "github.com/nj-eka/gobe1/env"
	"time"
)

var (
	Address   string
	Protocol string
	KeepAlive time.Duration
	TestTimeout time.Duration
)

func init(){
	flag.StringVar(&Address, "a", GetEnv("SMCHAT_ADDRESS", ":8000"), "address")
	flag.StringVar(&Protocol, "p", GetEnv("SMCHAT_PROTOCOL", "tcp"), "protocol")
	flag.DurationVar(&KeepAlive, "k", 1 * time.Minute, "keep alive")
	flag.DurationVar(&TestTimeout, "t", 10 * time.Second, "test timeout")
	flag.Parse()
}
