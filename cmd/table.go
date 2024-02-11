package cmd

import (
	"fmt"

	t "github.com/Luigy102/south-american-qualifiers-cli/types"
	"github.com/fatih/color"
)

func PrintTable(t t.Table) {
	title := color.New(color.FgWhite, color.Bold).Add(color.Underline)
	title.Println("World Cup Qualifying - South America Table")

	fmt.Printf("%-2s %s \t%-3s %-3s %-3s %-3s %-3s %-3s \n", "", "Teams", "P", "W", "D", "L", "GD", "PTS")

	for _, v := range t.Positions {
		info := fmt.Sprintf("%-2s %s \t%-3s %-3s %-3s %-3s %-3s %-3s ",
			v.Position, v.Country, v.MatchesPlayes, v.Won, v.Draw, v.Losses, v.GoalDifference, v.Points)
		switch v.Position {
		case "1":
			color.Green(info)
		case "2":
			color.Green(info)
		case "3":
			color.Green(info)
		case "4":
			color.Green(info)
		case "5":
			color.Green(info)
		case "6":
			color.Green(info)
		case "7":
			color.White(info)
		case "8":
			color.Red(info)
		case "9":
			color.Red(info)
		case "10":
			color.Red(info)
		}
	}
}
