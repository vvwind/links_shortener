package main

import (
	db "axitech/pkg/db"
	pkg "axitech/pkg/services"
	pb "axitech/protos"
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

type simpleServer struct {
	pb.UnimplementedDataServiceServer
	data          map[string]string
	data_reversed map[string]string
	mu            sync.Mutex
	rnd           *rand.Rand
	mode          string
	db            *db.DbData
}

func (s *simpleServer) Init() {
	s.data = make(map[string]string)
	s.data_reversed = make(map[string]string)
}
func (s *simpleServer) Shorten(ctx context.Context, in *pb.OriginalURL) (*pb.NewURL, error) {
	s.mu.Lock()
	if _, ok := s.data[in.Body]; !ok {
		newUrl := pkg.RandStringBytes(s.rnd)
		s.data[in.Body] = newUrl
		for {
			if _, present := s.data_reversed[newUrl]; present {
				newUrl = pkg.RandStringBytes(s.rnd)
				s.data[in.Body] = newUrl
			} else {
				s.data_reversed[newUrl] = in.Body
				break
			}
		}
		if s.mode == "db" {
			errW := s.db.WriteOriginal(in.Body, newUrl)
			if errW != nil {
				return nil, errW
			}

			errW2 := s.db.WriteNew(newUrl, in.Body)
			if errW2 != nil {
				return nil, errW
			}

		}
	} else {
		s.mu.Unlock()
		return &pb.NewURL{Body: ""}, errors.New("this link was already shortened!")
	}

	s.mu.Unlock()
	return &pb.NewURL{Body: s.data[in.Body]}, nil
}
func (s *simpleServer) Reveal(ctx context.Context, in *pb.NewURL) (*pb.OriginalURL, error) {
	fmt.Println(in.Body)
	_, ok := s.data_reversed[in.Body]
	fmt.Println(ok)
	if s.mode == "db" {
		data, err := s.db.Reveal(in.Body)
		if err != nil {
			return nil, err
		} else {
			return &pb.OriginalURL{Body: data}, nil
		}
	}
	return &pb.OriginalURL{Body: s.data_reversed[in.Body]}, nil
}

func main() {
	flag1 := flag.String("m", "mem", "data storage type")
	flag.Parse()
	if (*flag1 != "db") && (*flag1 != "mem") {
		log.Fatalf("incorrect storage type")
	}
	myport := ":9005"
	lis, err := net.Listen("tcp", myport)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	myServer := &simpleServer{mode: *flag1}
	myServer.Init()

	pb.RegisterDataServiceServer(s, myServer)
	log.Printf("Server is listening at %v", lis.Addr())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	myServer.rnd = r1

	if *flag1 == "db" {
		database := db.Init()
		errSt := database.StartDB()
		myServer.db = database
		if errSt != nil {
			log.Printf("Failed to start DB %v", errSt)
		} else {
			log.Println("DB has succesfully started!")
		}
	}
	if err2 := s.Serve(lis); err2 != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
