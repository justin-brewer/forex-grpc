package main

import (
	pb "../converter"
	"context"
	"encoding/json"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var (
	addr = flag.String("addr", ":50051", "Network host:port to listen on for gRPC connections.")
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

type rates struct {
	Rates map[string]float64
	Base  string
	Date  string
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetConversion(ctx context.Context, in *pb.ConversionRequest) (*pb.ConversionReply, error) {
	log.Printf("Handling GetConversion request [%v] with context %v", in, ctx)
	target := getConversion(ctx, in)
	return &pb.ConversionReply{Message: "Hello " + in.Name}, nil
}

func getConversion(ctx context.Context, in *pb.ConversionRequest) float32 {
	if conversionTable == nil {
		conversionTable = getRates()
	}

}

func getRates() {
	resp, err := http.Get("https://api.exchangeratesapi.io/latest")
	if err != nil {
		log.Printf("Error getting exchange rates:\n")
		log.Printf("%v\n", err)
	}

	xrates := rates{}

	conversionTable := json.UnMarshall([]byte(resp), &xrates)
	return resp
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterConverterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func helper() {

}
