package main

import (
	"GoRepositories/Mongo"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mws/products"
	"github.com/svvu/gomws/mwsHttps"
	// "github.com/moovweb/gokogiri"
	// "github.com/moovweb/gokogiri/xpath"
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
	asin := "B00AVWKUJS"
	itemCondition := "New"
	matchProduct := parseMatchingProduct("B00AVWKUJS")
	if matchProduct == nil {
		fmt.Println("failed to get item information")
		return
	}
	// 36000 requests per hour
	lowestOffers := parseLowestOfferListingsForASIN("B00AVWKUJS")
	if lowestOffers == nil {
		fmt.Println("failed to load offers from bigger route")
		return
	} else {
		productData := mergeMatchWithLowestOfferListingsForASIN(matchProduct, lowestOffers)
		res2B1mws, _ := json.Marshal(productData)
		fmt.Println("item : ", string(res2B1mws))
		fmt.Println("--------------------------------------------------------------------------------------------------------------------------------- ")
	}

	// 200 requests per hour
	lowestPriceOffers := parseLowestPricedOffersForASIN(asin, itemCondition)
	if lowestPriceOffers == nil {
		fmt.Println("failed to load offers from bigger route")
		return
	} else {
		productData := mergeMatchWithLowestPricedOffersForASIN(matchProduct, lowestPriceOffers)
		res2B1mws, _ := json.Marshal(productData)
		fmt.Println("item : ", string(res2B1mws))
		fmt.Println("--------------------------------------------------------------------------------------------------------------------------------- ")
	}

}

func mergeMatchWithLowestPricedOffersForASIN(itemInfo *GetMatchingProductResponse, mws *GetLowestPricedOffersForASINResponse) []*ProductTracking {
	itemData := []*ProductTracking{}
	// ignoring the index
	for _, result := range mws.Results {
		for _, offer := range result.Offers {
			itemInfoResult := itemInfo.Results[0]
			newItem := NewProductTracking(result.ASIN)

			newItem.Domain = "www_amazon_com"
			newItem.Title = itemInfoResult.Product.ProductTitle

			newItem.Category = itemInfoResult.Product.ProductGroup
			newItem.Author = ""
			newItem.Condition = result.ItemCondition
			newItem.SubCondition = offer.ItemSubcondition

			newItem.PathName = "http://www.amazon.com/dp/" + newItem.Asin
			newItem.ImageURL = itemInfoResult.Product.ProductImage

			newItem.CurrencyCode = offer.ListingPriceCurrencyCode

			newItem.RegularAmount = offer.LandedPriceAmount
			newItem.SaleAmount = offer.ListingPriceAmount
			newItem.ShippingAmount = offer.ShippingPriceAmount

			if newItem.RegularAmount == 0 && newItem.SaleAmount > 0 {
				newItem.RegularAmount = newItem.SaleAmount
			}

			newItem.SalesRank = itemInfoResult.Product.SalesRank
			newItem.IsBuyBoxEligible = offer.IsBuyBoxWinner

			newItem.SellerFeedbackCount = offer.SellerFeedbackCount
			newItem.SellerPositiveFeedbackRating = offer.SellerPositiveFeedbackRating
			newItem.Count = result.TotalCount

			if offer.IsFulfilledByAmazon {
				newItem.Channel = "Amazon"
			} else {
				newItem.Channel = "Merchant"
			}

			newItem.IsSoldByAmazon = offer.IsFulfilledByAmazon

			newItem.PackageLength = itemInfoResult.Product.PackageLength
			newItem.PackageWidth = itemInfoResult.Product.PackageWidth
			newItem.PackageHeight = itemInfoResult.Product.PackageHeight
			newItem.PackageWeight = itemInfoResult.Product.PackageWeight

			itemData = append(itemData, newItem)
		}
	}
	return itemData
}

func mergeMatchWithLowestOfferListingsForASIN(itemInfo *GetMatchingProductResponse, mws *GetLowestOfferListingsForASINResponse) []*ProductTracking {
	itemData := []*ProductTracking{}
	// ignoring the index
	for _, result := range mws.Results {
		for _, offer := range result.Product.Offers {
			itemInfoResult := itemInfo.Results[0]
			newItem := NewProductTracking(itemInfoResult.ASIN)

			newItem.Domain = "www_amazon_com"
			newItem.Title = itemInfoResult.Product.ProductTitle

			newItem.Category = itemInfoResult.Product.ProductGroup
			newItem.Author = ""
			newItem.Condition = offer.ItemCondition
			newItem.SubCondition = offer.ItemSubcondition

			newItem.PathName = "http://www.amazon.com/dp/" + newItem.Asin
			newItem.ImageURL = itemInfoResult.Product.ProductImage

			newItem.CurrencyCode = offer.ListingPriceCurrencyCode

			newItem.RegularAmount = offer.LandedPriceAmount
			newItem.SaleAmount = offer.ListingPriceAmount
			newItem.ShippingAmount = offer.ShippingPriceAmount

			if newItem.RegularAmount == 0 && newItem.SaleAmount > 0 {
				newItem.RegularAmount = newItem.SaleAmount
			}

			newItem.SalesRank = itemInfoResult.Product.SalesRank
			newItem.IsBuyBoxEligible = false

			newItem.SellerFeedbackCount = offer.SellerFeedbackCount
			newItem.SellerPositiveFeedbackRating = offer.SellerPositiveFeedbackRating
			newItem.Count = offer.NumberOfOfferListingsConsidered
			newItem.Channel = offer.FulfillmentChannel
			newItem.IsSoldByAmazon = false

			newItem.PackageLength = itemInfoResult.Product.PackageLength
			newItem.PackageWidth = itemInfoResult.Product.PackageWidth
			newItem.PackageHeight = itemInfoResult.Product.PackageHeight
			newItem.PackageWeight = itemInfoResult.Product.PackageWeight

			itemData = append(itemData, newItem)
		}
	}
	return itemData
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

func parseLowestPricedOffersForASIN(asin, itemCondition string) *GetLowestPricedOffersForASINResponse {
	response := getLowestPricedOffersForASIN(asin, itemCondition)
	xmlParser := gmws.NewXMLParser(response)
	// Check whether or not API send back error message
	if xmlParser.HasError() {
		fmt.Println("error : ", xmlParser.GetError())
		return nil
	}
	mws := new(GetLowestPricedOffersForASINResponse)
	errMws := xmlParser.Parse(mws)
	if errMws != nil {
		fmt.Println("Error unmarshalling from errMws", errMws)
		return nil
	}

	return mws
}

// getLowestOfferListingsForASIN - 36000 requests per hour
func getLowestOfferListingsForASIN(asin, itemCondition string) *mwsHttps.Response {
	productsClient, _ := products.NewClient(getConfigFile())
	optional := gmws.Parameters{
		"ItemCondition": itemCondition,
	}
	response := productsClient.GetLowestOfferListingsForASIN([]string{asin}, optional)
	if response.Error != nil {
		fmt.Println(response.Error.Error())
		return nil
	}
	return response
}

func parseLowestOfferListingsForASIN(asin string) *GetLowestOfferListingsForASINResponse {
	response := getLowestOfferListingsForASIN(asin, "New")

	xmlParser := gmws.NewXMLParser(response)
	// Check whether or not API send back error message
	if xmlParser.HasError() {
		fmt.Println("error : ", xmlParser.GetError())
		return nil
	}

	mws := new(GetLowestOfferListingsForASINResponse)
	errMws := xmlParser.Parse(mws)
	if errMws != nil {
		fmt.Println("Error unmarshalling from errMws", errMws)
		return nil
	}
	return mws
}

func getMatchingProduct(asin string) *mwsHttps.Response {
	productsClient, _ := products.NewClient(getConfigFile())
	response := productsClient.GetMatchingProduct([]string{asin})
	if response.Error != nil {
		fmt.Println(response.Error.Error())
		return nil
	}

	return response
}

func parseMatchingProduct(asin string) *GetMatchingProductResponse {
	response := getMatchingProduct(asin)

	xmlParser := gmws.NewXMLParser(response)
	// Check whether or not API send back error message
	if xmlParser.HasError() {
		fmt.Println("error : ", xmlParser.GetError())
		return nil
	}

	gmp := new(GetMatchingProductResponse)
	errMws := xmlParser.Parse(gmp)
	if errMws != nil {
		fmt.Println("Error unmarshalling from errMws", errMws)
		return nil
	}
	return gmp
}

// GetLowestPricedOffersForASINResponse model
type GetLowestPricedOffersForASINResponse struct {
	XMLName xml.Name                             `xml:"GetLowestPricedOffersForASINResponse"`
	Results []GetLowestPricedOffersForASINResult `xml:"GetLowestPricedOffersForASINResult"`
}

// GetLowestPricedOffersForASINResult model
type GetLowestPricedOffersForASINResult struct {
	XMLName       xml.Name                                   `xml:"GetLowestPricedOffersForASINResult"`
	ASIN          string                                     `xml:"Identifier>ASIN"`
	Status        string                                     `xml:"status,attr"`
	MarketplaceID string                                     `xml:"Identifier>MarketplaceId"`
	ItemCondition string                                     `xml:"Identifier>ItemCondition"`
	TotalCount    int                                        `xml:"Summary>TotalOfferCount"`
	Offers        []GetLowestPricedOffersForASINResultOffers `xml:"Offers>Offer"`
}

// GetLowestPricedOffersForASINResultOffers model
type GetLowestPricedOffersForASINResultOffers struct {
	XMLName          xml.Name `xml:"Offer"`
	ItemSubcondition string   `xml:"SubCondition"`

	IsFulfilledByAmazon bool `xml:"IsFulfilledByAmazon"`
	IsBuyBoxWinner      bool `xml:"IsBuyBoxWinner"`

	SellerPositiveFeedbackRating string `xml:"SellerFeedbackRating>SellerPositiveFeedbackRating"`
	SellerFeedbackCount          int    `xml:"SellerFeedbackRating>FeedbackCount"`

	ListingPriceCurrencyCode string  `xml:"ListingPrice>CurrencyCode"`
	ListingPriceAmount       float64 `xml:"ListingPrice>Amount"`

	LandedPriceCurrencyCode string  `xml:"LandedPrice>CurrencyCode"`
	LandedPriceAmount       float64 `xml:"LandedPrice>Amount"`

	ShippingPriceCurrencyCode string  `xml:"Shipping>CurrencyCode"`
	ShippingPriceAmount       float64 `xml:"Shipping>Amount"`
}

// GetMatchingProductResponse model
type GetMatchingProductResponse struct {
	XMLName xml.Name                   `xml:"GetMatchingProductResponse"`
	Results []GetMatchingProductResult `xml:"GetMatchingProductResult"`
}

// GetMatchingProductResult model
type GetMatchingProductResult struct {
	XMLName xml.Name `xml:"GetMatchingProductResult"`
	ASIN    string   `xml:"ASIN,attr"`
	Status  string   `xml:"status,attr"`
	Product *GetMatchingProductResultProduct
}

// GetMatchingProductResultProduct model
type GetMatchingProductResultProduct struct {
	XMLName       xml.Name `xml:"Product"`
	MarketplaceID string   `xml:"Identifiers>MarketplaceASIN>MarketplaceId"`
	ASIN          string   `xml:"Identifiers>MarketplaceASIN>ASIN"`
	Binding       string   `xml:"AttributeSets>ItemAttributes>Binding"`
	Brand         string   `xml:"AttributeSets>ItemAttributes>Brand"`
	PackageHeight float64  `xml:"AttributeSets>ItemAttributes>PackageDimensions>Height"`
	PackageLength float64  `xml:"AttributeSets>ItemAttributes>PackageDimensions>Length"`
	PackageWidth  float64  `xml:"AttributeSets>ItemAttributes>PackageDimensions>Width"`
	PackageWeight float64  `xml:"AttributeSets>ItemAttributes>PackageDimensions>Weight"`
	ProductGroup  string   `xml:"AttributeSets>ItemAttributes>ProductGroup"`
	ProductImage  string   `xml:"AttributeSets>ItemAttributes>SmallImage>URL"`
	ProductTitle  string   `xml:"AttributeSets>ItemAttributes>Title"`
	SalesRank     int      `xml:"SalesRankings>SalesRank>Rank"`
}

// GetLowestOfferListingsForASINResponse model
type GetLowestOfferListingsForASINResponse struct {
	XMLName xml.Name                           `xml:"GetLowestOfferListingsForASINResponse"`
	Results []LowestOfferListingsForASINResult `xml:"GetLowestOfferListingsForASINResult"`
}

// LowestOfferListingsForASINResult model
type LowestOfferListingsForASINResult struct {
	XMLName xml.Name `xml:"GetLowestOfferListingsForASINResult"`
	ASIN    string   `xml:"ASIN,attr"`
	Status  string   `xml:"status,attr"`
	Product *LowestOfferListingsForASINResultProduct
}

// LowestOfferListingsForASINResultProduct model
type LowestOfferListingsForASINResultProduct struct {
	XMLName xml.Name                                `xml:"Product"`
	Offers  []LowestOfferListingsForASINResultOffer `xml:"LowestOfferListings>LowestOfferListing"`
}

// LowestOfferListingsForASINResultOffer model
type LowestOfferListingsForASINResultOffer struct {
	XMLName                         xml.Name `xml:"LowestOfferListing"`
	ItemCondition                   string   `xml:"Qualifiers>ItemCondition"`
	ItemSubcondition                string   `xml:"Qualifiers>ItemSubcondition"`
	FulfillmentChannel              string   `xml:"Qualifiers>FulfillmentChannel"`
	ShipsDomestically               string   `xml:"Qualifiers>ShipsDomestically"`
	SellerPositiveFeedbackRating    string   `xml:"Qualifiers>SellerPositiveFeedbackRating"`
	SellerFeedbackCount             int      `xml:"SellerFeedbackCount"`
	NumberOfOfferListingsConsidered int      `xml:"NumberOfOfferListingsConsidered"`
	LandedPriceCurrencyCode         string   `xml:"Price>LandedPrice>CurrencyCode"`
	LandedPriceAmount               float64  `xml:"Price>LandedPrice>Amount"`
	ListingPriceCurrencyCode        string   `xml:"Price>ListingPrice>CurrencyCode"`
	ListingPriceAmount              float64  `xml:"Price>ListingPrice>Amount"`
	ShippingPriceCurrencyCode       string   `xml:"Price>Shipping>CurrencyCode"`
	ShippingPriceAmount             float64  `xml:"Price>Shipping>Amount"`
}

// ProductTracking model
type ProductTracking struct {
	// (ReadOnly) Id of the document, created by system utilizing Mongo Bson Id
	ID string `json:"id"  bson:"_id" binding:"required"`
	// (ReadOnly) Date the document was created (UTC)
	DateCreated time.Time `json:"dateCreated" bson:"DateCreated" binding:"required"`
	// (ReadOnly) Date the document was modified (UTC)
	DateModified time.Time `json:"-" bson:"DateModified"`
	// (ReadOnly) What user modified the document
	UserModified string `json:"-" bson:"UserModified,omitempty"`
	// TenantID pertains to the tenant the document belongs to.  this is not serialized out
	TenantID string `json:"-" bson:"TenantId"`
	// Asin - Amazon Number
	Asin string `json:"asin" bson:"Asin"`
	// Upc Number
	Upc string `json:"UPC" bson:"Upc"`
	// Domain
	Domain string `json:"domain" bson:"Domain"`
	// Title
	Title string `json:"title" bson:"Title"`
	// Category
	Category string `json:"category" bson:"Category"`
	// Author
	Author string `json:"author" bson:"Author"`
	// Condition
	Condition string `json:"condition" bson:"Condition"`
	// SubCondition
	SubCondition string `json:"subCondition" bson:"SubCondition"`
	// PathName - URL Path to the item
	PathName string `json:"pathName" bson:"PathName"`
	// ImageUrl
	ImageURL string `json:"imageUrl" bson:"ImageUrl"`
	// CurrencyCode
	CurrencyCode string `json:"currencyCode" bson:"CurrencyCode"`
	// RegularAmount
	RegularAmount float64 `json:"regularPrice" bson:"RegularAmount"`
	// SaleAmount
	SaleAmount float64 `json:"salePrice" bson:"SaleAmount"`
	// ShippingAmount
	ShippingAmount float64 `json:"shippingPrice" bson:"ShippingAmount"`
	// SalesRank
	SalesRank int `json:"salesRank" bson:"SalesRank"`
	// SellerFeedbackCount
	SellerFeedbackCount int `json:"sellerFeedbackCount" bson:"SellerFeedbackCount"`
	// SellerPositiveFeedbackRating
	SellerPositiveFeedbackRating string `json:"sellerPositiveFeedbackRating" bson:"SellerPositiveFeedbackRating"`
	// Count - Count of items in inventory (for amazon this can be split by merchant vs amazon [channels])
	Count int `json:"count" bson:"Count"`
	// Channel
	Channel string `json:"channel" bson:"Channel"`
	// IsSoldByAmazon
	IsSoldByAmazon bool `json:"isSoldByAmazon" bson:"IsSoldByAmazon"`
	// IsBuyBoxEligible
	IsBuyBoxEligible bool `json:"isBuyBoxEligible" bson:"IsBuyBoxEligible"`
	// PackageLength
	PackageLength float64 `json:"packageLength" bson:"PackageLength"`
	// PackageWidth
	PackageWidth float64 `json:"packageWidth" bson:"PackageWidth"`
	// PackageHeight
	PackageHeight float64 `json:"packageHeight" bson:"PackageHeight"`
	// PackageWeight
	PackageWeight float64 `json:"packageWeight" bson:"PackageWeight"`
	// AmazonFees
	AmazonFees float64 `json:"amazonFees" bson:"_"`
}

// NewProductTracking gets a new object
func NewProductTracking(asin string) *ProductTracking {
	newItem := new(ProductTracking)
	newItem.InitProductTracking(asin)
	return newItem
}

// InitProductTracking set defaults
func (p *ProductTracking) InitProductTracking(asin string) {
	p.Asin = asin
	p.CurrencyCode = "USD"
	p.SalesRank = 9223372036854775807
	p.Channel = "Merchant"

	eventTime := time.Now().UTC()
	p.ID = Mongo.GetNewBsonIDString()
	p.DateCreated = eventTime
	p.DateModified = eventTime
}
