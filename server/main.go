package main

import (
	pb "../converter"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var (
	//	addr   = flag.String("addr", ":50051", "Network host:port to listen on for gRPC connections.")
	port   = ":50051"
	rates  = xrates{}
	source = "get"
)

type server struct{}

type xrates struct {
	Rates map[string]float32
	Base  string
	Date  string
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetConversion(ctx context.Context, in *pb.ConversionRequest) (*pb.ConversionReply, error) {
	log.Printf("Handling GetConversion request [%v] with context %v", in, ctx)
	fmt.Printf("Args: Source: %v, Target: %v, Amount: %v", in.Source, in.Target, in.Amount)
	target := getConversion(ctx, in)
	return &pb.ConversionReply{Amount: target}, nil
}

func listCurrencies() string {

}

func getConversion(ctx context.Context, in *pb.ConversionRequest) float32 {
	if rates.Date == "" {
		getRates()
	}

	newAmount := (rates.Rates[in.Target] / rates.Rates[in.Source]) * in.Amount
	return newAmount
}

func getRates() {
	var rsp string
	if source == "get" {
		resp, err := http.Get("https://api.exchangeratesapi.io/latest")
		if err != nil {
			log.Printf("error with GET")
			log.Printf("%v\n", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("%v\n", err)
		}
		rsp = string(body)
	} else {
		rsp = `{"rates":{"CAD":1.5524,"HKD":8.7466,"ISK":148.64,"PHP":57.645,"DKK":7.4727,"HUF":338.37,"CZK":26.203,"AUD":1.7674,"RON":4.8213,"SEK":10.8945,"IDR":16434.0,"INR":83.468,"BRL":5.5081,"RUB":84.0284,"HRK":7.6,"JPY":116.84,"THB":35.586,"CHF":1.0549,"SGD":1.5779,"PLN":4.3599,"BGN":1.9558,"TRY":7.0361,"CNY":7.8877,"NOK":11.3682,"NZD":1.8173,"ZAR":18.4447,"USD":1.124,"MXN":24.8028,"ILS":4.0909,"GBP":0.88623,"KRW":1359.4,"MYR":4.7944},"base":"EUR","date":"2020-03-12"}`
	}

	errR := json.Unmarshal([]byte(rsp), &rates)
	if errR != nil {
		log.Printf("err:\n%v", errR)
	}
	rates.Rates[rates.Base] = 1.00
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("started listening on port: %v\n", 50051)
	s := grpc.NewServer()
	pb.RegisterConverterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
