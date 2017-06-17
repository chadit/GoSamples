package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

// sudo ./example BB-E99F

func main() {

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	bb8 := bb8.NewDriver(bleAdaptor)
	fmt.Println("1")
	work := func() {
		gobot.Every(1*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			bb8.SetRGB(r, g, b)
		})
	}
	fmt.Println("2")
	robot := gobot.NewRobot("bb",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{bb8},
		work,
	)
	fmt.Println("3")
	err := robot.Start()
	fmt.Println(err)
}
