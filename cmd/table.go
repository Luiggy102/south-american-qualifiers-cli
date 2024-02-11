package cmd

import (
	"fmt"

	t "github.com/Luigy102/south-american-qualifiers-cli/types"
	"github.com/fatih/color"
)

func PrintTable(t t.Table) {
	title := color.New(color.FgWhite, color.Bold).Add(color.Underline)
	title.Println("\nWorld Cup Qualifying - South America Table")

	fmt.Printf("%-2s %s \t%-3s %-3s %-3s %-3s %-3s %-3s \n", "", "Teams", "P", "W", "D", "L", "GD", "PTS")

	for _, v := range t.Positions {
		info := fmt.Sprintf("%-2s %s \t%-3s %-3s %-3s %-3s %-3s %-3s ",
			v.Position, v.Country, v.MatchesPlayes, v.Won, v.Draw, v.Losses, v.GoalDifference, v.Points)
		switch v.Position {
		case "1":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Green(info)
			case "La posición ha subido":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "2":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Green(info)
			case "La posición ha subido":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "3":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Green(info)
			case "La posición ha subido":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "4":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Green(info)
			case "La posición ha subido":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "5":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Green(info)
			case "La posición ha subido":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "6":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Green(info)
			case "La posición ha subido":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgGreen).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "7":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.White(info)
			case "La posición ha subido":
				color.New(color.FgWhite).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgWhite).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "8":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Red(info)
			case "La posición ha subido":
				color.New(color.FgRed).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgRed).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "9":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Red(info)
			case "La posición ha subido":
				color.New(color.FgRed).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgRed).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		case "10":
			switch v.Changes {
			case "No hay cambios en la posición":
				color.Red(info)
			case "La posición ha subido":
				color.New(color.FgRed).PrintfFunc()("%s", info)
				color.Cyan("↑")
			case "La posición ha bajado":
				color.New(color.FgRed).PrintfFunc()("%s", info)
				color.Cyan("↓")
			}
		}
	}
}
