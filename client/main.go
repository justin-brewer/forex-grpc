package main

import (
	"context"
	//	"flag"
	converter "../converter"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	defaultServer = "localhost:50051"
)

// args (expected number = 4)
// args[0] : go executable name
// args[1] source string: currency symbol (e.g. "USD")
// args[2] target string: currency symbol (e.g. "EUR")
// args[3] amount float32: source converted to target
func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage:\ngo run main.go <source-curreny> <target-currency> <amount>")
		return
	}
	source := os.Args[1]
	target := os.Args[2]
	amount, err := strconv.ParseFloat(os.Args[3], 32)
	if err != nil {
		log.Println("amount must be a float")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(defaultServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	log.Println("connected to server")
	defer conn.Close()
	c := converter.NewConverterClient(conn)
	r, err := c.GetConversion(ctx, &converter.ConversionRequest{Source: source, Target: target, Amount: float32(amount)})
	if err != nil {
		log.Printf("Error:\n%v", err)
	}
	//	log.Printf("Server Reply: %s\n", r.GetStatus())
	log.Printf("%v of %v = %v of %v\n", amount, source, r, target)

}
