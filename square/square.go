package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

func main() {
	gbot := gobot.NewGobot()

	adaptor := sphero.NewSpheroAdaptor("sphero", "/dev/tty.Sphero-BBY-AMP-SPP")
	driver := sphero.NewSpheroDriver(adaptor, "sphero")
	driver.SetRGB(0, 0, 255)

	square := func() {
		direction := 0
		driver.SetRGB(0, 255, 0)

		gobot.Every(1*time.Second, func() {

			if direction == 360 {
				driver.SetRGB(255, 0, 0)

				driver.Stop()
			}
			driver.Roll(50, uint16(direction))
			direction = direction + 90
		})
	}

	escape := func() {

	}

	_ = escape
	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		square,
	)
	gbot.AddRobot(robot)
	gbot.Start()
}
