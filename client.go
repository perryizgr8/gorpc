package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"perryizgr8.com/gorpc/procedures"
)

func main() {
	fmt.Println("starting client...")
	cc, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer cc.Close()
	c := procedures.NewPingClient(cc)
	res, err := c.Ping(context.Background(), &procedures.EchoRequest{Msg: "Ping!"})
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
	fmt.Printf("Ping!... %v\n", res)
}
