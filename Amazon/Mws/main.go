package main

import (
	"fmt"

	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mws/products"
	"github.com/svvu/gomws/mwsHttps"
)

var (
	// SellerID or merchant id from user
	SellerID = "-"
	// AuthToken from user
	AuthToken = "-"
	// Region from user
	Region = "US"
	// AccessKey is from main account
	AccessKey = "-"
	// SecretKey is from main account
	SecretKey = "-"
)

func getConfigFile() gmws.MwsConfig {
	return gmws.MwsConfig{
		SellerId:  SellerID,
		AuthToken: AuthToken,
		Region:    Region,

		// Optional if already set in env variable
		AccessKey: AccessKey,
		SecretKey: SecretKey,
	}
}

func main() {

	response := getMatchingProduct([]string{"B00AVWKUJS"})

	xmlParser := gmws.NewXMLParser(response)
	// Check whether or not API send back error message
	if xmlParser.HasError() {
		fmt.Println(xmlParser.GetError())
	}

	gmp := products.GetMatchingProductResult{}
	xmlParser.Parse(&gmp)
	// Individual result might have error
	for _, r := range gmp.Results {
		if r.Error != nil {
			fmt.Println(r.Error)
		} else {
			fmt.Println(r.Products)
		}
	}
}

// getLowestPricedOffersForASIN  - 200 requests per hour
func getLowestPricedOffersForASIN(asin, itemCondition string) *mwsHttps.Response {
	productsClient, _ := products.NewClient(getConfigFile())
	response := productsClient.GetLowestPricedOffersForASIN(asin, itemCondition)
	if response.Error != nil {
		fmt.Println(response.Error.Error())
		return nil
	}
	return response
}

func getMatchingProduct(asin []string) *mwsHttps.Response {
	productsClient, _ := products.NewClient(getConfigFile())
	fmt.Println("------GetMatchingProduct------")
	response := productsClient.GetMatchingProduct([]string{"B00AVWKUJS"})
	if response.Error != nil {
		fmt.Println(response.Error.Error())
		return nil
	}
	return response
}
