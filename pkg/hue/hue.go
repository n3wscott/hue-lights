package hue

type Hue struct {
	Bridge  *Bridge
	User    *User
	Groups  []Group
	Catalog *Catalog
}

type Bridge struct {
	Nunup *Nunup
}

type User struct {
	Username *string
}

type Catalog struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	Name     string   `json:"name"`
	LightIds []string `json:"lights"`
	Lights   []Light  `json:"lightObjects"`
	Type     string   `json:"type"`
	State    State    `json:"state"`
	Class    string   `json:"class"`
	Action   Action   `json:"action"`
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

type Nunup struct {
	Id                string `json:"id"`
	InternalIPAddress string `json:"internalipaddress"`
}
