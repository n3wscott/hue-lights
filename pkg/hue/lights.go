package hue

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/n3wscott/hue-lights/pkg/apis/hue/v1"
)

func (h *Hue) FindLights() {
	glog.Info("Finding Lights")

	body, err := h.get("/api/{username}/lights")
	if err != nil {
		// handle error...
		glog.Error("Finding lights returned error: ", err)
		return
	}
	var lights v1.Lights
	json.Unmarshal(body, &lights)

	h.Lights = &lights
}

func (h *Hue) findLight(lightId string) (*v1.Light, error) {
	body, err := h.get("/api/{username}/lights/" + lightId)
	if err != nil {
		// handle error...
		glog.Error("Finding light "+lightId+" returned error: ", err)
		return nil, err
	}
	var light v1.Light
	json.Unmarshal(body, &light)

	glog.Info("Light: ", light)

	return &light, nil
}
