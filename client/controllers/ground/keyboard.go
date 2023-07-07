package ground

import (
	"log"
	"math"

	"github.com/Speshl/goremotecontrol_local/client/controllers"
	"github.com/Speshl/goremotecontrol_local/client/models"
	"gobot.io/x/gobot/v2/platforms/keyboard"
)

var _ controllers.ControllerIFace[GroundControllerKeyboard, models.GroundControlData] = (*GroundControllerKeyboard)(nil)

type GroundControllerKeyboard struct {
	GroundControllerBase
	keyboardDriver *keyboard.Driver
	keyMap         map[string]int
}

func NewGroundControllerKeyboard() *GroundControllerKeyboard {
	log.Printf("Creating Keyboard Controller...")
	controller := GroundControllerKeyboard{
		keyboardDriver: keyboard.NewDriver(),
	}

	controller.LoadKeyMap()

	controller.keyboardDriver.On(keyboard.Key, controller.keyHandler)

	return &controller
}

func (c *GroundControllerKeyboard) LoadKeyMap() {
	keyMap := make(map[string]int, 0)

	keyMap["THROTTLE"] = keyboard.W
	keyMap["BRAKE"] = keyboard.S
	keyMap["LEFT"] = keyboard.A
	keyMap["RIGHT"] = keyboard.D

	keyMap["PAN_LEFT"] = keyboard.Q
	keyMap["PAN_RIGHT"] = keyboard.E

	keyMap["TRIM_LEFT"] = keyboard.ArrowLeft
	keyMap["TRIM_RIGHT"] = keyboard.ArrowRight

	//TEMP KEYS
	keyMap["CENTER"] = keyboard.ArrowUp
	keyMap["COAST"] = keyboard.ArrowDown

	c.keyMap = keyMap
}

func (c *GroundControllerKeyboard) keyHandler(data interface{}) {
	key := data.(keyboard.KeyEvent)
	//log.Printf("Got Key Event...")
	for _, action := range GroundControlActions {
		//fmt.Printf("Comparing %d to %s value %d\n", key.Key, action, c.keyMap[action])
		if key.Key == c.keyMap[action] {
			//log.Printf("%s key pressed!\n", key.Char)
			c.PerformAction(action)
			log.Printf("ControlData: \n %+v\n\n", c.controlData)
			break
		}
	}
}

func (c *GroundControllerKeyboard) Start() error {
	log.Printf("Starting Keyboard Controller...")
	return c.keyboardDriver.Start()
}

func (c *GroundControllerKeyboard) Stop() error {
	log.Printf("Stopping Keyboard Controller...")
	return c.keyboardDriver.Halt()
}

func (c *GroundControllerKeyboard) PerformAction(action string) {
	//log.Printf("Performing Controller Action...")
	switch action {
	case "THROTTLE":
		c.controlData.ESC = math.MaxUint8
	case "BRAKE":
		c.controlData.ESC = 0
	case "COAST":
		c.controlData.ESC = math.MaxUint8 / 2

	case "LEFT":
		c.controlData.Steering = controllers.TrimInBounds(0, c.controlData.Trim)
	case "RIGHT":
		c.controlData.Steering = controllers.TrimInBounds(math.MaxUint8, c.controlData.Trim)
	case "CENTER":
		c.controlData.Steering = controllers.TrimInBounds(math.MaxUint8/2, c.controlData.Trim)
	case "TRIM_LEFT":
		c.controlData.Trim -= 1
	case "TRIM_RIGHT":
		c.controlData.Trim += 1

	case "PAN_LEFT":
		c.controlData.Pan = 0
	case "PAN_RIGHT":
		c.controlData.Pan = math.MaxUint8
	}

}
