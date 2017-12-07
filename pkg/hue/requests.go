package hue

import (
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/golang/glog"
)

func (h *Hue) get(url string) ([]byte, error) {

	if strings.HasPrefix(url, "/") {
		url = h.Bridge.InternalIPAddress + url
	}

	if !strings.Contains(url, "http://") {
		url = "http://" + url
	}

	if strings.Contains(url, "{username}") {
		r := strings.NewReplacer("{username}", *h.User.Username)
		url = r.Replace(url)
	}

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		glog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	glog.Info(string(body))
	return []byte(body), nil
}
