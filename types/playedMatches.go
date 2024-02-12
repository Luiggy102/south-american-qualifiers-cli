package types

import "time"

type Played struct {
	Matchday6 []struct {
		FirstTeam struct {
			Country string `json:"country"`
			Goals   string `json:"goals"`
		} `json:"first_team"`
		SecondTeam struct {
			Country string `json:"country"`
			Goals   string `json:"goals"`
		} `json:"second_team"`
		Winner string    `json:"winner"`
		Date   time.Time `json:"date"`
	} `json:"Jornada 6"`
	Matchday5 []struct {
		FirstTeam struct {
			Country string `json:"country"`
			Goals   string `json:"goals"`
		} `json:"first_team"`
		SecondTeam struct {
			Country string `json:"country"`
			Goals   string `json:"goals"`
		} `json:"second_team"`
		Winner string    `json:"winner"`
		Date   time.Time `json:"date"`
	} `json:"Jornada 5"`
	Matchday4 []struct {
		FirstTeam struct {
			Country string `json:"country"`
			Goals   string `json:"goals"`
		} `json:"first_team"`
		SecondTeam struct {
			Country string `json:"country"`
			Goals   string `json:"goals"`
		} `json:"second_team"`
		Winner string    `json:"winner"`
		Date   time.Time `json:"date"`
	} `json:"Jornada 4"`
	Matchday3 []struct {
		FirstTeam struct {
			Country string `json:"country"`
			Goals   string `json:"goals"`
		} `json:"first_team"`
		SecondTeam struct {
			Country string `json:"country"`
			Goals   string `json:"goals"`
		} `json:"second_team"`
		Winner string    `json:"winner"`
		Date   time.Time `json:"date"`
	} `json:"Jornada 3"`
}
