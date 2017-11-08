package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

// Extract, Transform, and Load
func main() {
	strart := time.Now()

	extractChannel := make(chan *Order)
	transformChannel := make(chan *Order)
	doneChannel := make(chan bool)

	go extract(extractChannel)
	go tranform(extractChannel, transformChannel)
	go load(transformChannel, doneChannel)

	<-doneChannel
	fmt.Println(time.Since(strart))
}

type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

type Order struct {
	CustomerNumber int
	PartNumber     string
	Quantity       int

	UnitCost  float64
	UnitPrice float64
}

func extract(ch chan *Order) {
	f, _ := os.Open("./orders.txt")
	r := csv.NewReader(f)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		ch <- order
	}
	close(ch)
}

func tranform(extractChannel, transformChannel chan *Order) {
	f, _ := os.Open("./productList.txt")
	defer f.Close()
	r := csv.NewReader(f)

	records, _ := r.ReadAll()
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}

	var numMessages int64
	for o := range extractChannel {
		atomic.AddInt64(&numMessages, 1)
		//numMessages++
		//	fmt.Println("e : ", numMessages)
		go func(o *Order) {
			// simulate web call
			time.Sleep(3 * time.Millisecond)
			o.UnitCost = productList[o.PartNumber].UnitCost
			o.UnitPrice = productList[o.PartNumber].UnitPrice
			transformChannel <- o
			atomic.AddInt64(&numMessages, -1)
			//numMessages--
		}(o)
	}

	// checks to see if messages are still being processed before closing the channel
	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("transform closed")
	close(transformChannel)
}

func load(transformChannel chan *Order, doneChannel chan bool) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()

	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n",
		"Part Number", "Quantity", "Unit Cost",
		"Unit Price", "Total Cost", "Total Price")

	var numMessages int64
	for o := range transformChannel {
		//	fmt.Println("t : ", numMessages)
		atomic.AddInt64(&numMessages, 1)
		go func(o *Order) {
			time.Sleep(1 * time.Millisecond)
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
				o.PartNumber, o.Quantity,
				o.UnitCost, o.UnitPrice,
				o.UnitCost*float64(o.Quantity),
				o.UnitPrice*float64(o.Quantity))
			atomic.AddInt64(&numMessages, -1)
			//numMessages--
		}(o)
	}
	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("done called")
	doneChannel <- true
}
