package Mws

import "testing"

func TestGetDefaultAmazonFeeCount(t *testing.T) {
	expected := 91
	defaultAmazonItems := getDefaultamazonFeeOptions()
	actual := len(defaultAmazonItems)
	if actual != expected {
		// t.Errorf("Test failed, expected: '%s', got:  '%s'", expectedString, actualString)
		t.Errorf("expected: '%d', got:  '%d'", expected, actual)
	}
}

func TestToyAmazonFeeOptions(t *testing.T) {
	expected := 15
	defaultAmazonItems := getDefaultamazonFeeOptions()
	actual := defaultAmazonItems["Toy"]
	if actual.ReferralFeesPercent != expected {
		t.Errorf("expected: '%d', got:  '%d'", expected, actual.ReferralFeesPercent)
	}
}
