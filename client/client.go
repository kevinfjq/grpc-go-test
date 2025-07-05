package main

import (
	"context"
	pb "github.com/kevinfjq/proto_example/coffeeshop_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatalf("Failed to call GetMenu: %v", err)
	}

	done := make(chan bool)

	var items []*pb.Item

	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Failed to recv: %v", err)
			}

			items = resp.Items
			log.Printf("Received items: %v", resp.Items)
		}
	}()

	<-done
	receipt, err := c.PlaceOrder(ctx, &pb.Order{Items: items})
	log.Printf("Place order: %v", receipt)

	status, err := c.GetOrderStatus(ctx, receipt)
	log.Printf("Get order status: %v", status)
}
