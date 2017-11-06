package outputter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTableOutput(t *testing.T) {
	table := NewTableOutput()
	assert.IsType(t, new(TableOutput), table)
}

func TestTableOutputNew(t *testing.T) {
	var buf bytes.Buffer
	table := NewTableOutputWithWriter(&buf)
	table.SetHeaders([]string{"header1", "header2"})
	err := table.AddRow([]string{"value1", "value2"})
	assert.NoError(t, err)
	table.Draw()
	expectedOutput := `+---------+---------+
| HEADER1 | HEADER2 |
+---------+---------+
| value1  | value2  |
+---------+---------+
`
	assert.Equal(t, expectedOutput, buf.String())
}

func TestTableTooManyValues(t *testing.T) {
	table := NewTableOutput()
	table.SetHeaders([]string{"header1", "header2"})
	err := table.AddRow([]string{"value1", "value2", "value3"})
	assert.Equal(t, ErrorOutputAddRowTooFewHeaders, err)
}

func TestTableNoHeaders(t *testing.T) {
	table := NewTableOutput()
	err := table.AddRow([]string{"value1", "value2", "value3"})
	assert.Equal(t, ErrorOutputAddRowNoHeaders, err)
}
func TestTableOutputFewerValues(t *testing.T) {
	var buf bytes.Buffer
	table := NewTableOutputWithWriter(&buf)
	table.SetHeaders([]string{"header1", "header2"})
	err := table.AddRow([]string{"value1"})
	assert.NoError(t, err)
	table.Draw()
	expectedOutput := `+---------+---------+
| HEADER1 | HEADER2 |
+---------+---------+
| value1  |
+---------+---------+
`
	assert.Equal(t, expectedOutput, buf.String())
}
