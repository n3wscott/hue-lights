package hue

import (
	"github.com/golang/glog"
	"github.com/n3wscott/hue-lights/pkg/apis/hue/v1"
)

func (h *Hue) FindUser() {
	glog.Info("User Bridge")

	// TODO: need to save the user somehow... for now we are going to use return the default user.

	username := "unknown"
	if h.Bridge.Id == "mock" {
		username = "newdeveloper"
	} else if h.Bridge.Id == "001788fffe17f327" {
		username = "3x-swlnrBtlzmcxFejPUNdKQDdL0otT9DYzYn-j9"
	}

	h.User = &v1.User{Username: &username}
}
