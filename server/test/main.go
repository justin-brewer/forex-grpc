package main

import (
	"encoding/json"
	"fmt"
	//	"reflect"
	//	"io/ioutil"
	"log"
	//	"net/http"
)

func getRates() {
	//	resp, err := http.Get("https://api.exchangeratesapi.io/latest")
	//	if err != nil {
	//		log.Printf("error with GET")
	//		log.Printf("%v\n", err)
	//	}
	//	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		log.Printf("%v\n", err)
	//	}

	//rsp := string(body)
	rsp := `{"rates":{"CAD":1.5524,"HKD":8.7466,"ISK":148.64,"PHP":57.645,"DKK":7.4727,"HUF":338.37,"CZK":26.203,"AUD":1.7674,"RON":4.8213,"SEK":10.8945,"IDR":16434.0,"INR":83.468,"BRL":5.5081,"RUB":84.0284,"HRK":7.6,"JPY":116.84,"THB":35.586,"CHF":1.0549,"SGD":1.5779,"PLN":4.3599,"BGN":1.9558,"TRY":7.0361,"CNY":7.8877,"NOK":11.3682,"NZD":1.8173,"ZAR":18.4447,"USD":1.124,"MXN":24.8028,"ILS":4.0909,"GBP":0.88623,"KRW":1359.4,"MYR":4.7944},"base":"EUR","date":"2020-03-12"}`

	var xrates interface{}
	errR := json.Unmarshal([]byte(rsp), &xrates)
	if errR != nil {
		log.Printf("err:\n%v", errR)
	}

	m := xrates.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	//	fmt.Println(m["base"])
	//	rt := m["rates"]
	//	fmt.Println("rates:")
	//	fmt.Println(rt)
	//	fmt.Println("type:")
	//	fmt.Println(reflect.TypeOf(rt))
	//	fmt.Println(rt["AUD"])

	//	map[base:EUR date:2020-03-12 rates:map[AUD:1.7674 BGN:1.9558 BRL:5.5081 CAD:1.5524 CHF:1.0549 CNY:7.8877 CZK:26.203 DKK:7.4727 GBP:0.88623 HKD:8.7466 HRK:7.6 HUF:338.37 IDR:16434 ILS:4.0909 INR:83.468 ISK:148.64 JPY:116.84 KRW:1359.4 MXN:24.8028 MYR:4.7944 NOK:11.3682 NZD:1.8173 PHP:57.645 PLN:4.3599 RON:4.8213 RUB:84.0284 SEK:10.8945 SGD:1.5779 THB:35.586 TRY:7.0361 USD:1.124 ZAR:18.4447]]
}

func main() {
	getRates()
}
