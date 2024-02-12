package types

import "time"

type NextMatches struct {
	Matchday7 []struct {
		FirstTeam  string    `json:"first_team"`
		SecondTeam string    `json:"second_team"`
		Date       time.Time `json:"date"`
	} `json:"Jornada 7"`
	Matchday8 []struct {
		FirstTeam  string    `json:"first_team"`
		SecondTeam string    `json:"second_team"`
		Date       time.Time `json:"date"`
	} `json:"Jornada 8"`
	Matchday9 []struct {
		FirstTeam  string    `json:"first_team"`
		SecondTeam string    `json:"second_team"`
		Date       time.Time `json:"date"`
	} `json:"Jornada 9"`
}
