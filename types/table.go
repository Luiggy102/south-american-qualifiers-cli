package types

type Table struct {
	Positions []struct {
		Country  string `json:"country"`
		Position string `json:"position"`
	} `json:"results"`
}
