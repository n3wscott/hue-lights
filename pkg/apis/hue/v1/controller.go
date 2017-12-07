package v1

import "github.com/n3wscott/hue-lights/pkg/hue"

type Controller interface {
	GetLights(username string) (*Lights, error)
	GetLight(lightId string) (*Light, error)
	//GetNewLights()

	GetGroups(username string) (*Groups, error)
	GetGroup(username string) (*Group, error)

	GetPerson(id string) (*Person, error)
	CreatePerson(id string, person *CreatePersonRequest) (*Person, error)
	DeletePerson(id string) error

	GetBridge() (*hue.Bridge, error)
	GetUser() (*hue.User, error)

	//GetCatalog() (*hue.Catalog, error)
}
