package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=googl")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	quote := new(QuoteResponse)
	xml.Unmarshal(body, &quote)

	fmt.Printf("%s: %.2f", quote.Name, quote.LastPrice)
	fmt.Printf("\ntime took=%s", time.Since(start))
}

type QuoteResponse struct {
	Status           string
	Name             string
	Symbol           string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	Timestamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
