package hue

import (
	"encoding/json"

	"github.com/golang/glog"
)

func (h *Hue) FindLights() {
	glog.Info("Finding Lights")

	for i := range h.Groups {
		h.findLightsForGroup(&h.Groups[i])
		glog.Info(h.Groups[i])
	}
}

func (h *Hue) findLightsForGroup(group *Group) {
	for _, lightId := range group.LightIds {
		light, err := h.findLight(lightId)
		if err != nil {
			return
		}
		group.Lights = append(group.Lights, *light)
	}
}

func (h *Hue) findLight(lightId string) (*Light, error) {
	body, err := h.get("/api/{username}/lights/" + lightId)
	if err != nil {
		// handle error...
		glog.Error("Finding light "+lightId+" returned error: ", err)
		return nil, err
	}
	var light Light
	json.Unmarshal(body, &light)

	glog.Info("Light: ", light)

	return &light, nil
}
