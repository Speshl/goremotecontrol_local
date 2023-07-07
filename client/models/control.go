package models

type GroundControlData struct {
	ESCData
	SteeringData
	CameraData
}

type ESCData struct {
	ESC uint8
}

type SteeringData struct {
	Steering uint8
	Trim     int8
}

type CameraData struct {
	Pan  uint8
	Tilt uint8
}

func (d *GroundControlData) SetStartValues() {
	d.Steering = 127 //Default middle position
	d.ESC = 127      //Default middle position
}

func (d *GroundControlData) GetBytes() []byte {
	returnBytes := make([]byte, 4)
	returnBytes[0] = d.Steering
	returnBytes[1] = d.ESC
	returnBytes[2] = d.Pan
	returnBytes[3] = d.Tilt
	return returnBytes
}

func (d *SteeringData) SetScaledValue(value int16) {
	d.Steering = mapToRange(int(value), -32768, 32768, 0, 255)
}

func (d *SteeringData) LeftTrim() {
	if d.Trim > -60 {
		d.Trim--
	}
}

func (d *SteeringData) RightTrim() {
	if d.Trim < 60 {
		d.Trim++
	}
}

func (d *ESCData) SetScaledGasValue(value int16) {
	if value == -32768 { //Sometimes wraps around at full throttle
		return
	}

	d.ESC = mapToRange(int(value), -32768, 32768, 127, 255)
}

func (d *ESCData) SetScaledBrakeValue(value int16) {
	if value == -32768 { //Sometimes wraps around at full brake
		return
	}
	d.ESC = mapToRange(int(value), -32768, 32768, 127, 0)
}

func mapToRange(value int, min int, max int, minReturn int, maxReturn int) byte {
	return byte(int(maxReturn-minReturn)*(value-min)/(max-min) + int(minReturn))
}
