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
		glog.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	glog.Info("Bridge: ", string(body))

	nunups := []Nunup{}
	json.Unmarshal([]byte(body), &nunups)
	glog.Info("Obj: ", nunups)

	if len(nunups) > 0 {
		h.Bridge.Nunup = &nunups[0]
	} else {
		h.Bridge.Nunup = &Nunup{Id: "mock", InternalIPAddress: "localhost:5000"}
	}
}
