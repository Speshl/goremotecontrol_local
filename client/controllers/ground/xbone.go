package ground

import (
	"fmt"
	"log"

	"github.com/Speshl/goremotecontrol_local/client/controllers"
	"github.com/Speshl/goremotecontrol_local/client/models"
	"gobot.io/x/gobot/v2/platforms/joystick"
)

var _ controllers.ControllerIFace[GroundControllerXBone, models.GroundControlData] = (*GroundControllerXBone)(nil)

type GroundControllerXBone struct {
	GroundControllerBase
	joystickAdaptor *joystick.Adaptor
	joystickDriver  *joystick.Driver
}

func NewGroundControllerXbone() *GroundControllerXBone {
	log.Printf("Creating Xbone Controller...")

	joystickAdaptor := joystick.NewAdaptor()
	joystick := joystick.NewDriver(joystickAdaptor, joystick.XboxOne)

	controller := GroundControllerXBone{
		joystickDriver:  joystick,
		joystickAdaptor: joystickAdaptor,
	}

	controller.LoadKeyMap()
	return &controller
}

func (c *GroundControllerXBone) LoadKeyMap() {
	//Not filled out because xbox controller input is not detected for some reason
	c.joystickDriver.On(c.joystickDriver.Event("start_press"), func(data interface{}) {
		fmt.Println("start_press")
	})

	// rt trigger
	c.joystickDriver.On("rt", func(data interface{}) {
		fmt.Println("rt", data)
	})
}

func (c *GroundControllerXBone) Start() error {
	log.Printf("Starting XBone Controller...")
	err := c.joystickAdaptor.Connect()
	if err != nil {
		return err
	}
	return c.joystickDriver.Start()
}

func (c *GroundControllerXBone) Stop() error {
	log.Printf("Stopping XBone Controller...")
	err := c.joystickDriver.Halt()
	if err != nil {
		return err
	}
	return c.joystickAdaptor.Finalize()
}
