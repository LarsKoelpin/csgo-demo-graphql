package domain

// Bomb is the structure representing the json of the bombstate in a tick.
type Bomb struct {
	LastOnGroundPosition Position `json:"lastOnGroundPosition"`
	Carrier              Player   `json:"carrier"`
}
