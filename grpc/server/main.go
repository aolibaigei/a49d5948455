package main

import (
	"context"
	"fmt"
	"log"
	"net"

	msgexamplepb "../pbfile"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Send(ctx context.Context, req *msgexamplepb.MsgRequest) (*msgexamplepb.MsgResp, error) {
	fmt.Println("rpc invoked : %v", req)

	text := req.GetMsg().GetMsgText()

	res_text := text + "aaaa"

	res := &msgexamplepb.MsgResp{
		Recvmsg: res_text,
	}
	return res, nil
}

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

	msgexamplepb.RegisterMsgServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("RPC Serve Error: %v", err)
	}

}
