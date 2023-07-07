package ground

import "github.com/Speshl/goremotecontrol_local/client/models"

type GroundControllerBase struct {
	controlData models.GroundControlData
}

//convert this to an enum
var GroundControlActions = []string{
	//ESC
	"THROTTLE",
	"BRAKE",
	"COAST",

	//STEERING
	"LEFT",
	"RIGHT",
	"CENTER",

	"TRIM_LEFT",
	"TRIM_RIGHT",

	//CAMERA
	"PAN_LEFT",
	"PAN_RIGHT",
}

func (c *GroundControllerBase) GetControlData() models.GroundControlData {
	return c.controlData
}
