package hue

type Hue struct {
	Bridge *Bridge
}

type Bridge struct {
	Nunup *Nunup
}

type Nunup struct {
	Id                string `json:"id"`
	InternalIPAddress string `json:"internalipaddress"`
}
