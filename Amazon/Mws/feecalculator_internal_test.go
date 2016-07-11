package Mws

import "testing"

// start open testing -- used to test issues that pop up
func TestSmallStandardRandom(t *testing.T) {
	merchantExpected := 1.0
	expectedAmazon := 4.05
	productItem := new(ProductTracking)
	productItem.Asin = "B01739Y1KU"
	productItem.Category = "Toy"
	productItem.PackageLength = 8.7
	productItem.PackageWidth = 6.6
	productItem.PackageHeight = 1.5
	productItem.PackageWeight = 0.25
	productItem.Channel = "Merchant"
	productItem.RegularAmount = 5.99
	productItem.SaleAmount = 5.99

	actual := CalculateFees(productItem)
	if actual != merchantExpected {
		t.Errorf("merchantExpected: '%g', got: '%g'", merchantExpected, actual)
	}

	productItem.Channel = "Amazon"
	actual = CalculateFees(productItem)
	if actual != expectedAmazon {
		t.Errorf("expectedAmazon: '%g', got: '%g'", expectedAmazon, actual)
	}
}

// end open testing

// start Small Standard Size Test
func TestSmallStandardMedia(t *testing.T) {
	merchantExpected := 2.1
	expectedAmazon := 3.66
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
	if actual != merchantExpected {
		t.Errorf("merchantExpected: '%g', got: '%g'", merchantExpected, actual)
	}

	productItem.Channel = "Amazon"
	actual = CalculateFees(productItem)
	if actual != expectedAmazon {
		t.Errorf("expectedAmazon: '%g', got: '%g'", expectedAmazon, actual)
	}
}

func TestSmallStandardNoMedia(t *testing.T) {
	merchantExpected := 9.86
	expectedAmazon := 12.44
	productItem := new(ProductTracking)
	productItem.Asin = "B00AG0D5MO"
	productItem.Category = "Digital Device Accessory"
	productItem.PackageLength = 9.8
	productItem.PackageWidth = 9.5
	productItem.PackageHeight = 0.6
	productItem.PackageWeight = 0.4
	productItem.Channel = "Merchant"
	productItem.RegularAmount = 21.9
	productItem.SaleAmount = 21.9

	actual := CalculateFees(productItem)
	if actual != merchantExpected {
		t.Errorf("merchantExpected: '%g', got: '%g'", merchantExpected, actual)
	}

	productItem.Channel = "Amazon"
	actual = CalculateFees(productItem)
	if actual != expectedAmazon {
		t.Errorf("expectedAmazon: '%g', got: '%g'", expectedAmazon, actual)
	}
}

// end Small Standard Size Test

// start Small Over Size Test
func TestSmallOverSize(t *testing.T) {
	merchantExpected := 4.87
	expectedAmazon := 20.53
	productItem := new(ProductTracking)
	productItem.Asin = "B004I40SD8"
	productItem.Category = "Lawn & Patio"
	productItem.PackageLength = 55.0
	productItem.PackageWidth = 9.8
	productItem.PackageHeight = 6.9
	productItem.PackageWeight = 12.9
	productItem.Channel = "Merchant"
	productItem.RegularAmount = 32.47
	productItem.SaleAmount = 32.47

	actual := CalculateFees(productItem)
	if actual != merchantExpected {
		t.Errorf("merchantExpected: '%g', got: '%g'", merchantExpected, actual)
	}

	productItem.Channel = "Amazon"
	actual = CalculateFees(productItem)
	if actual != expectedAmazon {
		t.Errorf("expectedAmazon: '%g', got: '%g'", expectedAmazon, actual)
	}
}

func TestSmallOverSizeCEItem(t *testing.T) {
	merchantExpected := 6.94
	expectedAmazon := 26.48
	productItem := new(ProductTracking)
	productItem.Asin = "B00AVWKUJS"
	productItem.Category = "CE"
	productItem.PackageLength = 22.2
	productItem.PackageWidth = 19.6
	productItem.PackageHeight = 12.0
	productItem.PackageWeight = 30.55
	productItem.Channel = "Merchant"
	productItem.RegularAmount = 86.71
	productItem.SaleAmount = 86.71

	actual := CalculateFees(productItem)
	if actual != merchantExpected {
		t.Errorf("merchantExpected: '%g', got: '%g'", merchantExpected, actual)
	}

	productItem.Channel = "Amazon"
	actual = CalculateFees(productItem)
	if actual != expectedAmazon {
		t.Errorf("expectedAmazon: '%g', got: '%g'", expectedAmazon, actual)
	}
}

// end Small Over Size Test

// start large standard size testing
func TestLargeStandardSize(t *testing.T) {
	merchantExpected := 6.64
	expectedAmazon := 12.02
	productItem := new(ProductTracking)
	productItem.Asin = "B00GSN5H2O"
	productItem.Category = "Toy"
	productItem.PackageLength = 10.5
	productItem.PackageWidth = 10.2
	productItem.PackageHeight = 6.0
	productItem.PackageWeight = 1.25
	productItem.Channel = "Merchant"
	productItem.RegularAmount = 44.24
	productItem.SaleAmount = 44.24

	actual := CalculateFees(productItem)
	if actual != merchantExpected {
		t.Errorf("merchantExpected: '%g', got: '%g'", merchantExpected, actual)
	}

	productItem.Channel = "Amazon"
	actual = CalculateFees(productItem)
	if actual != expectedAmazon {
		t.Errorf("expectedAmazon: '%g', got: '%g'", expectedAmazon, actual)
	}
}

func TestLargeStandardSizeUnknownCategory(t *testing.T) {
	merchantExpected := 7.5
	expectedAmazon := 12.88
	productItem := new(ProductTracking)
	productItem.Asin = "B00P4L6DW4"
	productItem.Category = "Unknown"
	productItem.PackageLength = 12.0
	productItem.PackageWidth = 10.9
	productItem.PackageHeight = 4.8
	productItem.PackageWeight = 2.46
	productItem.Channel = "Merchant"
	productItem.RegularAmount = 50.0
	productItem.SaleAmount = 50.0

	actual := CalculateFees(productItem)
	if actual != merchantExpected {
		t.Errorf("merchantExpected: '%g', got: '%g'", merchantExpected, actual)
	}

	productItem.Channel = "Amazon"
	actual = CalculateFees(productItem)
	if actual != expectedAmazon {
		t.Errorf("expectedAmazon: '%g', got: '%g'", expectedAmazon, actual)
	}
}

// end large standard size testing

func TestGettingAmazonFeeOptionsByNameReturnMotorCycle(t *testing.T) {
	expected := 12.0
	actual := getAmazonFeeOptions("MotorCycle")
	if actual.ReferralFeesPercent != expected {
		t.Errorf("expected: '%g', got: '%g'", expected, actual.ReferralFeesPercent)
	}
}

func TestGettingAmazonFeeOptionsByNameReturnDefaultNoOptionFoundReturnDefault(t *testing.T) {
	expected := 15.0
	actual := getAmazonFeeOptions("Toy11")
	if actual.ReferralFeesPercent != expected {
		t.Errorf("expected: '%g', got: '%g'", expected, actual.ReferralFeesPercent)
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
