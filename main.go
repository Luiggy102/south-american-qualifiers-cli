package main

import (
	"fmt"
	"time"

	"github.com/Luigy102/south-american-qualifiers-cli/cmd"
	"github.com/Luigy102/south-american-qualifiers-cli/internal/requests"
	"github.com/Luigy102/south-american-qualifiers-cli/types"
)

func main() {
	var table types.Table = requests.Table()
	cmd.PrintTable(table)

	fmt.Println("\nUpdated at:")
	fmt.Println(time.Now().Format("Monday 01-02-2006 15:04:05"))
}
