package main

import (
	"fmt"
	protos "github.com/karanr1990/go-grpc-app/protos/translation"
	"github.com/karanr1990/go-grpc-app/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	fmt.Println("welcome to grpc")

	//create grpc server
	s := grpc.NewServer()

	// create new instance of Translation server
	trans := server.NewTranslation()

	reflection.Register(s)
	protos.RegisterTranslationServer(s, trans)

	// create socket to listen to requests
	tl, err := net.Listen("tcp", "localhost:8765")
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8765", err))
	}

	// start listening
	s.Serve(tl)

}
