package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Luigy102/south-american-qualifiers-cli/cmd"
	"github.com/Luigy102/south-american-qualifiers-cli/internal/requests"
	"github.com/Luigy102/south-american-qualifiers-cli/types"
)

func main() {

	printTable := flag.Bool("table", false, "display current table")
	showFixtures := flag.Bool("fixtures", false, "show next matches")
	teamFixture := flag.String("fixture", " ", "show team fixture")
	flag.Parse()

	info := `Welcome to south-american-qualifiers-cli!
type -h or --help for more info`

	// Print Position table
	if *printTable {
		var table types.Table = requests.Table()
		cmd.PrintTable(table)
		apiRequestedMessage()

		return
	}

	if *showFixtures {
		var nextMatches types.NextMatches = requests.NextMatches()
		cmd.ShowFixture(nextMatches)
		apiRequestedMessage()

		return
	}

	var teams = [10]string{"Argentina", "Bolivia", "Brasil", "Chile", "Colombia", "Ecuador", "Paraguay", "Perú", "Uruguay", "Venezuela"}

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
			apiRequestedMessage()
			return
		case " ":
			//
		default:
			fmt.Println("Invalid option")
		}
	}

	// Print Teams fixture

	// if *showFixture {
	// 	var nextMatches types.NextMatches = requests.NextMatches()
	// 	cmd.ShowFixture(nextMatches)
	// 	apiRequestedMessage()
	//
	// 	return
	// }

	fmt.Println(info)

}

func apiRequestedMessage() {
	fmt.Println("\nApi requested at:")
	fmt.Println(time.Now().Format("Monday 01-02-2006 15:04:05"))
}
