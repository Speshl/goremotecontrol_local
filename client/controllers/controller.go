package controllers

import "math"

type ControllerIFace[ControllerType any, DataType any] interface {
	GetControlData() DataType
	Start() error
	Stop() error
}

func TrimInBounds(value uint8, trimValue int8) uint8 {
	if int(value)+int(trimValue) > math.MaxUint8 {
		return math.MaxUint8
	} else if int(value)+int(trimValue) < 0 {
		return 0
	} else {
		return value + uint8(trimValue)
	}
}
