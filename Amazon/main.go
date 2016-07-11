package main

// go run .\main.go .\model.go .\mwscalls.go .\xmlhelpers.go .\parsers.go

import (
	"GoSamples/Amazon/Mws"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("start")
	listProducts, _ := Mws.GetProductsByASIN("B00IGR5EQE")
	//getLowestOfferListingsForASIN("B00078ZLLI", "New")
	//	listProducts, _ := parseLowestOfferListingsForASIN("B00IGR5EQE", "New", ProductTracking{})
	//listProducts, _ := getProductsByKeyword("Spin Master Games - Moustache Smash", "New")
	// if listProductsError != nil {
	// 	fmt.Println("error getting products by keyword : " + listProductsError.Error())
	// 	return
	// }
	//
	res2B1mws, _ := json.Marshal(listProducts)
	fmt.Println("item : ", string(res2B1mws))

	fmt.Println("stop")
}
