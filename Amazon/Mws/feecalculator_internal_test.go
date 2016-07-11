package Mws

import "testing"

// Small Standard Size Test
func TestSmallStandardMediaMerchant(t *testing.T) {
	expected := 2.1
	productItem := new(ProductTracking)
	productItem.Asin = "B0000058HX"
	productItem.Category = "Music"
	productItem.PackageLength = 5.6
	productItem.PackageWidth = 4.9
	productItem.PackageHeight = 0.4
	productItem.PackageWeight = 0.25
	productItem.Channel = "Merchant"
	productItem.RegularAmount = 4.99
	productItem.SaleAmount = 4.99

	actual := CalculateFees(productItem)
	if actual != expected {
		t.Errorf("expected: '%g', got: '%g'", expected, actual)
	}
}

func TestGettingAmazonFeeOptionsByNameReturnMotorCycle(t *testing.T) {
	expected := 12
	actual := getAmazonFeeOptions("MotorCycle")
	if actual.ReferralFeesPercent != expected {
		t.Errorf("expected: '%d', got: '%d'", expected, actual.ReferralFeesPercent)
	}
}

func TestGettingAmazonFeeOptionsByNameReturnDefaultNoOptionFoundReturnDefault(t *testing.T) {
	expected := 15
	actual := getAmazonFeeOptions("Toy11")
	if actual.ReferralFeesPercent != expected {
		t.Errorf("expected: '%d', got: '%d'", expected, actual.ReferralFeesPercent)
	}
}

func TestConvertDecimalToPercentageBelowOne(t *testing.T) {
	expected := float64(45)
	actual := convertDecimalToPercentage(.45)
	if actual != expected {
		t.Errorf("expected: '%g', got: '%g'", expected, actual)
	}
}

func TestConvertDecimalToPercentageAboveOne(t *testing.T) {
	expected := float64(45)
	actual := convertDecimalToPercentage(45)
	if actual != expected {
		t.Errorf("expected: '%g', got: '%g'", expected, actual)
	}
}

func TestGetAmazonPackageSize(t *testing.T) {
	expected1 := 1.14
	expected2 := 1.10
	expected3 := 1.05

	actual := getAmazonPackageSize(expected2, expected1, expected3)
	if len(actual) == 0 {
		t.Error("expected : list greater than zero")
	}

	if actual[0] != expected1 {
		t.Errorf("expected1: '%g', got: '%g'", expected1, actual[0])
	}

	if actual[1] != expected2 {
		t.Errorf("expected2: '%g', got: '%g'", expected2, actual[1])
	}

	if actual[2] != expected3 {
		t.Errorf("expected3: '%g', got: '%g'", expected3, actual[2])
	}
}
