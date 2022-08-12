package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/adrianliechti/llama/pkg/provider/custom"

	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	s := &Server{}

	server := grpc.NewServer()

	custom.RegisterCompleterServer(server, s)

	if err := server.Serve(l