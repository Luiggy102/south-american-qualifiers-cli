package cmd

import (
	"fmt"

	m "github.com/Luigy102/south-american-qualifiers-cli/types"
	"github.com/fatih/color"
)

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
		fmt.Printf("* %s vs %s\n",
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
		fmt.Printf("* %s vs %s\n",
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
		fmt.Printf("* %s vs %s\n",
			v.FirstTeam, v.SecondTeam)
	}

}
