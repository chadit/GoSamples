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
	options := getAmazonFeeOptions(productItem.Category)
	commissionFees := getCommision(productItem.RegularAmount, productItem.Category, options)

	var vcfFee float64
	var fbaHandling float64
	var pickAndPack float64
	var weightHandlingFee float64
	var specialHandlingFee float64
	var storageFee float64
	var tvFee float64
	var greater float64

	if options.MediaFG {
		vcfFee = getClosingFees(false, false, productItem.PackageWeight, options)
	}

	if caseInsensitiveEquals(productItem.Channel, "Amazon") || caseInsensitiveEquals(productItem.Channel, "AmazonPrice") {
		amazonProductTier := getProductTierValue(productItem, options)
		// fmt.Println("productTier : ", amazonProductTier)
		// fmt.Println("Channel : ", productItem.Channel)

		if productItem.PackageHeight != 0 && productItem.PackageLength != 0 && productItem.PackageWidth != 0 && productItem.PackageWeight != 0 {
			unitVolume := productItem.PackageWidth * productItem.PackageLength * productItem.PackageHeight
			storageFee = getAmazonStorageFees(unitVolume, amazonProductTier)
			dimensionalWeight := unitVolume / 166
			if dimensionalWeight > productItem.PackageWeight {
				greater = dimensionalWeight
			} else {
				greater = productItem.PackageWeight
			}
			fbaHandling, pickAndPack, weightHandlingFee, specialHandlingFee = getFbaFees(amazonProductTier, options.MediaFG, greater, productItem.PackageWeight, productItem.RegularAmount)
			// end package check
		}
	}
	// if productItem.Asin == "B00AVWKUJS" {
	// 	fmt.Println("commissionFees (0.75) : ", commissionFees)
	// 	fmt.Println("vcfFee (1.35) : ", vcfFee)
	// 	fmt.Println("fbaHandling (0) : ", fbaHandling)
	// 	fmt.Println("pickAndPack (1.06) : ", pickAndPack)
	// 	fmt.Println("weightHandlingFee (0.5) : ", weightHandlingFee)
	// 	fmt.Println("storageFee (0) : ", storageFee)
	// 	fmt.Println("tvFee (0) : ", tvFee)
	// 	fmt.Println("specialHandlingFee (0) : ", specialHandlingFee)
	// 	fmt.Println("")
	// }
	return toFixed(commissionFees+vcfFee+fbaHandling+pickAndPack+weightHandlingFee+storageFee+tvFee+specialHandlingFee, 2)
}

func getAmazonFeeOptions(productGroup string) amazonFeeOption {
	allAmazonFeeOptions := getDefaultamazonFeeOptions()
	productGroupItem := allAmazonFeeOptions[productGroup]
	if productGroupItem.ReferralFeesPercent != 0 {
		return productGroupItem
	}
	return allAmazonFeeOptions["Any Other Products"]
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
