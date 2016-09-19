package main

import (
	"fmt"
	"time"
)

func main() {
	parseDate()
	// 1469909149 = 7/30/2016 20:5:48
	//daysSince()
	//daysBetween1()
	//	convertIntToTime()
	// minusTwoHoursFromNow()
	//fiveDaysFromNow()
}

func parseDate() {
	s := "2017-09-18"
	v := "2006-01-02"
	ptime, ptimeErr := time.Parse(v, s)
	if ptimeErr != nil {
		fmt.Println(ptimeErr)
	}

	fmt.Println(ptime)
}

func daysSince() {
	// year, month, day, hour, seconds, millisonds, nanoseconds, format type
	longTimeAgo := time.Date(2016, time.January, 2, 20, 0, 0, 0, time.UTC)
	fmt.Println("TrialDate", longTimeAgo)

	hoursSince := int(time.Since(longTimeAgo).Hours())
	fmt.Println("hourSince : ", hoursSince)
	if hoursSince < 0 {
		// swap it and make it a positive
		hoursSince = hoursSince * -1
		modifier := 0
		if hoursSince >= 24 {
			modifier = int(hoursSince / 24)
			fmt.Println("modifier : ", modifier)
		}

		fmt.Println("trial is greater then : ", 30+modifier)
		return
	}

	daysForTrial := 30
	hoursForTrial := daysForTrial * 24
	hoursFromTrial := hoursForTrial - hoursSince

	fmt.Println("hoursForTrial : ", hoursForTrial)
	fmt.Println("hoursFromTrial : ", hoursFromTrial)

	fmt.Println("Days : ", (hoursForTrial-hoursSince)/24)
}

func convertIntToTime() {
	var epochTimeStamp int64
	epochTimeStamp = 1469909149
	timeInfo := time.Unix(epochTimeStamp, 0)
	//30 Jul 2016 20:05:49 GMT
	fmt.Println(timeInfo.UTC(), "7/30/2016 20:05:49")
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

func daysBetween1() {
	now := time.Now()

	fmt.Println("Today : ", now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	longTimeAgo := time.Date(2016, time.August, 31, 0, 0, 0, 0, time.UTC)

	// calculate the time different between today
	// and long time ago

	diff := now.Sub(longTimeAgo)

	// convert diff to days
	days := int(diff.Hours() / 24)

	fmt.Printf("31st July 2016 was %d days ago \n", days)
}
