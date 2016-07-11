package Mws

import (
	"math"
	"sort"
	"strings"
)

var (
	smallStandardSize = "Small Standard"
	largeStandardSize = "Large Standard"
	smallOverSize     = "Small Over"
	mediumOverSize    = "Medium Over"
	largeOverSize     = "Large Over"
	specialOverSize   = "Special Over"
	noSize            = ""
)

// CalculateFees gets the fee prices
func CalculateFees(productItem *ProductTracking) float64 {
	amazonFeeOption := getAmazonFeeOptions(productItem.Category)
	commissionFees := getCommision(productItem.RegularAmount, productItem.Category, amazonFeeOption)
	var vcfFee float64
	var fbaHandling float64
	var pickAndPack float64
	var weightHandlingFee float64
	var specialHandlingFee float64
	var storageFee float64
	var tvFee float64

	if amazonFeeOption.MediaFG {
		vcfFee = getClosingFees(false, false, productItem.PackageWeight, amazonFeeOption)
	}

	if caseInsensitiveEquals(productItem.Channel, "Amazon") || caseInsensitiveEquals(productItem.Channel, "AmazonPrice") || caseInsensitiveEquals(productItem.Channel, "Merchant") {
		amazonProductTier := getProductTierValue(productItem, amazonFeeOption)
		var greater float64
		if productItem.PackageHeight != 0 && productItem.PackageLength != 0 && productItem.PackageWidth != 0 && productItem.PackageWeight != 0 {
			unitVolume := productItem.PackageWidth * productItem.PackageLength * productItem.PackageHeight
			storageFee = getAmazonStorageFees(unitVolume, amazonProductTier)
			dimensionalWeight := unitVolume / 166
			if dimensionalWeight > productItem.PackageWeight {
				greater = dimensionalWeight
			} else {
				greater = productItem.PackageWeight
			}
			fbaHandling, pickAndPack, weightHandlingFee, specialHandlingFee = getFbaFees(amazonProductTier, amazonFeeOption.MediaFG, greater, productItem.PackageWeight, productItem.RegularAmount)
			// end package check
		}
	}
	return commissionFees + vcfFee + fbaHandling + pickAndPack + weightHandlingFee + storageFee + tvFee + specialHandlingFee
}

// getFbaFees - rules as of 18 Feb 2016
func getFbaFees(productTier int, isMedia bool, greater float64, unitWeight float64, price float64) (float64, float64, float64, float64) {
	var fbaHandling, pickAndPack, weightHandlingFee, specialHandlingFee, outboundWeight float64
	switch {
	// smallStandardSize
	case productTier == 1 && price < 300:
		// Order Handling
		if isMedia {
			fbaHandling = 0
		} else {
			fbaHandling = 1
		}
		// Pick and Pack
		pickAndPack = 1.06
		// Weight Handling
		weightHandlingFee = .5
		// largeStandardSize
	case productTier == 2 && price < 300:
		// Order Handling
		if isMedia {
			fbaHandling = 0
		} else {
			fbaHandling = 1
		}
		// Pick and Pack
		pickAndPack = 1.06

		// Weight Handling
		// set the item wieght value
		if isMedia {
			outboundWeight = math.Ceil(unitWeight + 0.125)
			if outboundWeight <= 1 {
				weightHandlingFee = 0.85
			} else if outboundWeight > 1 && outboundWeight <= 2 {
				weightHandlingFee = 1.24
			} else {
				weightHandlingFee = 1.24 + (outboundWeight-2)*.41
			}
		} else {
			// non-media
			if unitWeight > 1 {
				outboundWeight = math.Ceil(greater + 0.25)
			} else {
				outboundWeight = math.Ceil(unitWeight + 0.25)
			}
			if outboundWeight <= 1 {
				weightHandlingFee = 0.96
			} else if outboundWeight > 1 && outboundWeight <= 2 {
				weightHandlingFee = 1.95
			} else {
				weightHandlingFee = 1.95 + (outboundWeight-2)*0.39
			}
		}
		// smallOverSize
	case productTier == 3:
		// Order Handling
		fbaHandling = 0
		// pick and pack
		pickAndPack = 4.09
		// weight Handling
		outboundWeight = math.Ceil(greater + 1)
		if outboundWeight >= 0 && outboundWeight <= 2 {
			weightHandlingFee = 2.06
		} else {
			weightHandlingFee = 2.06 + (outboundWeight-2)*0.39
		}
		// mediumOverSize
	case productTier == 4:
		// order Handling
		fbaHandling = 0
		// pick and pack
		pickAndPack = 5.2
		// weight handling
		outboundWeight = math.Ceil(greater + 1)
		if outboundWeight >= 0 && outboundWeight <= 2 {
			weightHandlingFee = 2.73
		} else {
			weightHandlingFee = 2.73 + (outboundWeight-2)*0.39
		}
		// largeOverSize
	case productTier == 5:
		// order Handling
		fbaHandling = 0
		// pick and pack
		pickAndPack = 8.4
		// weight handling
		outboundWeight = math.Ceil(greater + 1)
		if outboundWeight >= 0 && outboundWeight <= 90 {
			weightHandlingFee = 63.98
		} else {
			weightHandlingFee = 63.98 + (outboundWeight-90)*0.8
		}
		// specialOverSize
	case productTier == 6:
		// order Handling
		fbaHandling = 0
		// pick and pack
		pickAndPack = 10.53
		// weight handling
		outboundWeight = math.Ceil(greater + 1)
		if outboundWeight >= 0 && outboundWeight <= 90 {
			weightHandlingFee = 124.58
		} else {
			weightHandlingFee = 124.58 + (outboundWeight-90)*0.92
		}
		// end switch
	}

	return fbaHandling, pickAndPack, weightHandlingFee, specialHandlingFee
}

func getAmazonFeeOptions(productGroup string) amazonFeeOption {
	allAmazonFeeOptions := getDefaultamazonFeeOptions()
	productGroupItem := allAmazonFeeOptions[productGroup]
	if productGroupItem.ReferralFeesPercent != 0 {
		return productGroupItem
	}
	return allAmazonFeeOptions["Any Other Products"]
}

func getCommision(price float64, productGroup string, option amazonFeeOption) float64 {
	if productGroup == "CE" {
		return getCeCommission(price, productGroup, option)
	}

	referral := toFixed(price*float64(option.ReferralFeesPercent), 2)
	if option.MinReferralFees1 && referral < 1 {
		referral = 1
	}
	if option.MinReferralFees2 && referral < 2 {
		referral = 2
	}

	if caseInsensitiveEquals(productGroup, "digital accessories 5") || caseInsensitiveEquals(productGroup, "digital device accessory") {
		return toFixed(price*.45, 2)
	}

	if caseInsensitiveEquals(productGroup, "gps or navigation system") || caseInsensitiveEquals(productGroup, "home theater") || caseInsensitiveEquals(productGroup, "major appliances") {
		if price <= 100 {
			referral = toFixed(price*.15, 2)
		} else {
			referral = toFixed(price*.08, 2)
		}
	}
	return referral
}

func getCeCommission(price float64, productGroup string, option amazonFeeOption) float64 {
	var referral float64
	for _, name := range getAmazonProductTypeNames() {
		if caseInsensitiveContains(name, productGroup) {
			referral = toFixed(price*.15, 2)
		} else {
			referral = toFixed(price*float64(option.ReferralFeesPercent), 2)
		}

		if option.MinReferralFees1 && referral < 1 {
			referral = 1
		}

		if option.MinReferralFees2 && referral < 2 {
			referral = 2
		}
	}
	return referral
}

func getClosingFees(isInternational bool, isExpedited bool, packageWeight float64, option amazonFeeOption) float64 {
	fee := option.VcfDomesticStandard
	if isInternational && option.VcfInternational != nil {
		fee = *option.VcfInternational
	}

	if isExpedited {
		fee = option.VcfDomesticExpedited
	}

	if option.VcfPerPound != nil {
		fee += *option.VcfPerPound * packageWeight
	}

	return fee
}

func getProductTierValue(product *ProductTracking, option amazonFeeOption) int {
	productTierOptions := getAmazonProductTiers()

	if product.PackageHeight == 0 && product.PackageLength == 0 && product.PackageWidth == 0 && product.PackageWeight == 0 {
		return productTierOptions[largeStandardSize]
	}

	productPackageSizes := getAmazonPackageSize(product.PackageHeight, product.PackageLength, product.PackageWidth)
	longest := productPackageSizes[0]
	median := productPackageSizes[1]
	shortest := productPackageSizes[2]

	//isSmallStandard
	if longest <= 15 && median <= 12 && shortest <= .75 {
		if product.PackageWeight <= 12 {
			return productTierOptions[smallStandardSize]
		}
		if option.MediaFG && product.PackageWeight <= 14 {
			return productTierOptions[smallStandardSize]
		}
	}

	// isLargeStandard
	if longest <= 18 && median <= 14 && shortest <= 8 && product.PackageWeight <= 20 {
		return productTierOptions[largeStandardSize]
	}

	var dimensionalWeightThrottle float64 = 5184
	unitVolume := product.PackageWidth * product.PackageLength * product.PackageHeight

	// check weight
	checkPoint := product.PackageWeight
	checkPointVolumn := unitVolume / 166
	if unitVolume > dimensionalWeightThrottle && checkPoint < checkPointVolumn {
		checkPoint = checkPointVolumn
	}

	// check lengh + girth
	lengthGirth := product.PackageLength + 2*(shortest+median)

	// isSmallOverSize
	if longest <= 60 && median <= 30 && checkPoint <= 70 && lengthGirth <= 130 {
		return productTierOptions[smallOverSize]
	}

	// isMediumOverSize
	if longest <= 108 && checkPoint <= 150 && lengthGirth <= 130 {
		return productTierOptions[mediumOverSize]
	}

	// isLargeOverSize
	if longest <= 108 && checkPoint <= 150 && lengthGirth <= 165 {
		return productTierOptions[largeOverSize]
	}

	// isSpecialOversize
	if longest > 108 || checkPoint > 150 || lengthGirth > 165 {
		return productTierOptions[specialOverSize]
	}

	return productTierOptions[noSize]
}

// helpers
func getAmazonProductTiers() map[string]int {
	productTiers := make(map[string]int)
	productTiers[noSize] = 0
	productTiers[smallStandardSize] = 1
	productTiers[largeStandardSize] = 2
	productTiers[smallOverSize] = 3
	productTiers[mediumOverSize] = 4
	productTiers[largeOverSize] = 5
	productTiers[specialOverSize] = 6
	return productTiers
}

func getAmazonPackageSize(packageHeight float64, packageLength float64, packageWidth float64) []float64 {
	size := []float64{packageHeight, packageLength, packageWidth}
	sort.Sort(sort.Reverse(sort.Float64Slice(size)))
	return size
}

func convertDecimalToPercentage(decimal float64) float64 {
	if decimal < 1 {
		return decimal * 100
	}
	return decimal
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func caseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func caseInsensitiveEquals(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return s == substr
}
