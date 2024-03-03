package cmd

import (
	"fmt"

	m "github.com/Luiggy102/south-american-qualifiers-cli/types"
	"github.com/fatih/color"
)

// function to show fixture of a specified team
// --fixture `team`
func TeamFixture(t string, m m.NextMatches) {
	// title formated
	title := color.New(color.FgWhite, color.Bold).Add(color.Underline)
	dateFormat := "Mon 02 Jan 2006"

	// 7
	for _, v := range m.Matchday7 {
		title.Printf("\nMatchday 7 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	for _, v := range m.Matchday7 {
		if v.FirstTeam == t || v.SecondTeam == t {
			fmt.Printf("%s vs %s\n",
				v.FirstTeam, v.SecondTeam)
		} else {
			continue
		}
	}

	// 8
	for _, v := range m.Matchday8 {
		title.Printf("\nMatchday 8 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	for _, v := range m.Matchday8 {
		if v.FirstTeam == t || v.SecondTeam == t {
			fmt.Printf("%s vs %s\n",
				v.FirstTeam, v.SecondTeam)
		} else {
			continue
		}
	}

	// 9
	for _, v := range m.Matchday9 {
		title.Printf("\nMatchday 9 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	for _, v := range m.Matchday9 {
		if v.FirstTeam == t || v.SecondTeam == t {
			fmt.Printf("%s vs %s\n",
				v.FirstTeam, v.SecondTeam)
		} else {
			continue
		}
	}
}

// function to display fixture
// --fixture
func ShowFixture(m m.NextMatches) {
	// show next three matchdays
	title := color.New(color.FgWhite, color.Bold).Add(color.Underline)
	dateFormat := "Mon 02 Jan 2006"

	// Print title and date
	for _, v := range m.Matchday7 {
		title.Printf("\nMatchday 7 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	// Print matches
	for _, v := range m.Matchday7 {
		// Teams
		fmt.Printf("%s vs %s\n",
			v.FirstTeam, v.SecondTeam)
	}

	// Print title and date
	for _, v := range m.Matchday8 {
		title.Printf("\nMatchday 8 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	// Print matches
	for _, v := range m.Matchday8 {
		// Teams
		fmt.Printf("%s vs %s\n",
			v.FirstTeam, v.SecondTeam)
	}

	// Print title and date
	for _, v := range m.Matchday9 {
		title.Printf("\nMatchday 9 - %v\n",
			v.Date.Format(dateFormat))
		break
	}
	// Print matches
	for _, v := range m.Matchday9 {
		// Teams
		fmt.Printf("%s vs %s\n",
			v.FirstTeam, v.SecondTeam)
	}

}
