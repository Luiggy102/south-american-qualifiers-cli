package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Luiggy102/south-american-qualifiers-cli/cmd"
	"github.com/Luiggy102/south-american-qualifiers-cli/internal/requests"
	"github.com/Luiggy102/south-american-qualifiers-cli/types"
)

func main() {

	// program flags
	// print positions table
	printTable := flag.Bool("table", false, "display current table")

	// boolean flags
	// show next matches
	showFixtures := flag.Bool("show-fixture", false, "show next matches")
	// show previous matches
	showPrevious := flag.Bool("show-previous", false, "show previous matches")

	// strings flags
	// show team next matches
	teamFixture := flag.String("fixture", " ", "show team fixture")
	// show team previous matches
	teamPrevious := flag.String("previous", " ", "show played matches for a team")
	flag.Parse()

	// boolean flags

	// Print Position table
	if *printTable {
		var table types.Table = requests.Table()
		cmd.PrintTable(table)
		apiRequestedInfo()

		return
	}

	// show all fixture
	if *showFixtures {
		var nextMatches types.NextMatches = requests.NextMatches()
		cmd.ShowFixture(nextMatches)
		apiRequestedInfo()
		return
	}

	// show all previous matches
	if *showPrevious {
		var prevMatches types.PreviousMatches = requests.PreviousMatches()
		cmd.Played(prevMatches)
		apiRequestedInfo()
		return
	}

	// strings flags

	var teams = [10]string{
		"Argentina",
		"Bolivia",
		"Brasil",
		"Chile",
		"Colombia",
		"Ecuador",
		"Paraguay",
		"Perú",
		"Uruguay",
		"Venezuela",
	}

	/// next matches
	for _, v := range teams {
		// Pass
		if v != *teamFixture {
			continue
		}
		// check the teams
		switch *teamFixture {
		case v:
			var nextMatches types.NextMatches = requests.NextMatches()
			cmd.TeamFixture(v, nextMatches)
			apiRequestedInfo()
			return
		case " ":
			//
		default:
			fmt.Println("Invalid option")
		}
	}

	/// played matches
	for _, v := range teams {
		// Pass
		if v != *teamPrevious {
			continue
		}
		// check the teams
		switch *teamPrevious {
		case v:
			var prevMatches types.PreviousMatches = requests.PreviousMatches()
			cmd.ShowPrevious(v, prevMatches)
			apiRequestedInfo()
			return
		case " ":
			//
		default:
			fmt.Println("Invalid option")
		}
	}

	// if none of the flags was used print this message
	info := `Welcome to south-american-qualifiers-cli!
type -h or --help for more info`
	fmt.Println(info)

}

func apiRequestedInfo() {
	fmt.Println("\nApi requested at:")
	fmt.Println(time.Now().Format("Monday 01-02-2006 15:04:05"))
}
