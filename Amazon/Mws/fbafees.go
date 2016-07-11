package Mws

import "math"

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
