package config

import (
	"flag"
	. "github.com/nj-eka/gobe1/env"
	"time"
)

var (
	Address   string
	Protocol string
	TickInterval time.Duration
	KeepAlive time.Duration
)

func init(){
	flag.StringVar(&Address, "a", GetEnv("TIMETICK_ADDRESS", ":8000"), "address")
	flag.StringVar(&Protocol, "p", GetEnv("TIMETICK_PROTOCOL", "tcp"), "protocol")
	flag.DurationVar(&KeepAlive, "keepAlive", 1 * time.Minute, "keep alive duration")
	flag.DurationVar(&TickInterval, "tick", 1 * time.Second, "tick interval")
	flag.Parse()
}
