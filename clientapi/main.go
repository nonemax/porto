package main

import (
	"flag"
	"log"
	"os"
	"time"

	t "github.com/nonemax/porto-transport"
	"github.com/nonemax/porto/clientapi/parser"
	"github.com/nonemax/porto/clientapi/server"
	"google.golang.org/grpc"
)

var (
	listenAddress string
	address       string
	filename      string
	bufferSize    int
	wait          int
)

func init() {
	flag.StringVar(&address, "addr", "localhost:8080", "Server address")
	flag.StringVar(&listenAddress, "lst_addr", ":8081", "Listen address")
	flag.StringVar(&filename, "file_name", "ports.json", "Listen address")
	flag.IntVar(&bufferSize, "buf_size", 10, "Parser buffer size")
	flag.IntVar(&wait, "wait", 0, "Parser buffer size")
}

func main() {
	flag.Parse()
	if len(listenAddress) == 0 || len(address) == 0 || len(filename) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if wait > 0 {
		time.Sleep(time.Duration(wait) * time.Second)
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cant connect: %v", err)
	}

	c := t.NewTransportClient(conn)
	sender := parser.NewSender(c)
	p := parser.New(c, bufferSize, filename, &sender)
	log.Println("Start parser")
	go p.Start()

	s, err := server.New(listenAddress, c)
	if err != nil {
		log.Fatalf("Cant start server %v", err)
	}
	log.Println("Start server")
	s.Start()
}
