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

	if len(os.Args) == 2 {
		if os.Args[1] == "list" {
			listCurrencies()
		}
	} else {
		getConversion(os.Args)
	}
}

func listCurrencies() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(defaultServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := converter.NewConverterClient(conn)
	r, err := c.GetCurrencyList(ctx, &converter.ListRequest{})
	if err != nil {
		log.Printf("Error:\n%v", err)
	}
	//	log.Printf("Server Reply: %s\n", r.GetStatus())
	log.Printf("available currencies:\n%v", r.Reply)
}

func getConversion(args []string) {

	if len(args) != 4 {
		fmt.Println("Usage:\ngo run main.go <source-curreny> <target-currency> <amount>")
		return
	}
	source := args[1]
	target := args[2]
	amount, err := strconv.ParseFloat(args[3], 32)
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
	defer conn.Close()
	c := converter.NewConverterClient(conn)
	r, err := c.GetConversion(ctx, &converter.ConversionRequest{Source: source, Target: target, Amount: float32(amount)})
	if err != nil {
		log.Printf("Error:\n%v", err)
	}
	//	log.Printf("Server Reply: %s\n", r.GetStatus())
	fmt.Printf("%v of %v = %v of %v\n", amount, source, r, target)
}
