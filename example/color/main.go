package main

import (
	"flag"
	"log"

	color "github.com/fatih/color"
	outputter "github.com/lusis/outputter"
)

var outputFormat = flag.String("format", "tabular", "format for output")

func main() {
	flag.Parse()
	var outputFormatter outputter.Outputter
	switch *outputFormat {
	case "json":
		outputFormatter = outputter.NewJSONOutput()
	case "tabular":
		outputFormatter = outputter.NewTabularOutput()
	case "table":
		outputFormatter = outputter.NewTableOutput()
	default:
		outputFormatter = outputter.NewTabularOutput()
	}

	// set color output based on if the outputter supports it
	// not that `ColorSupport()` returns false if not supported
	// so inverse will need to be passed to `color.NoColor`
	color.NoColor = !outputFormatter.ColorSupport()
	outputFormatter.SetHeaders([]string{
		"header1",
		"header2",
		"header3",
	})
	rowErr := outputFormatter.AddRow([]string{
		color.YellowString("value1"),
		color.YellowString("value2"),
		color.YellowString("value3"),
	})
	if rowErr != nil {
		log.Fatalf("error adding row: %s", rowErr.Error())
	}

	outputFormatter.Draw()
}
