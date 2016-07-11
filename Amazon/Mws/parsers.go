package Mws

import (
	"errors"
	"fmt"

	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

func setMwsKeys(amazonSellerID, amazonAuthToken, amazonRegion, amazonAccessKey, amazonSecretKey string) {
	sellerID = amazonSellerID
	authToken = amazonAuthToken
	if amazonRegion == "" {
		region = "US"
	} else {
		region = amazonRegion
	}
	accessKey = amazonAccessKey
	secretKey = amazonSecretKey
}

// ----------------------------------------------------------------------------------------------
//                                   Get Products by ASIN
// ----------------------------------------------------------------------------------------------

// GetProductsByASIN return an item from MWS by its ASIN
func GetProductsByASIN(asin, amazonSellerID, amazonAuthToken, amazonRegion, amazonAccessKey, amazonSecretKey string) (map[string]*[]ProductTracking, error) {
	setMwsKeys(amazonSellerID, amazonAuthToken, amazonRegion, amazonAccessKey, amazonSecretKey)
	productGroupings := make(map[string]*[]ProductTracking)
	// calls MWS to find product by keyword
	response := getMatchingProductForASIN(asin)
	if response == nil {
		return productGroupings, errors.New("no response from MWS ProductsByASIN")
	}

	// parse that response into an array of ProductTracking items
	return parseMatchingProductForASIN(response, "New")
}

func parseMatchingProductForASIN(response *mwsHttps.Response, itemCondition string) (map[string]*[]ProductTracking, error) {
	productGroupings := make(map[string]*[]ProductTracking)

	xmlNode, xmlNodeError := gmws.GenerateXMLNode(response.Body)
	if xmlNodeError != nil {
		return productGroupings, xmlNodeError
	}

	products := xmlNode.FindByPath("GetMatchingProductResponse.GetMatchingProductResult.Product")
	if len(products) == 0 {
		return productGroupings, errors.New("could not parse product")
	}

	for _, product := range products {
		itemAttributes := product.FindByPath("AttributeSets.ItemAttributes")
		if len(itemAttributes) == 0 {
			fmt.Println("could not parse itemAttributes")
			continue
		}

		itemMarketPlaceIdentifiers := product.FindByPath("Identifiers.MarketplaceASIN")
		if len(itemMarketPlaceIdentifiers) == 0 {
			fmt.Println("could not parse itemMarketPlaceIdentifiers")
			continue
		}

		asin := getXMLStringValue(itemMarketPlaceIdentifiers[0].FindByPath("ASIN"))
		newProductTracking := NewProductTracking(asin)
		newProductTracking.Domain = "www_amazon_com"
		newProductTracking.Title = getXMLStringValue(itemAttributes[0].FindByPath("Title"))
		newProductTracking.ImageURL = getXMLStringValue(itemAttributes[0].FindByPath("SmallImage.URL"))
		newProductTracking.Category = getXMLStringValue(itemAttributes[0].FindByPath("ProductGroup"))
		newProductTracking.PathName = "http://www.amazon.com/dp/" + asin

		salesRankings := product.FindByPath("SalesRankings.SalesRank.Rank")
		if len(salesRankings) > 0 {
			newProductTracking.SalesRank = getXMLIntValue((salesRankings))
		}

		itemDemensions := itemAttributes[0].FindByPath("PackageDimensions")
		if len(itemDemensions) > 0 {
			newProductTracking.PackageHeight = getProductDemensions(itemDemensions[0], "Height.#text")
			newProductTracking.PackageLength = getProductDemensions(itemDemensions[0], "Length.#text")
			newProductTracking.PackageWeight = getProductDemensions(itemDemensions[0], "Weight.#text")
			newProductTracking.PackageWidth = getProductDemensions(itemDemensions[0], "Width.#text")
		}

		productsWithPricing, productsWithPricingError := getPricingInformationForProduct(*newProductTracking, itemCondition)
		if productsWithPricingError == nil {
			productGroupings[asin] = productsWithPricing
		}
	}

	return productGroupings, nil
}

// ----------------------------------------------------------------------------------------------
//                                   Get Products by Keyword
// ----------------------------------------------------------------------------------------------

// GetProductsByKeyword return an item from MWS by its ASIN
func GetProductsByKeyword(keyword, itemCondition, amazonSellerID, amazonAuthToken, amazonRegion, amazonAccessKey, amazonSecretKey string) (map[string]*[]ProductTracking, error) {
	setMwsKeys(amazonSellerID, amazonAuthToken, amazonRegion, amazonAccessKey, amazonSecretKey)
	productGroupings := make(map[string]*[]ProductTracking)
	// calls MWS to find product by keyword
	response := getMatchingProductForKeyword(keyword, itemCondition)
	if response == nil {
		return productGroupings, errors.New("no response from MWS MatchProduct")
	}

	// parse that response into an array of ProductTracking items
	return parseMatchingProductForKeyword(response, itemCondition)
}

func parseMatchingProductForKeyword(response *mwsHttps.Response, itemCondition string) (map[string]*[]ProductTracking, error) {
	productGroupings := make(map[string]*[]ProductTracking)

	xmlNode, xmlNodeError := gmws.GenerateXMLNode(response.Body)
	if xmlNodeError != nil {
		return productGroupings, xmlNodeError
	}

	products := xmlNode.FindByPath("ListMatchingProductsResponse.ListMatchingProductsResult.Products.Product")
	if len(products) == 0 {
		return productGroupings, errors.New("could not parse product")
	}

	for _, product := range products {
		itemAttributes := product.FindByPath("AttributeSets.ItemAttributes")
		if len(itemAttributes) == 0 {
			fmt.Println("could not parse itemAttributes")
			continue
		}

		itemMarketPlaceIdentifiers := product.FindByPath("Identifiers.MarketplaceASIN")
		if len(itemMarketPlaceIdentifiers) == 0 {
			fmt.Println("could not parse itemMarketPlaceIdentifiers")
			continue
		}

		asin := getXMLStringValue(itemMarketPlaceIdentifiers[0].FindByPath("ASIN"))
		newProductTracking := NewProductTracking(asin)
		newProductTracking.Domain = "www_amazon_com"
		newProductTracking.Title = getXMLStringValue(itemAttributes[0].FindByPath("Title"))
		newProductTracking.ImageURL = getXMLStringValue(itemAttributes[0].FindByPath("SmallImage.URL"))
		newProductTracking.Category = getXMLStringValue(itemAttributes[0].FindByPath("ProductGroup"))
		newProductTracking.PathName = "http://www.amazon.com/dp/" + asin

		salesRankings := product.FindByPath("SalesRankings.SalesRank.Rank")
		if len(salesRankings) > 0 {
			newProductTracking.SalesRank = getXMLIntValue((salesRankings))
		}

		itemDemensions := itemAttributes[0].FindByPath("PackageDimensions")
		if len(itemDemensions) > 0 {
			newProductTracking.PackageHeight = getProductDemensions(itemDemensions[0], "Height.#text")
			newProductTracking.PackageLength = getProductDemensions(itemDemensions[0], "Length.#text")
			newProductTracking.PackageWeight = getProductDemensions(itemDemensions[0], "Weight.#text")
			newProductTracking.PackageWidth = getProductDemensions(itemDemensions[0], "Width.#text")
		}

		productsWithPricing, productsWithPricingError := getPricingInformationForProduct(*newProductTracking, itemCondition)
		if productsWithPricingError == nil {
			productGroupings[asin] = productsWithPricing
		}
	}

	return productGroupings, nil
}

// ----------------------------------------------------------------------------------------------
//                    Get Products pricing information and sales information
// ----------------------------------------------------------------------------------------------

func getPricingInformationForProduct(product ProductTracking, itemCondition string) (*[]ProductTracking, error) {
	var priceResponse *[]ProductTracking
	var priceResponseError error

	// call good route and nil is not returned
	// parse xml to get product pricing info
	priceResponse, priceResponseError = parseLowestPricedOffersForASIN(product.Asin, itemCondition, product)
	if priceResponseError == nil {
		return priceResponse, priceResponseError
	}
	// fmt.Println("Mws.parser.getPricingInformationForProduct : " + priceResponseError.Error())

	priceResponse, priceResponseError = parseLowestOfferListingsForASIN(product.Asin, itemCondition, product)
	if priceResponseError == nil {
		return priceResponse, priceResponseError
	}
	// parse xml to get product pricing info
	return priceResponse, errors.New("could not get prices")
}

// ----------------------------------------------------------------------------------------------
//      Get Products pricing information and sales information - LowsestPricedOffersForASIN
//                         Good Route 200 Request per hour Throttle
// ----------------------------------------------------------------------------------------------

// parseLowestPricedOffersForASIN good route -- 200 request per hour
func parseLowestPricedOffersForASIN(asin, itemCondition string, product ProductTracking) (*[]ProductTracking, error) {
	listProducts := &[]ProductTracking{}
	response := getLowestPricedOffersForASIN(asin, itemCondition)
	if response == nil {
		fmt.Println("parseLowestPricedOffersForASIN : no response ")
		return listProducts, errors.New("no response")
	}

	xmlNode, xmlNodeError := gmws.GenerateXMLNode(response.Body)
	if xmlNodeError != nil {
		return listProducts, xmlNodeError
	}

	if response.Error != nil {
		// fmt.Println("parseLowestPricedOffersForASIN : response error ", response.Error.Error())
		// xmlNode.PrintXML()
		// fmt.Println("--------------------------------------------")
		return listProducts, response.Error
	}

	// fmt.Println("parseLowestPricedOffersForASIN")
	offers := xmlNode.FindByPath("GetLowestPricedOffersForASINResponse.GetLowestPricedOffersForASINResult.Offers.Offer")
	if len(offers) == 0 {
		return listProducts, errors.New("could not parse offers")
	}

	headerNode := xmlNode.FindByPath("GetLowestPricedOffersForASINResponse.GetLowestPricedOffersForASINResult")
	if len(headerNode) == 0 {
		return listProducts, errors.New("could not parse headerNode")
	}
	newHeaderAttributes := lowestPricedOffersAttribute{}
	_ = headerNode[0].ToStruct(&newHeaderAttributes)

	offerCounts := getXMLIntValue(xmlNode.FindByPath("GetLowestPricedOffersForASINResponse.GetLowestPricedOffersForASINResult.Summary.TotalOfferCount"))

	for _, offer := range offers {
		newProduct := product
		newProduct.Channel = getChannel(getXMLBoolFromXMLNode(offer.FindByPath("IsFulfilledByAmazon")))
		newProduct.Count = offerCounts
		newProduct.Condition = newHeaderAttributes.ItemCondition
		newProduct.SubCondition = getXMLStringValue(offer.FindByPath("SubCondition"))
		newProduct.IsBuyBoxEligible = getXMLBoolFromXMLNode(offer.FindByPath("IsBuyBoxWinner"))
		newProduct.SellerFeedbackCount = getXMLIntValue(offer.FindByPath("SellerFeedbackRating.FeedbackCount"))
		newProduct.SellerPositiveFeedbackRating = getXMLStringValue(offer.FindByPath("SellerFeedbackRating.SellerPositiveFeedbackRating"))
		newProduct.SaleAmount = getXMLFloat64Value(offer.FindByPath("ListingPrice.Amount"))
		newProduct.RegularAmount = getXMLFloat64Value(offer.FindByPath("ListingPrice.Amount"))
		newProduct.ShippingAmount = getXMLFloat64Value(offer.FindByPath("Shipping.Amount"))
		newProduct.AmazonFees = CalculateFees(&newProduct)
		*listProducts = append(*listProducts, newProduct)
	}

	return listProducts, nil
}

// ----------------------------------------------------------------------------------------------
//      Get Products pricing information and sales information - LowestOfferListingsForASIN
//                         Good Route 36000 Request per hour Throttle
// ----------------------------------------------------------------------------------------------

// parseLowestOfferListingsForASIN good route -- 36000 request per hour
func parseLowestOfferListingsForASIN(asin, itemCondition string, product ProductTracking) (*[]ProductTracking, error) {
	listProducts := &[]ProductTracking{}
	response := getLowestOfferListingsForASIN(asin, itemCondition)
	if response == nil {
		fmt.Println("parseLowestOfferListingsForASIN : no response ")
		return listProducts, errors.New("no response")
	}

	xmlNode, xmlNodeError := gmws.GenerateXMLNode(response.Body)
	if xmlNodeError != nil {
		return listProducts, xmlNodeError
	}

	//	xmlNode.PrintXML()

	if response.Error != nil {
		// fmt.Println("parseLowestOfferListingsForASIN : response error ", response.Error.Error())
		// xmlNode.PrintXML()
		// fmt.Println("--------------------------------------------")
		return listProducts, response.Error
	}

	fmt.Println("parseLowestOfferListingsForASIN")
	offers := xmlNode.FindByPath("GetLowestOfferListingsForASINResponse.GetLowestOfferListingsForASINResult.Product.LowestOfferListings.LowestOfferListing")
	if len(offers) == 0 {
		return listProducts, errors.New("could not parse offers")
	}

	for _, offer := range offers {
		newProduct := product
		newProduct.Channel = getXMLStringValue(offer.FindByPath("Qualifiers.FulfillmentChannel"))
		newProduct.Count = getXMLIntValue(offer.FindByPath("NumberOfOfferListingsConsidered"))
		newProduct.Condition = getXMLStringValue(offer.FindByPath("Qualifiers.ItemCondition"))
		newProduct.SubCondition = getXMLStringValue(offer.FindByPath("Qualifiers.ItemSubcondition"))
		newProduct.SellerPositiveFeedbackRating = getXMLStringValue(offer.FindByPath("Qualifiers.SellerPositiveFeedbackRating"))
		newProduct.SellerFeedbackCount = getXMLIntValue(offer.FindByPath("SellerFeedbackCount"))

		newProduct.IsBuyBoxEligible = false
		newProduct.SellerFeedbackCount = getXMLIntValue(offer.FindByPath("SellerFeedbackRating.FeedbackCount"))

		newProduct.SaleAmount = getXMLFloat64Value(offer.FindByPath("Price.LandedPrice.Amount"))
		newProduct.RegularAmount = getXMLFloat64Value(offer.FindByPath("Price.ListingPrice.Amount"))
		newProduct.ShippingAmount = getXMLFloat64Value(offer.FindByPath("Price.Shipping.Amount"))
		newProduct.AmazonFees = CalculateFees(&newProduct)
		*listProducts = append(*listProducts, newProduct)
	}

	return listProducts, nil
}

func getChannel(isAmazon bool) string {
	if isAmazon {
		return "Amazon"
	}
	return "Merchant"
}
