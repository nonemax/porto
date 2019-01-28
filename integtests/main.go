package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/nonemax/porto/integtests/tester"
)

var (
	address string
	wait    int
)

func init() {
	flag.StringVar(&address, "addr", "localhost:8080", "Server address")
	flag.IntVar(&wait, "wait", 0, "Parser buffer size")
}

func main() {
	flag.Parse()
	if len(address) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if wait > 0 {
		time.Sleep(time.Duration(wait) * time.Second)
	}
	t := tester.New(address)
	err := t.Test()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("Everything looks good")
}
