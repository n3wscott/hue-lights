package hue

import (
	"github.com/golang/glog"
)

func (h *Hue) FindUser() {
	glog.Info("User Bridge")

	// TODO: need to save the user somehow... for now we are going to use return the default user.

	username := "unknown"
	if h.Bridge.Nunup.Id == "mock" {
		username = "newdeveloper"
	} else if h.Bridge.Nunup.Id == "001788fffe17f327" {
		username = "3x-swlnrBtlzmcxFejPUNdKQDdL0otT9DYzYn-j9"
	}

	h.User = &User{Username: &username}
}
