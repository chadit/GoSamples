package Mws

func getClosingFees(isInternational bool, isExpedited bool, packageWeight float64, option amazonFeeOption) float64 {
	fee := option.VcfDomesticStandard
	if isInternational && option.VcfInternational != nil {
		fee = *option.VcfInternational
	} else if isExpedited {
		fee = option.VcfDomesticExpedited
	}

	if option.VcfPerPound != nil {
		fee += *option.VcfPerPound * packageWeight
	}

	return fee
}
