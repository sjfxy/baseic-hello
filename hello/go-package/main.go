package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	data := [][]string{
		[]string{"Sin", "15", "10/21", "(10.22,30.25)"},
		[]string{"Sin", "15", "10/21", "(10.22,30.25)"},
		[]string{"Sin", "15", "10/21", "(10.22,30.25)"},
		[]string{"Sin", "15", "10/21", "(10.22,30.25)"},
		[]string{"Sin", "15", "10/21", "(10.22,30.25)"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NPC", "Speed", "Power", "Location"})
	table.AppendBulk(data)
	table.Render()

}
