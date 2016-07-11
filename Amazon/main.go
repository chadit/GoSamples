package main

// go run .\main.go .\model.go .\mwscalls.go .\xmlhelpers.go .\parsers.go

import (
	"GoSamples/Amazon/Mws"
	"encoding/json"
	"fmt"
)

var (
	// SellerID or merchant id from user
	SellerID = ""
	// AuthToken from user
	AuthToken = ""
	// Region from user
	Region = "US"
	// AccessKey is from main account
	AccessKey = ""
	// SecretKey is from main account
	SecretKey = ""
)

func main() {
	fmt.Println("start")

	//listProducts, listProductsError := Mws.GetProductsByASIN("B00IGR5EQE", SellerID, AuthToken, Region, AccessKey, SecretKey)
	//getLowestOfferListingsForASIN("B00078ZLLI", "New")
	//	listProducts, _ := parseLowestOfferListingsForASIN("B00IGR5EQE", "New", ProductTracking{})
	listProducts, listProductsError := Mws.GetProductsByKeyword("Spin Master Games - Moustache Smash", "New", SellerID, AuthToken, Region, AccessKey, SecretKey)
	// if listProductsError != nil {
	// 	fmt.Println("error getting products by keyword : " + listProductsError.Error())
	// 	return
	// }
	//
	if listProductsError != nil {
		fmt.Println("error")
	} else {

		res2B1mws, _ := json.Marshal(listProducts)
		fmt.Println("item : ", string(res2B1mws))
	}
	fmt.Println("stop")
}
