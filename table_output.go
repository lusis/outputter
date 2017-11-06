package outputter

import (
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

func init() {
	RegisterOutput("table", newTableFactoryOutput)
}

// TableOutput is an Outputter that draw data as a table
type TableOutput struct {
	table   *tablewriter.Table
	headers []string
}

func newTableFactoryOutput() Outputter {
	x := NewTableOutput()
	return x
}

// NewTableOutput creates a New TableOutput with os.Stdout as the destination
func NewTableOutput() *TableOutput {
	return NewTableOutputWithWriter(os.Stdout)
}

// NewTableOutputWithWriter creates a new instance of TableOutput with the provided io.Writer
func NewTableOutputWithWriter(i io.Writer) *TableOutput {
	t := &TableOutput{}
	tw := tablewriter.NewWriter(i)
	tw.SetAutoWrapText(false)
	tw.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	tw.SetAlignment(tablewriter.ALIGN_LEFT)
	t.table = tw
	return t
}

// SetHeaders sets the table's headers
func (t *TableOutput) SetHeaders(headers []string) {
	t.headers = headers
	t.table.SetHeader(headers)
}

// AddRow adds a row to the table
func (t *TableOutput) AddRow(row []string) error {
	if len(t.headers) == 0 {
		return ErrorOutputAddRowNoHeaders
	}
	if len(t.headers) < len(row) {
		return ErrorOutputAddRowTooFewHeaders
	}
	t.table.Append(row)
	return nil
}

// Draw displays the table to stdout
func (t *TableOutput) Draw() {
	t.table.Render()
}

// SetPretty returns a prettified version
func (t *TableOutput) SetPretty() {
	//noop for table
}

// ColorSupport specifies if the output supports colorized text or not
func (t *TableOutput) ColorSupport() bool {
	return true
}
