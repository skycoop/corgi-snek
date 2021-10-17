package types

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	Version    string `json:"version"`
}

type BattlesnakeMoveResponse struct {
	Move  Move   `json:"move"`
	Shout string `json:"shout,omitempty"`
}
