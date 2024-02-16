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
	showFixture := flag.Bool("fixture", false, "show next matches")
	flag.Parse()

	info := `Welcome to south-american-qualifiers-cli!
type -h or --help for more info`

	if *printTable {
		var table types.Table = requests.Table()
		cmd.PrintTable(table)
		apiRequestedMessage()

		return
	}

	if *showFixture {
		var nextMatches types.NextMatches = requests.NextMatches()
		cmd.ShowFixture(nextMatches)
		apiRequestedMessage()

		return
	}

	fmt.Println(info)

}

func apiRequestedMessage() {
	fmt.Println("\nApi requested at:")
	fmt.Println(time.Now().Format("Monday 01-02-2006 15:04:05"))
}
