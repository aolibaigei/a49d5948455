package main

import (
	"context"
	"fmt"
	"log"

	msgexamplepb "../pbfile"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Client")
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connect Error %v ", err)
	}

	defer conn.Close()

	c := msgexamplepb.NewMsgServiceClient(conn)
	req := &msgexamplepb.MsgRequest{
		Msg: &msgexamplepb.MSG{
			MsgId:   1,
			MsgText: "Hellow world",
		},
	}
	rep, err := c.Send(context.Background(), req)

	if err != nil {
		log.Fatal("send error")
	}
	fmt.Println(rep)
	// fmt.Println("client : %v", c)

}
