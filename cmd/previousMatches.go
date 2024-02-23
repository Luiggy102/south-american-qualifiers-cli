package cmd

import (
	"fmt"

	m "github.com/Luigy102/south-american-qualifiers-cli/types"
	"github.com/fatih/color"
)

func ShowPrevious(t string, m m.PreviousMatches) {
	title := color.New(color.FgWhite, color.Bold).Add(color.Underline)
	dateFormat := "Mon 02 Jan 2006"

	// 3
	for _, v := range m.Matchday3 {
		title.Printf("\nMatchday 3 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	for _, v := range m.Matchday3 {
		if v.FirstTeam.Country == t || v.SecondTeam.Country == t {
			if v.Winner == v.FirstTeam.Country {
				// first team Winner
				fmt.Println(
					color.GreenString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
					fmt.Sprintf(" vs "),
					color.RedString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
				)
			} else if v.Winner == v.SecondTeam.Country {
				// seccond team Winner
				fmt.Println(
					color.RedString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
					fmt.Sprintf(" vs "),
					color.GreenString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
				)
			} else {
				// tie
				fmt.Printf("%s %s vs %s %s\n",
					v.FirstTeam.Country, v.FirstTeam.Goals,
					v.SecondTeam.Country, v.SecondTeam.Goals,
				)
			}
		} else {
			continue
		}
	}
	// 4
	for _, v := range m.Matchday4 {
		title.Printf("\nMatchday 4 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	for _, v := range m.Matchday4 {
		if v.FirstTeam.Country == t || v.SecondTeam.Country == t {
			if v.Winner == v.FirstTeam.Country {
				// first team Winner
				fmt.Println(
					color.GreenString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
					fmt.Sprintf(" vs "),
					color.RedString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
				)
			} else if v.Winner == v.SecondTeam.Country {
				// seccond team Winner
				fmt.Println(
					color.RedString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
					fmt.Sprintf(" vs "),
					color.GreenString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
				)
			} else {
				// tie
				fmt.Printf("%s %s vs %s %s\n",
					v.FirstTeam.Country, v.FirstTeam.Goals,
					v.SecondTeam.Country, v.SecondTeam.Goals,
				)
			}
		} else {
			continue
		}
	}
	// 5
	for _, v := range m.Matchday5 {
		title.Printf("\nMatchday 5 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	for _, v := range m.Matchday5 {
		if v.FirstTeam.Country == t || v.SecondTeam.Country == t {
			if v.Winner == v.FirstTeam.Country {
				// first team Winner
				fmt.Println(
					color.GreenString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
					fmt.Sprintf(" vs "),
					color.RedString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
				)
			} else if v.Winner == v.SecondTeam.Country {
				// seccond team Winner
				fmt.Println(
					color.RedString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
					fmt.Sprintf(" vs "),
					color.GreenString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
				)
			} else {
				// tie
				fmt.Printf("%s %s vs %s %s\n",
					v.FirstTeam.Country, v.FirstTeam.Goals,
					v.SecondTeam.Country, v.SecondTeam.Goals,
				)
			}
		} else {
			continue
		}
	}
	// 6
	for _, v := range m.Matchday6 {
		title.Printf("\nMatchday 6 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	for _, v := range m.Matchday6 {
		if v.FirstTeam.Country == t || v.SecondTeam.Country == t {
			if v.Winner == v.FirstTeam.Country {
				// first team Winner
				fmt.Println(
					color.GreenString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
					fmt.Sprintf(" vs "),
					color.RedString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
				)
			} else if v.Winner == v.SecondTeam.Country {
				// seccond team Winner
				fmt.Println(
					color.RedString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
					fmt.Sprintf(" vs "),
					color.GreenString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
				)
			} else {
				// tie
				fmt.Printf("%s %s vs %s %s\n",
					v.FirstTeam.Country, v.FirstTeam.Goals,
					v.SecondTeam.Country, v.SecondTeam.Goals,
				)
			}
		} else {
			continue
		}
	}
}

func Played(m m.PreviousMatches) {
	// show next three matchdays
	title := color.New(color.FgWhite, color.Bold).Add(color.Underline)
	dateFormat := "Mon 02 Jan 2006"

	// Print title and date
	for _, v := range m.Matchday3 {
		title.Printf("\nMatchday 3 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	// Print matches
	for _, v := range m.Matchday3 {
		if v.Winner == v.FirstTeam.Country {
			// first team Winner
			fmt.Println(
				color.GreenString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
				fmt.Sprintf(" vs "),
				color.RedString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
			)
		} else if v.Winner == v.SecondTeam.Country {
			// seccond team Winner
			fmt.Println(
				color.RedString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
				fmt.Sprintf(" vs "),
				color.GreenString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
			)
		} else {
			// tie
			fmt.Printf("%s %s vs %s %s\n",
				v.FirstTeam.Country, v.FirstTeam.Goals,
				v.SecondTeam.Country, v.SecondTeam.Goals,
			)
		}
	}

	// Print title and date
	for _, v := range m.Matchday4 {
		title.Printf("\nMatchday 4 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	// Print matches
	for _, v := range m.Matchday4 {
		if v.Winner == v.FirstTeam.Country {
			// first team Winner
			fmt.Println(
				color.GreenString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
				fmt.Sprintf(" vs "),
				color.RedString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
			)
		} else if v.Winner == v.SecondTeam.Country {
			// seccond team Winner
			fmt.Println(
				color.RedString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
				fmt.Sprintf(" vs "),
				color.GreenString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
			)
		} else {
			// tie
			fmt.Printf("%s %s vs %s %s\n",
				v.FirstTeam.Country, v.FirstTeam.Goals,
				v.SecondTeam.Country, v.SecondTeam.Goals,
			)
		}

	}

	// Print title and date
	for _, v := range m.Matchday5 {
		title.Printf("\nMatchday 5 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	// Print matches
	for _, v := range m.Matchday5 {
		if v.Winner == v.FirstTeam.Country {
			// first team Winner
			fmt.Println(
				color.GreenString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
				fmt.Sprintf(" vs "),
				color.RedString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
			)
		} else if v.Winner == v.SecondTeam.Country {
			// seccond team Winner
			fmt.Println(
				color.RedString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
				fmt.Sprintf(" vs "),
				color.GreenString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
			)
		} else {
			// tie
			fmt.Printf("%s %s vs %s %s\n",
				v.FirstTeam.Country, v.FirstTeam.Goals,
				v.SecondTeam.Country, v.SecondTeam.Goals,
			)
		}
	}

	// Print title and date
	for _, v := range m.Matchday6 {
		title.Printf("\nMatchday 6 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	// Print matches
	for _, v := range m.Matchday6 {
		if v.Winner == v.FirstTeam.Country {
			// first team Winner
			fmt.Println(
				color.GreenString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
				fmt.Sprintf(" vs "),
				color.RedString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
			)
		} else if v.Winner == v.SecondTeam.Country {
			// seccond team Winner
			fmt.Println(
				color.RedString("%s %s", v.FirstTeam.Country, v.FirstTeam.Goals),
				fmt.Sprintf(" vs "),
				color.GreenString("%s %s", v.SecondTeam.Country, v.SecondTeam.Goals),
			)
		} else {
			// tie
			fmt.Printf("%s %s vs %s %s\n",
				v.FirstTeam.Country, v.FirstTeam.Goals,
				v.SecondTeam.Country, v.SecondTeam.Goals,
			)
		}
	}

}
