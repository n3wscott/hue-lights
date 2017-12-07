package v1

import "github.com/n3wscott/hue-lights/pkg/hue"

type Controller interface {
	GetPeople() ([]Person, error)

	GetPerson(id string) (*Person, error)
	CreatePerson(id string, person *CreatePersonRequest) (*Person, error)
	DeletePerson(id string) error

	GetBridge() (*hue.Bridge, error)
	GetUser() (*hue.User, error)

	GetCatalog() (*hue.Catalog, error)
}
