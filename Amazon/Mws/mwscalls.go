package Mws

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mws/products"
	"github.com/svvu/gomws/mwsHttps"
)

// "github.com/moovweb/gokogiri"
// "github.com/moovweb/gokogiri/xpath"

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

// ----------------------------------------------------------------------------
//                 Finds products
// ----------------------------------------------------------------------------

// getMatchingProduct get a product by asin
func getMatchingProductForASIN(asin string) *mwsHttps.Response {
	xmlFileResponse := createMwsHTTPResponse("examples/GetMatchingProductResponse.xml")
	if xmlFileResponse != nil {
		return xmlFileResponse
	}

	fmt.Println("-------- no example aws response ---------")
	fmt.Println("")
	productsClient, _ := products.NewClient(getConfigFile())
	response := productsClient.GetMatchingProduct([]string{asin})
	if response.Error != nil {
		fmt.Println(response.Error.Error())
		return nil
	}

	xmlNode, _ := gmws.GenerateXMLNode(response.Body)
	fmt.Println("------------------------ getMatchingProduct ------------------------")
	xmlNode.PrintXML()

	return response
}

// getMatchingProductForKeyword gets products by keyword
func getMatchingProductForKeyword(keyword string, itemCondition string) *mwsHttps.Response {
	// xmlFileResponse := createMwsHTTPResponse("examples/listmatchingproductsresponse.xml")
	// if xmlFileResponse != nil {
	// 	return xmlFileResponse
	// }

	// fmt.Println("-------- no example aws response ---------")
	// fmt.Println("")
	productsClient, _ := products.NewClient(getConfigFile())
	params := gmws.Parameters{
		"ItemCondition": itemCondition,
	}
	response := productsClient.ListMatchingProducts(keyword, params)
	if response.Error != nil {
		fmt.Println(response.Error.Error())
		return nil
	}

	return response
}

// ----------------------------------------------------------------------------
//                 Finds product prices
// ----------------------------------------------------------------------------

// getLowestPricedOffersForASIN  - 200 requests per hour
func getLowestPricedOffersForASIN(asin, itemCondition string) *mwsHttps.Response {
	// xmlFileResponse := createMwsHTTPResponse("examples/GetLowestPricedOffersForASINResponse.xml")
	// if xmlFileResponse != nil {
	// 	return xmlFileResponse
	// }

	// fmt.Println("-------- no example aws getLowestPricedOffersForASIN response ---------")
	// fmt.Println("")
	productsClient, _ := products.NewClient(getConfigFile())
	response := productsClient.GetLowestPricedOffersForASIN(asin, itemCondition)

	if response.Error != nil {
		fmt.Println("getLowestPricedOffersForASIN" + response.Error.Error())
		return response
	}
	// xmlNode, _ := gmws.GenerateXMLNode(response.Body)
	// fmt.Println("------------------------ getLowestPricedOffersForASIN ------------------------")
	// xmlNode.PrintXML()
	return response
}

// getLowestOfferListingsForASIN - 36000 requests per hour
func getLowestOfferListingsForASIN(asin, itemCondition string) *mwsHttps.Response {
	// xmlFileResponse := createMwsHTTPResponse("examples/GetLowestOfferListingsForASINResponse.xml")
	// if xmlFileResponse != nil {
	// 	return xmlFileResponse
	// }

	// fmt.Println("-------- no example aws getLowestOfferListingsForASIN response ---------")
	// fmt.Println("")
	productsClient, _ := products.NewClient(getConfigFile())
	optional := gmws.Parameters{
		"ItemCondition": itemCondition,
	}
	response := productsClient.GetLowestOfferListingsForASIN([]string{asin}, optional)
	if response.Error != nil {
		fmt.Println(response.Error.Error())
		return nil
	}
	// xmlNode, _ := gmws.GenerateXMLNode(response.Body)
	// fmt.Println("------------------------ getLowestOfferListingsForASIN ------------------------")
	// xmlNode.PrintXML()
	return response
}

// ------------------- MWS Test Helper --------------------------------

func createMwsHTTPResponse(fileLocation string) *mwsHttps.Response {
	xmlFile, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------")
	}

	if xmlFile != nil {
		response := new(mwsHttps.Response)
		defer xmlFile.Close()
		b, _ := ioutil.ReadAll(xmlFile)
		response.Body = b
		return response
	}
	return nil
}

func createErrorMwsHTTPResponse(fileLocation string) *mwsHttps.Response {
	xmlFile, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------")
	}

	if xmlFile != nil {
		response := new(mwsHttps.Response)
		defer xmlFile.Close()
		b, _ := ioutil.ReadAll(xmlFile)
		response.Body = b
		response.Error = errors.New("test")
		return response
	}
	return nil
}
