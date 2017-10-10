package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// gathers stock prices concurrently and processes them as they are returned

func main() {

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"neo",
	}

	numComplete := 0
	for _, symbol := range stockSymbols {
		// pass the symbol to our anonomous function, all go routines will fire concurrently
		// in the order they are created
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			var quote QuoteResponse
			xml.Unmarshal(body, &quote)
			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	// check if the go routines are done executing
	for numComplete < len(stockSymbols) {
		// sleep for 10 milliseconds before checking again
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volumn           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
