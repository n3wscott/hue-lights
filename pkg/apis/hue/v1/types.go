package v1

type Nunup struct {
	Id                string `json:"id"`
	InternalIPAddress string `json:"internalipaddress"`
}

// Post: /api
type UserRequest struct {
	DeviceType string `json:"devicetype"` // <application_name>#<devicename>
}

type UserResponse []UserSuccess

type UserSuccess struct {
	Success *User `json:"success"`
}

type User struct {
	Username *string `json:"username"`
}

// Get: /api/{username}/config
type Config struct {
	// TODO
	Username *string `json:"username"`
}

type LightId string

// Get: /api/{username}/lights
type Lights map[LightId]Light

// Get: /api/{username}/lights/{lightId}
type Light struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	State             State  `json:"state"`
	ModelId           string `json:"modelid"`
	UniqueId          string `json:"unitueid"`
	ManufacturerName  string `json:"manufacturername"`
	LuminaireUniqueId string `json:"luminaireuniqueid"`
	SWVersion         string `json:"swversion"`
}

type GroupId string

// Get: /api/{username}/groups
type Groups map[GroupId]Group

// Get: /api/{username}/groups/{groupId}
type Group struct {
	Name   string    `json:"name"`
	Lights []LightId `json:"lights"`
	Type   string    `json:"type"`
	State  State     `json:"state"`
	Class  string    `json:"class"`
	Action Action    `json:"action"`
}

type Action struct {
	On               bool   `json:"on"`
	Brightness       int    `json:"bri"`
	Hue              int    `json:"hue"`
	Saturation       int    `json:"sat"`
	XY               []int  `json:"xy"`
	ColorTemperature int    `json:"ct"`
	Alert            string `json:"alert"`
	Effect           string `json:"effect"`
	ColorMode        string `json:"colormode"`
}

type State struct {
	On               bool   `json:"on"`
	Brightness       int    `json:"bri"`
	Hue              int    `json:"hue"`
	Saturation       int    `json:"sat"`
	XY               []int  `json:"xy"`
	ColorTemperature int    `json:"ct"`
	Alert            string `json:"alert"`
	Effect           string `json:"effect"`
	ColorMode        string `json:"colormode"`
	Reachable        bool   `json:"reachable"`
}

type JsonMap map[string]interface{}

// Get: /api/{username}
type DataStore struct {
	Lights    Lights  `json:"lights"`
	Groups    Groups  `json:"groups"`
	Config    Config  `json:"config"`
	Schedules JsonMap `json:"schedules"`
	Scenes    JsonMap `json:"scenes"`
	Sensors   JsonMap `json:"sensors"`
	Rules     JsonMap `json:"rules"`
}
