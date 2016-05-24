package main

import (
	"fmt"
	"time"
)

func main() {
	minusTwoHoursFromNow()
	//fiveDaysFromNow()
}

func minusTwoHoursFromNow() {
	now := time.Now()

	fmt.Println("Today : ", now.Format("Mon, Jan 2, 2006 at 3:04pm"))
	dur, _ := time.ParseDuration("-2h")
	diff := now.Add(dur)

	fmt.Println("Two hours agao : ", diff.Format(time.ANSIC))
}

func fiveDaysFromNow() {
	now := time.Now()

	fmt.Println("Today : ", now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	fiveDays := time.Hour * 24 * 5

	diff := now.Add(fiveDays)

	fmt.Println("Five days from now will be : ", diff.Format(time.ANSIC))
}

func daysBetween() {
	now := time.Now()

	fmt.Println("Today : ", now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	longTimeAgo := time.Date(2010, time.May, 18, 23, 0, 0, 0, time.UTC)

	// calculate the time different between today
	// and long time ago

	diff := now.Sub(longTimeAgo)

	// convert diff to days
	days := int(diff.Hours() / 24)

	fmt.Printf("18th May 2010 was %d days ago \n", days)
}
