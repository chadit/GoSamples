package main

import (
	"fmt"
	"os"
	"os/exec"
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
	cmd := exec.Command("java", "-jar", "D:\\Projects\\GoWorkspace\\src\\GoSamples\\selenium\\selenium-server-standalone.jar")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	//cmd.Wait()

	runSeleniumTest()
	cmd.Process.Signal(os.Kill)
	fmt.Println("end")
}

func runSeleniumTest() {
	// FireFox driver without specific version
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, _ := selenium.NewRemote(caps, "")
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
