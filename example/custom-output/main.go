package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/lusis/outputter"
)

// NewCustomFactoryOutput returns my custom outputter
func NewCustomFactoryOutput() outputter.Outputter {
	return NewCustomOutput()
}

// CustomOutput is an Outputter that draws data in custom format
type CustomOutput struct {
	keys   []string
	writer io.Writer
	data   []map[string]string
	sync.RWMutex
}

// NewCustomOutput creates a New TablularOutput with os.Stdout as the destination
func NewCustomOutput() *CustomOutput {
	t := NewCustomOutputWithWriter(os.Stdout)
	return t
}

// NewCustomOutputWithWriter creates a new instance of TabularOutput with the provided io.Writer
func NewCustomOutputWithWriter(i io.Writer) *CustomOutput {
	t := &CustomOutput{writer: i}
	return t
}

// SetHeaders sets the table's headers
func (t *CustomOutput) SetHeaders(headers []string) {
	t.Lock()
	defer t.Unlock()
	t.keys = headers
}

// AddRow adds a row to the table
func (t *CustomOutput) AddRow(row []string) error {
	if len(t.keys) == 0 {
		return outputter.ErrorOutputAddRowNoHeaders
	}
	if len(t.keys) < len(row) {
		return outputter.ErrorOutputAddRowTooFewHeaders
	}
	t.Lock()
	defer t.Unlock()
	m := make(map[string]string)
	if len(row) < len(t.keys) {
		difference := len(t.keys) - len(row)
		// we have to account for this and fill in empty values
		missingVals := make([]string, difference)
		row = append(row, missingVals...)
	}
	for keyIdx, keyName := range t.keys {
		m[keyName] = row[keyIdx]
	}
	t.data = append(t.data, m)
	return nil
}

// SetPretty sets pretty output
func (t *CustomOutput) SetPretty() {
	//noop
}

// Draw displays the table to stdout
func (t *CustomOutput) Draw() {
	var res []string
	res = append(res, "here's my data!")
	for _, headerAndValues := range t.data {
		for k, v := range headerAndValues {
			res = append(res, fmt.Sprintf("%s:%s", k, v))
		}
	}
	fmt.Fprintf(t.writer, strings.Join(res, "\n")+"\n")
}

// ColorSupport specifies if the output supports colorized text or not
func (t *CustomOutput) ColorSupport() bool {
	return true
}
func main() {
	outputter.RegisterOutput("myoutput", NewCustomFactoryOutput)
	outputters := outputter.GetOutputters()
	log.Printf("found outputters: %s", strings.Join(outputters, ","))
	outputFormatter, err := outputter.NewOutputter("myoutput")
	if err != nil {
		log.Fatalf("unable to get an outputter: %s", err.Error())
	}
	outputFormatter.SetHeaders([]string{"header1", "header2", "header3"})
	rowErr := outputFormatter.AddRow([]string{"value1", "value2", "value3"})
	if rowErr != nil {
		log.Fatalf("error adding row: %s", rowErr.Error())
	}
	outputFormatter.Draw()
}
