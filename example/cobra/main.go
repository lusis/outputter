package main

import (
	"log"

	outputter "github.com/lusis/outputter"
	"github.com/spf13/cobra"
)

var outputFormat string
var outputFormatter outputter.Outputter

func doOutput(cmd *cobra.Command, args []string) {
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
func main() {
	var rootCmd = &cobra.Command{
		Use:   "formatter",
		Short: "diplays output in different formats",
	}
	var outputCmd = &cobra.Command{
		Use: "run",
		Run: doOutput,
	}

	rootCmd.PersistentFlags().StringVar(&outputFormat, "format", "tabular", "Specify the output format: json, table, tabular")
	rootCmd.AddCommand(outputCmd)

	_ = rootCmd.Execute()
}
