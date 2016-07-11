package Mws

import "time"

func getAmazonStorageFees(unitVolume float64, productTier int) float64 {
	var multiplier float64 = 1
	currentDateTime := time.Now().UTC()
	currentMonth := currentDateTime.Month()
	// standard small/large staorage fees
	if productTier == 1 || productTier == 2 {
		if currentMonth > 0 && currentMonth < 10 {
			multiplier = 0.54
		} else {
			multiplier = 0.72
		}
	} else {
		// everything else
		if currentMonth > 0 && currentMonth < 10 {
			multiplier = 0.43
		} else {
			multiplier = 0.57
		}
	}

	return toFixed((unitVolume*0.0005787)*multiplier, 2)
}
