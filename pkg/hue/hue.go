package hue

import "github.com/n3wscott/hue-lights/pkg/apis/hue/v1"

type Hue struct {
	Bridge *v1.Nunup
	User   *v1.User
	Groups *v1.Groups
	Lights *v1.Lights
}
