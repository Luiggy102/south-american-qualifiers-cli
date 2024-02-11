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
	flag.Parse()

	if *printTable {
		var table types.Table = requests.Table()
		cmd.PrintTable(table)
		fmt.Println("\nUpdated at:")
		fmt.Println(time.Now().Format("Monday 01-02-2006 15:04:05"))
		return
	}

	fmt.Println("Welcome to south-american-qualifiers-cli!")
	fmt.Println("type -h or --help for more info")

}
