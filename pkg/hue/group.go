package hue

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/n3wscott/hue-lights/pkg/apis/hue/v1"
)

func (h *Hue) FindGroups() {
	glog.Info("Get Groups")

	body, err := h.get("/api/{username}/groups")
	if err != nil {
		// handle error...
		glog.Error("Finding group returned error: ", err)
	}

	var groups v1.Groups

	json.Unmarshal(body, &groups)

	glog.Info("Groups: ", groups)

	//i := 0
	//for k, v := range groups {
	//	fmt.Printf("key[%s] value[%s]\n", k, v)
	//	h.Groups[i] = v
	//	i++
	//}

	h.Groups = &groups

	h.FindLights()
}
