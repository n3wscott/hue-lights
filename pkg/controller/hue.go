package controller

import (
	"github.com/n3wscott/hue-lights/pkg/hue"
)

type HueController struct {
	Hue *hue.Hue
}

func (hc *HueController) GetBridge() (*hue.Nunup, error) {

	hc.Hue.FindBridge()

	return hc.Hue.Bridge.Nunup, nil
}
