 package main

 import (
	"log"

	"context"
	"google.golang.org/grpc"
	
	"server/chat"
 )

 funct main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dail(:9000, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClieent(conn)

	message := chat.Message{
		Body: "Hello from the client!"
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nill {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
 }
