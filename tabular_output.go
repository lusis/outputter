package console

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

func init() {
	RegisterOutput("tabular", newTabularFactoryOutput)
}

func newTabularFactoryOutput() Outputter {
	return NewTabularOutput()
}

// TabularOutput is an Outputter that draws data in tabular format
type TabularOutput struct {
	table *tabwriter.Writer
}

// NewTabularOutput creates a New TablularOutput with os.Stdout as the destination
func NewTabularOutput() *TabularOutput {
	t := NewTabularOutputWithWriter(os.Stdout)
	return t
}

// NewTabularOutputWithWriter creates a new instance of TabularOutput with the provided io.Writer
func NewTabularOutputWithWriter(i io.Writer) *TabularOutput {
	t := &TabularOutput{}
	w := tabwriter.NewWriter(i, 0, 2, 2, ' ', 0)
	t.table = w
	return t
}

// SetHeaders sets the table's headers
func (t *TabularOutput) SetHeaders(headers []string) {
	fmt.Fprintf(t.table, strings.Join(headers, "\t")+"\n")
}

// AddRow adds a row to the table
func (t *TabularOutput) AddRow(row []string) error {
	fmt.Fprintf(t.table, strings.Join(row, "\t")+"\n")
	return nil
}

// SetPretty sets pretty output
func (t *TabularOutput) SetPretty() {
	//noop
}

// Draw displays the table to stdout
func (t *TabularOutput) Draw() {
	_ = t.table.Flush()
}

// ColorSupport specifies if the output supports colorized text or not
func (t *TabularOutput) ColorSupport() bool {
	// have to turn off color on tabular output for now =(
	// http://stackoverflow.com/questions/35398497/how-do-i-get-colors-to-work-with-golang-tabwriter
	return false
}
