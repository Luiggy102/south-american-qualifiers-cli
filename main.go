package main

import (
	"github.com/Luigy102/south-american-qualifiers-cli/cmd"
	"github.com/Luigy102/south-american-qualifiers-cli/internal/requests"
	"github.com/Luigy102/south-american-qualifiers-cli/types"
)

func main() {
	var table types.Table = requests.Table()
	cmd.PrintTable(table)
}
