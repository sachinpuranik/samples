package main

import (
	"flag"
	"log"

	"github.com/sachinpuranik/samples/micro-service/pkg/cmd"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	staticAdd := "127.0.0.1:5050"

	address = &staticAdd

	// Set up a connection to the server.

	c := cmd.NewClient()
	c.ConnectGRPCServer(*address)

	cfg := cmd.HttpServerConfig{7070}
	c.RunClientService(cfg)

	log.Printf("started server at 127.0.0.1:%d", cfg.Port)

	halt := make(chan int, 1)
	<-halt
}
