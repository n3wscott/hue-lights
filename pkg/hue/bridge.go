package hue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

func (h *Hue) FindBridge() {
	glog.Info("Finding Bridge")

	resp, err := http.Get("https://www.meethue.com/api/nupnp")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	glog.Info("Bridge: ", string(body))

	nunups := make([]Nunup, 0)
	json.Unmarshal([]byte(body), &nunups)
	glog.Info("Obj: ", nunups)

	h.Bridge.Nunup = &nunups[0] // TODO check if there is one.
}
