package main

import (
	pb "axitech/protos"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	flag1 := flag.String("sh", "_", "shorten link")
	flag2 := flag.String("rev", "_", "reveal a link")
	flag.Parse()
	if ((*flag1 == "_") && (*flag2 == "_")) || ((*flag1 != "_") && (*flag2 != "_")) {
		log.Fatalf("incorrect input data")
	}

	conn, err := grpc.Dial(":9005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect! %v", err)
	}
	defer conn.Close()
	c := pb.NewDataServiceClient(conn)

	if *flag1 != "_" {
		msg := pb.OriginalURL{Body: *flag1}
		data, err2 := c.Shorten(context.Background(), &msg)
		if err2 != nil {
			log.Panicln("Error while shortening a link: ", err2)
		}
		log.Println("Result: ", data.Body)
	}

	if *flag2 != "_" {
		msg2 := pb.NewURL{Body: *flag2}
		data2, err2 := c.Reveal(context.Background(), &msg2)
		if err2 != nil {
			log.Panicln("Error while revealing a link: ", err2)
		}
		log.Println("Result: ", data2.Body)
	}
}
