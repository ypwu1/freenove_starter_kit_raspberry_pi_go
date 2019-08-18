package examples

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"testing"
	"time"
)

func TestBlink(t *testing.T) {

	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "11")

	work := func() {
		gobot.Every(1*time.Second, func() {
			if err := led.Toggle(); err != nil {
				t.FailNow()
			}
			if led.State() {
				t.Log("led is on")
			} else {
				t.Log("led is off")
			}

		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	go robot.Start()

	time.Sleep(5 * time.Second)
}