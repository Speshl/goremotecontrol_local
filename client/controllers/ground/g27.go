package ground

import (
	"log"

	"github.com/Speshl/goremotecontrol_local/client/controllers"
	"github.com/Speshl/goremotecontrol_local/client/models"
	"gobot.io/x/gobot/v2/platforms/joystick"
)

var _ controllers.ControllerIFace[GroundControllerG27, models.GroundControlData] = (*GroundControllerG27)(nil)

type GroundControllerG27 struct {
	GroundControllerBase
	joystickAdaptor *joystick.Adaptor
	joystickDriver  *joystick.Driver
}

func NewGroundControllerG27() *GroundControllerG27 {
	log.Printf("Creating G27 Controller...")

	joystickAdaptor := joystick.NewAdaptor()
	joystick := joystick.NewDriver(joystickAdaptor, joystick.LogitechG27)

	controller := GroundControllerG27{
		joystickDriver:  joystick,
		joystickAdaptor: joystickAdaptor,
	}

	controller.controlData.SetStartValues()

	controller.LoadKeyMap()
	return &controller
}

func (c *GroundControllerG27) LoadKeyMap() {
	c.joystickDriver.On("wheel", func(data interface{}) {
		c.controlData.SteeringData.SetScaledValue(data.(int16))
	})

	c.joystickDriver.On("gas", func(data interface{}) {
		c.controlData.ESCData.SetScaledGasValue(data.(int16) * -1)
	})

	c.joystickDriver.On("brake", func(data interface{}) {
		c.controlData.ESCData.SetScaledBrakeValue(data.(int16) * -1)
	})

	c.joystickDriver.On("wheel_left_top_press", func(data interface{}) {
		c.controlData.SteeringData.LeftTrim()
	})

	c.joystickDriver.On("wheel_right_top_press", func(data interface{}) {
		c.controlData.SteeringData.RightTrim()
	})
}

func (c *GroundControllerG27) Start() error {
	log.Printf("Starting G27 Controller...")
	err := c.joystickAdaptor.Connect()
	if err != nil {
		return err
	}
	return c.joystickDriver.Start()
}

func (c *GroundControllerG27) Stop() error {
	log.Printf("Stopping G27 Controller...")
	err := c.joystickDriver.Halt()
	if err != nil {
		return err
	}
	return c.joystickAdaptor.Finalize()
}
