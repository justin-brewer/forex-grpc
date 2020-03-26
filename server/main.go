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

//const endpoint = "http://data.fixer.io/api/latest?access_key=4196e8bc7ff14104a867e574540057a9"
const api_url = "http://data.fixer.io/api/latest"
const api_key = "4196e8bc7ff14104a867e574540057a9"

var endpoint = fmt.Sprintf("%v?access_key=%v", api_url, api_key)

type server struct{}

type xrates struct {
	Success   bool
	Timestamp int
	Base      string
	Date      string
	Rates     map[string]float32
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetConversion(ctx context.Context, in *pb.ConversionRequest) (*pb.ConversionReply, error) {
	log.Printf("Handling GetConversion request [%v] with context %v", in, ctx)
	fmt.Printf("Args: Source: %v, Target: %v, Amount: %v", in.Source, in.Target, in.Amount)
	target := getConversion(ctx, in)
	return &pb.ConversionReply{Amount: target}, nil
}

func listCurrencies() string {
	list := ""
	return list
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
		resp, err := http.Get(endpoint)
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
		rsp = `{"success":true,"timestamp":1585261445,"base":"EUR","date":"2020-03-26","rates":{"AED":4.049721,"AFN":84.673933,"ALL":126.190057,"AMD":548.247044,"ANG":1.977177,"AOA":591.733047,"ARS":70.707016,"AUD":1.826082,"AWG":1.984644,"AZN":1.877768,"BAM":1.971106,"BBD":2.22708,"BDT":93.761897,"BGN":1.957068,"BHD":0.415387,"BIF":2090.491524,"BMD":1.10258,"BND":1.587575,"BOB":7.610279,"BRL":5.538588,"BSD":1.10457,"BTC":0.000164,"BTN":83.141858,"BWP":13.087382,"BYN":2.84671,"BYR":21610.566386,"BZD":2.220085,"CAD":1.550029,"CDF":1888.16528,"CHF":1.062335,"CLF":0.03313,"CLP":914.152134,"CNY":7.799541,"COP":4389.370652,"CRC":637.319176,"CUC":1.10258,"CUP":29.218368,"CVE":111.305632,"CZK":27.174624,"DJF":195.950458,"DKK":7.464086,"DOP":59.6714,"DZD":136.221534,"EGP":17.363651,"ERN":16.538525,"ETB":36.197879,"EUR":1,"FJD":2.532402,"FKP":0.908343,"GBP":0.907863,"GEL":3.782136,"GGP":0.908343,"GHS":6.334285,"GIP":0.908343,"GMD":56.13783,"GNF":10391.815887,"GTQ":8.746019,"GYD":231.124908,"HKD":8.547529,"HNL":27.614123,"HRK":7.612302,"HTG":104.299143,"HUF":355.724974,"IDR":17925.078343,"ILS":3.948758,"IMP":0.908343,"INR":82.379372,"IQD":1312.070102,"IRR":46424.127048,"ISK":154.008623,"JEP":0.908343,"JMD":149.536971,"JOD":0.781742,"JPY":120.639915,"KES":115.869724,"KGS":88.428559,"KHR":4470.961465,"KMF":495.471827,"KPW":992.39065,"KRW":1337.726607,"KWD":0.34114,"KYD":0.920459,"KZT":492.425087,"LAK":9848.243674,"LBP":1669.305763,"LKR":206.602374,"LRD":217.925518,"LSL":19.174144,"LTL":3.255632,"LVL":0.66694,"LYD":1.582243,"MAD":10.838783,"MDL":19.925267,"MGA":4130.264123,"MKD":62.176865,"MMK":1547.418432,"MNT":3059.395947,"MOP":8.820296,"MRO":393.621369,"MUR":43.169637,"MVR":16.990557,"MWK":810.395935,"MXN":25.730775,"MYR":4.808627,"MZN":73.409652,"NAD":19.173762,"NGN":404.646657,"NIO":37.536058,"NOK":11.487675,"NPR":133.028748,"NZD":1.855631,"OMR":0.424559,"PAB":1.104459,"PEN":3.74602,"PGK":3.770752,"PHP":56.876591,"PKR":181.236604,"PLN":4.520302,"PYG":7269.40543,"QAR":4.014475,"RON":4.834257,"RSD":117.502243,"RUB":85.422711,"RWF":1030.912223,"SAR":4.141394,"SBD":9.122788,"SCR":15.117289,"SDG":60.970245,"SEK":11.00981,"SGD":1.576023,"SHP":0.908343,"SLL":10708.804977,"SOS":645.008867,"SRD":8.223068,"STD":24314.943833,"SVC":9.663894,"SYP":567.224038,"SZL":19.174269,"THB":35.946929,"TJS":11.271856,"TMT":3.870056,"TND":3.175675,"TOP":2.614987,"TRY":7.053816,"TTD":7.463242,"TWD":33.282482,"TZS":2551.481109,"UAH":31.170132,"UGX":4291.077752,"USD":1.10258,"UYU":49.46636,"UZS":10491.048092,"VEF":11.01202,"VND":25630.021476,"VUV":138.518054,"WST":3.074351,"XAF":661.06824,"XAG":0.076723,"XAU":0.000676,"XCD":2.979777,"XDR":0.816614,"XOF":663.753047,"XPF":120.649772,"YER":276.031219,"ZAR":19.102743,"ZMK":9924.541857,"ZMW":19.683266,"ZWL":355.030735}}`
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
