package Mws

func getCommision(price float64, productGroup string, option amazonFeeOption) float64 {
	if productGroup == "CE" {
		return getCeCommission(price, productGroup, option)
	}

	referral := toFixed(price*option.ReferralFeesPercent/100, 2)
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
	referral := toFixed(price*float64(option.ReferralFeesPercent/100), 2)

	if option.MinReferralFees1 && referral < 1 {
		referral = 1
	}

	if option.MinReferralFees2 && referral < 2 {
		referral = 2
	}
	//	}
	return referral
}
