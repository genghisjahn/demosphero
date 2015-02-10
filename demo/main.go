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
	driver.SetRGB(0, 255, 0)
	b := false
	s := 1

	work := func() {
		sp := 100
		d := 180
		gobot.Every(1*time.Second, func() {

			if s == 1 {
				d = 0
				sp = 100
			}
			if s == 2 {
				d = 90
				sp = 50
			}
			if s == 3 {
				d = 180
				sp = 100
			}
			if s == 4 {
				d = 270
				sp = 50
			}
			s = s + 1
			if s > 4 {
				s = 1
			}
			driver.Roll(uint8(sp), uint16(d))

			b = !b
			if b == false {
				driver.SetRGB(255, 0, 0)
			} else {
				driver.SetRGB(0, 255, 0)
			}

		})
	}
	_ = work
	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
