package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

func main() {
	gbot := gobot.NewGobot()

	adaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/tty.Sphero-BBY-AMP-SPP")
	spheroDriver := sphero.NewSpheroDriver(adaptor, "sphero")
	collConfig := sphero.DefaultCollisionConfig()
	collConfig.Xt = 10
	collConfig.Xs = 10
	collConfig.Yt = 10
	collConfig.Ys = 10

	locConfig := sphero.DefaultLocatorConfig()

	spheroDriver.ConfigureCollisionDetection(collConfig)
	spheroDriver.ConfigureLocator(locConfig)
	spheroDriver.SetRGB(0, 255, 0)
	direction := 0
	speed := uint8(110)
	zeroCount := 0

	pingpong := func() {

		gobot.Every(400*time.Millisecond, func() {
			vals := spheroDriver.ReadLocator()
			if len(vals) > 0 {
				fmt.Println("X:", vals[0], "Y:", vals[1], "XS:", vals[2], "YS:", vals[3], "SOG:", vals[4])
				if vals[4] < 5 && zeroCount <= 4 {
					zeroCount++
					spheroDriver.Roll(speed, uint16(direction))
				} else {
					spheroDriver.SetRGB(0, 255, 0)
				}
			} else {
				if zeroCount <= 4 {
					spheroDriver.Roll(speed, uint16(direction))
				}
				zeroCount++
			}
			if zeroCount > 2 && len(vals) > 0 {
				if vals[2] < 5 && vals[3] < 5 && vals[4] < 5 {
					if direction == 0 {
						direction = 180
					} else {
						direction = 0
					}
					log.Println("Barrier!")
				}
			}

			if zeroCount > 4 {
				spheroDriver.Stop()
				spheroDriver.SetRGB(255, 0, 0)
				if direction == 0 {
					direction = 180
				} else {
					direction = 0
				}
			}
			if zeroCount > 6 {
				zeroCount = 0
			}
		})

		gobot.On(spheroDriver.Event("collision"), func(data interface{}) {
			fmt.Printf("Collision Detected! %+v\n", data)
		})

	}

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{spheroDriver},
		pingpong,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
