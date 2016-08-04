package main

import (
	"fmt"
	"math"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

var code = `
package main
import "fmt"

func main() {
	fmt.Println("Hello111 WebDriver!\n")
}
`

func main() {
	fmt.Println("start")

	isServiceUp := isSeleniumServiceRunning()
	if isServiceUp {
		fmt.Println("service is running")
	} else {
		startupError := startSeleniumService()
		if startupError != nil {
			fmt.Println("failed to start selenium service : ", startupError)
			panic("boom")
		}
		isServiceUp = isSeleniumServiceRunning()
		if isServiceUp {
			fmt.Println("service is running")
		} else {
			fmt.Println("service is not running")
			panic("boom-boom")
		}
	}

	fmt.Println("end")
}

func startSeleniumService() error {
	cmd := exec.Command("java", "-jar", "D:\\Projects\\GoWorkspace\\src\\GoSamples\\selenium\\selenium-server-standalone.jar")
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func isSeleniumServiceRunning() bool {
	defer func() {
		if r := recover(); r != nil {
			// if selenium server is not running the driver will panic
		}
	}()

	wd := getWebDriver()
	defer wd.Quit()

	_, seleniumStatusError := wd.Status()
	if seleniumStatusError != nil {
		// if there is some other error that occured trying to start the driver and get the status
		return false
	}
	return true
}

func getWebDriver() selenium.WebDriver {
	//caps := selenium.Capabilities(map[string]interface{}{"browserName": "firefox"})
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "firefox"})
	wd, _ := selenium.NewRemote(caps, "")
	return wd
}

func getToysRUsTest() (string, WebSelector) {
	url := "http://www.toysrus.com/product/index.jsp?productId=24943206&cp=2255956.2273442.67716186.2255960.70114206.3808765&parentPage=family"
	// url := "http://www.tooooysrus.com/product/index.jsp?productId=24943206&cp=2255956.2273442.67716186.2255960.70114206.3808765&parentPage=family"
	webSelector := WebSelector{}
	webSelector.ItemTitle = []string{"divBADBADBAD#lTitle h1", "div#lTitle h1\t"}
	webSelector.RegularPrice = []string{"div#price ul li.retail.fl.withLP span"}
	webSelector.SalePrice = []string{"div111#price ul li.retail.fl.withLP span"}
	return url, webSelector
}

func getSamsClubTest() (string, WebSelector) {
	url := "http://www.samsclub.com/sams/utility-table-cherry-24-x-60/prod750132.ip?origin=home_page.rr1&campaign=rr&sn=MultiItemPersonalizedViewCP&"
	webSelector := WebSelector{}
	webSelector.ItemTitle = []string{"div#content-nonmember section.mainContainer.forItemPage.zoomImg.v2-pdp div.container section.baseContainer.clearfix div.span12.prodTitle.offset0 h1 span"}
	webSelector.RegularPrice = []string{"div#itemPageMoneyBox.span6.offset0.itemPageMb div.span6.offset0.pricingInfo ul.lgFont li"}
	webSelector.SalePrice = []string{"div111#price ul li.retail.fl.withLP span"}
	webSelector.UPC = []string{}
	return url, webSelector
}

func runScrappingTest() {
	wd := getWebDriver()
	defer wd.Quit()

	url, webSelector := getToysRUsTest()

	wd.Get(url)
	time.Sleep(1 * time.Second)

	var itemTitle string
	var regularPrice float64
	var salePrice float64
	var itemUPC string

	args := []interface{}{}

	for _, valueSelector := range webSelector.ItemTitle {
		results, resultsError := wd.ExecuteScript("return document.querySelector('"+cleanScript(valueSelector)+"').innerHTML;", args)
		if resultsError != nil {
			//	fmt.Println(resultsError)
			continue
		}
		if results != nil && results != "" {
			itemTitle = strings.TrimSpace(results.(string))
			break
		}
	}

	for _, valueSelector := range webSelector.RegularPrice {
		results, resultsError := wd.ExecuteScript("return document.querySelector('"+cleanScript(valueSelector)+"').innerHTML;", args)
		if resultsError != nil {
			//	fmt.Println(resultsError)
			continue
		}

		if results != nil && results != "" {
			regularPrice = cleanPriceValue(results)
			break
		}
	}

	for _, valueSelector := range webSelector.SalePrice {
		results, resultsError := wd.ExecuteScript("return document.querySelector('"+cleanScript(valueSelector)+"').innerHTML;", args)
		if resultsError != nil {
			//	fmt.Println(resultsError)
			continue
		}

		if results != nil && results != "" {
			salePrice = cleanPriceValue(results)
			break
		}
	}

	for _, valueSelector := range webSelector.UPC {
		results, resultsError := wd.ExecuteScript("return document.querySelector('"+cleanScript(valueSelector)+"').innerHTML;", args)
		if resultsError != nil {
			//	fmt.Println(resultsError)
			continue
		}
		if results != nil && results != "" {
			itemUPC = strings.TrimSpace(results.(string))
			break
		}
	}

	fmt.Println("Title :" + itemTitle)
	fmt.Println("resultRegularPrice : ", regularPrice)
	fmt.Println("salePriceSelector : ", salePrice)
	fmt.Println("itemUPC :" + itemUPC)
}

func cleanPriceValue(item interface{}) float64 {
	if item == "" {
		return 0
	}

	innerHTML := item.(string)

	regex, _ := regexp.Compile("[^0-9.]+")
	cleanScript1 := regex.ReplaceAllString(innerHTML, " ")
	//fmt.Println("cleanScript1 : ", cleanScript1)
	var prices []float64
	items := strings.Split(cleanScript1, " ")
	itemsSize := len(items)
	if itemsSize == 0 {
		return 0
	}
	for i := range items {
		if items[i] != "" {
			// check item to see if price contains a . for the cent amount
			if strings.Index(items[i], ".") != -1 {
				convertedPrice := convertStringToFloat(items[i])
				if convertedPrice > 0 && hasDecimalPlaces(1, convertedPrice) {
					return convertedPrice
				}
				prices = append(prices, convertedPrice)
			} else if i+1 < itemsSize && strings.Index(items[i+1], ".") == -1 && len(items[i+1]) == 2 {
				convertedPrice := convertStringToFloat(items[i] + "." + items[i+1])
				if convertedPrice > 0 && hasDecimalPlaces(1, convertedPrice) {
					return convertedPrice
				}
				prices = append(prices, convertedPrice)
			} else {
				convertedPrice := convertStringToFloat(items[i])
				if convertedPrice > 0 && hasDecimalPlaces(1, convertedPrice) {
					return convertedPrice
				}
				prices = append(prices, convertedPrice)
			}
		}
	}

	if len(prices) == 0 {
		return 0
	}

	_ = sort.Reverse(sort.Float64Slice(prices))
	return prices[0]
}

func convertStringToFloat(item string) float64 {
	convertedPrice, convertedPriceError := strconv.ParseFloat(item, 64)
	if convertedPriceError != nil {
		fmt.Println("convert error : ", convertedPriceError)
		return 0
	}
	return convertedPrice
}

func hasDecimalPlaces(i int, value float64) bool {
	valuef := value * float64(math.Pow(10.0, float64(i)))
	extra := valuef - float64(int(valuef))
	return extra != 0
}

func cleanScript(itemSelector string) string {
	regex, err := regexp.Compile("\t|\n|\r")
	if err == nil {
		return regex.ReplaceAllString(itemSelector, "")
	}
	return itemSelector
}

func runSeleniumTest() {
	// // FireFox driver without specific version
	// caps := selenium.Capabilities{"browserName": "firefox"}
	// wd, _ := selenium.NewRemote(caps, "")
	wd := getWebDriver()
	defer wd.Quit()

	// Get simple playground interface
	wd.Get("http://play.golang.org/?simple=1")

	// Enter code in textarea
	elem, _ := wd.FindElement(selenium.ByCSSSelector, "#code")
	elem.Clear()
	elem.SendKeys(code)

	// Click the run button
	btn, _ := wd.FindElement(selenium.ByCSSSelector, "#run")
	btn.Click()

	// Get the result
	div, _ := wd.FindElement(selenium.ByCSSSelector, "#output")

	output := ""
	// Wait for run to finish
	for {
		output, _ = div.Text()
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("Got: %s\n", output)
}

// WebSelector model
type WebSelector struct {
	ItemTitle    []string `json:"itemTitle"  bson:"ItemTitle" binding:"required"`
	SalePrice    []string `json:"salePrice"  bson:"SalePrice"`
	RegularPrice []string `json:"regularPrice"  bson:"RegularPrice"`
	UPC          []string `json:"UPC"  bson:"UPC"`
}
