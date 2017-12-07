package hue

import (
	"encoding/json"

	"fmt"

	"github.com/golang/glog"
)

func (h *Hue) FindGroups() {
	glog.Info("Get Groups")

	body, err := h.get("/api/{username}/groups")
	if err != nil {
		// handle error...
		glog.Error("Finding group returned error: ", err)
	}

	var groups map[string]Group

	json.Unmarshal(body, &groups)

	glog.Info("Groups: ", groups)

	h.Groups = make([]Group, len(groups))

	i := 0
	for k, v := range groups {
		fmt.Printf("key[%s] value[%s]\n", k, v)
		v.Lights = make([]Light, 0)
		h.Groups[i] = v
		i++
	}

	h.Catalog = &Catalog{Groups: h.Groups}

	h.FindLights()
}
