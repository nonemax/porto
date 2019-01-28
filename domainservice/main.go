package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/nonemax/porto/domainservice/db/pq"
	"github.com/nonemax/porto/domainservice/server"
)

var (
	psqlClientConfig string
	listenAddress    string
	wait             int
)

func init() {
	flag.StringVar(&psqlClientConfig, "psql_cfg", "", "PostgreSQL client config, e.g. \"postgres://user:password@localhost/dbname?sslmode=disable\"")
	flag.StringVar(&listenAddress, "lst_addr", ":8080", "Listen address")
	flag.IntVar(&wait, "wait", 0, "Parser buffer size")
}

func main() {
	flag.Parse()
	if len(psqlClientConfig) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if wait > 0 {
		time.Sleep(time.Duration(wait) * time.Second)
	}
	db, err := pq.NewDB(psqlClientConfig)
	if err != nil {
		log.Printf("Failed to create psql connection: %s", err.Error())
		os.Exit(1)
	}
	s := server.New(listenAddress, db)
	log.Println("Start server")
	server.Start(&s)
}
