package Mws

import (
	"GoRepositories/Mongo"
	"time"
)

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

type lowestPricedOffersAttribute struct {
	MarketplaceID string `json:"-MarketplaceID"`
	ASIN          string `json:"-ASIN"`
	ItemCondition string `json:"-ItemCondition"`
}

type amazonFeeOption struct {
	ReferralFeesPercent  int
	VcfDomesticStandard  float64
	VcfDomesticExpedited float64
	VcfInternational     *float64
	MediaFG              bool
	VcfPerPound          *float64
	ShippingCreditID     int
	MinReferralFees1     bool
	MinReferralFees2     bool
}

func getDefaultamazonFeeOptions() map[string]amazonFeeOption {
	amazonitems := make(map[string]amazonFeeOption)
	amazonitems["Amazon Kindle"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Automotive Parts and Accessories"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.65, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Baby Products (excluding baby apparel)"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.65, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Books"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 1.35, VcfDomesticExpedited: 1.35, VcfInternational: createFloat64Pointer(1.35), MediaFG: true, VcfPerPound: nil, ShippingCreditID: 2, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Camera and Photo"] = amazonFeeOption{ReferralFeesPercent: 8, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Cell Phone Accessories"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Consumer Electronics"] = amazonFeeOption{ReferralFeesPercent: 8, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Home & Garden (including Pet Supplies)**"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Industrial & Scientific (including Food Service and Janitorial & Sanitation)"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Kindle Accessories"] = amazonFeeOption{ReferralFeesPercent: 25, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Office Products"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Personal Computers"] = amazonFeeOption{ReferralFeesPercent: 6, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Software & Computer Games"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 1.35, VcfDomesticExpedited: 1.35, VcfInternational: nil, MediaFG: true, VcfPerPound: nil, ShippingCreditID: 1, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Sporting Goods*"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Tires & Wheels"] = amazonFeeOption{ReferralFeesPercent: 10, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Tools & Home Improvement**"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Toys"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Video & DVD"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.80, VcfDomesticExpedited: 0.80, VcfInternational: nil, MediaFG: true, VcfPerPound: nil, ShippingCreditID: 5, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Video Game Consoles"] = amazonFeeOption{ReferralFeesPercent: 8, VcfDomesticStandard: 1.35, VcfDomesticExpedited: 1.35, VcfInternational: nil, MediaFG: true, VcfPerPound: nil, ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Watches"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.45, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Unlocked Cell Phones"] = amazonFeeOption{ReferralFeesPercent: 8, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Any Other Products"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Amazon Devices"] = amazonFeeOption{ReferralFeesPercent: 25, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Apparel"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Art"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Art and Craft Supply"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Audible"] = amazonFeeOption{ReferralFeesPercent: 25, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Automotive Al"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Baby Product"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: nil, ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Beauty"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["BISS"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Book"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Car Audio or Theater"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["CE"] = amazonFeeOption{ReferralFeesPercent: 8, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Classical"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Collectibles"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Digital Accessories 5"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Digital Device Accessory"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Digital Music Album"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Digital Software"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Digital Video Games"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["DVD"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 1.35, VcfDomesticExpedited: 1.35, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["eBook"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Entertainment Memorabilia"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Furniture"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Gift Card"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Gourmet"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["GPS or Navigation System"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Grocery"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Health and Beauty"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Hobby"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Home"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Home Improvement"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Home Theater"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Jewelry"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: true}
	amazonitems["Kitchen"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Lawn & Patio"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Lighting"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Luggage"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Major Appliances"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["MotorCycle"] = amazonFeeOption{ReferralFeesPercent: 12, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Movie"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Music"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 1.35, VcfDomesticExpedited: 1.35, VcfInternational: createFloat64Pointer(1.35), MediaFG: true, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Musical Instruments"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Network Media Player"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Office Electronics"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Office Product"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Pantry"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["PC Accessory"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Personal Computer"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Pet Products"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Photography"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Portable Audio Video"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Prescription Drugs"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Prestige Beauty"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Receiver or Amplifier"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Shoes"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Single Detail Page Misc"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Software"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Speakers"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Sports"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Television"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Toy"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["TV Series Episode Video on Demand"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.54, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Video"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Video Games"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 1.35, VcfDomesticExpedited: 1.35, VcfInternational: nil, MediaFG: true, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 1, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Watch"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: true}
	amazonitems["Wireless"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: true, MinReferralFees2: false}
	amazonitems["Wireless Phone"] = amazonFeeOption{ReferralFeesPercent: 8, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Wireless Phone Accessory"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}
	amazonitems["Wireless Plan Option"] = amazonFeeOption{ReferralFeesPercent: 15, VcfDomesticStandard: 0.45, VcfDomesticExpedited: 0.65, VcfInternational: nil, MediaFG: false, VcfPerPound: createFloat64Pointer(0.05), ShippingCreditID: 6, MinReferralFees1: false, MinReferralFees2: false}

	return amazonitems
}

func createFloat64Pointer(value float64) *float64 {
	newValue := value
	return &newValue
}

func getAmazonProductTypeNames() []string {
	return []string{
		"ACCESSORY",
		"ACCESSORY_OR_PART_OR_SUPPLY",
		"AMAZON_TABLET_ACCESSORY",
		"ANTENNA",
		"AUTO_ACCESSORY",
		"BATTERY",
		"CABLE_OR_ADAPTER",
		"CAMERA_OTHER_ACCESSORIES",
		"CAMERA_POWER_SUPPLY",
		"CE_ACCESSORY",
		"CE_CARRYING_CASE_OR_BAG",
		"COMPUTER_COOLING_DEVICE",
		"COMPUTER_INPUT_DEVICE",
		"COMPUTER_OUTPUT_DEVICE_ACCESSORY",
		"COMPUTER_SPEAKER",
		"COMPUTER_VIDEO_GAME_CONTOLLER",
		"CONSUMER_ELECTRONICS_PARTS",
		"DIGITAL_DEVICE_ACCESSORY",
		"ELECTRONIC_COMPONENT",
		"GPS_OR_NAVIGATION_ACCESSORY",
		"HOME_LIGHTING_ACCESSORY",
		"INSTRUMENT_PARTS_AND_ACCESSORIES",
		"MOTHERBOARD",
		"PHONE_ACCESSORY",
		"PORTABLE_AV_DEVICE",
		"SOUND_CARD",
		"SPEAKERS",
		"SYSTEM_POWER_DEVICE",
		"VIDEO_GAME_ACCESSORIES",
		"WIRELESS_ACCESSORY",
		"AUDIO_OR_VIDEO",
		"COMPUTER_COMPONENT",
		"COMPUTER_ADD_ON",
		"HEADPHONES",
		"RAM_MEMORY",
		"REMOTE_CONTROL",
	}
}
