package main

import (
	"fmt"
	"log"
	"net"

	msgexamplepb "./pbfile"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	msg := &msgexamplepb.MSG{
		MsgId:   12,
		MsgText: "hello world",
	}
	fmt.Println(msg)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Listen Error %v", err)
	}
	s := grpc.NewServer()

	msgexamplepb.RegisterSendMSGServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("RPC Serve Error: %v", err)
	}

}
