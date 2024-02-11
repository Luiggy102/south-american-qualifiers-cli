package types

type Table struct {
	Positions []struct {
		Country        string `json:"country"`
		Position       string `json:"position"`
		Points         string `json:"points"`
		MatchesPlayes  string `json:"matches_played"`
		Won            string `json:"won"`
		Draw           string `json:"tied"`
		Losses         string `json:"losses"`
		Changes        string `json:"label"`
		GoalDifference string `json:"goal_difference"`
	} `json:"results"`
}
