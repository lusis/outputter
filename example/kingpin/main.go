package main

import (
	"log"

	outputter "github.com/lusis/outputter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var outputFormat string
	kingpin.Flag("format", "format for output").
		Default("tabular").
		EnumVar(&outputFormat, "table", "json", "tabular")
	kingpin.Parse()
	var outputFormatter outputter.Outputter
	switch outputFormat {
	case "json":
		outputFormatter = outputter.NewJSONOutput()
	case "tabular":
		outputFormatter = outputter.NewTabularOutput()
	case "table":
		outputFormatter = outputter.NewTableOutput()
	default:
		outputFormatter = outputter.NewTabularOutput()
	}

	outputFormatter.SetHeaders([]string{"header1", "header2", "header3"})
	rowErr := outputFormatter.AddRow([]string{"value1", "value2", "value3"})
	if rowErr != nil {
		log.Fatalf("error adding row: %s", rowErr.Error())
	}
	outputFormatter.Draw()
}
