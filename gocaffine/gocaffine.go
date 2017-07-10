package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"

	"github.com/micmonay/keybd_event"
)

// to run on windows, need git installed and stty mapped to the path

func main() {
	// Create a channel to detect interruptions, so we can shutdown the different components appropiately.
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		keyboardListener()
	}()

	go func() {
		defer wg.Done()

		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}
		// For linux, it is very important wait 2 seconds
		if runtime.GOOS == "linux" {
			time.Sleep(2 * time.Second)
		}
		count := 0
		for {
			select {
			case <-exit:
				fmt.Println("keyboard loop ended : ")
				return
			default:
				time.Sleep(1 * time.Second)
				count++
				if count > 59 {
					count = 0
					if err := keyboardEvent(kb); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()

	wg.Wait()
	fmt.Println("ended")
}

func keyboardListener() {
	fmt.Println("keyboardListener started ")
	keys := keyboard.NewDriver()

	work := func() {

		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)
			fmt.Println("----------------")
			if key.Key == keyboard.A {
				fmt.Println("A pressed!")
			} else {
				fmt.Println("keyboard event!", key, key.Char)
			}
			fmt.Println("----------------")
		})
	}

	robot := gobot.NewRobot("keyboardbot",
		[]gobot.Connection{},
		[]gobot.Device{keys},
		work,
	)

	err := robot.Start()
	fmt.Println("keyboardListener ended : ", err)
}

func keyboardEvent(kb keybd_event.KeyBonding) error {
	//	fmt.Println("beep")
	kb.SetKeys(keybd_event.VK_F15)
	return kb.Launching()
}
