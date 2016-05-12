package contract

type Table struct {
	Type      string
	Rows      []*RowOfValues
	Signature *RowOfValues
}

func (table *Table) AddRow(row *RowOfValues) {
	table.Rows = append(table.Rows, row)
}

func NewTable(inaugralRow *RowOfValues) *Table {
	table := Table{
		Type:      "Table",
		Rows:      []*RowOfValues{inaugralRow},
		Signature: inaugralRow,
	}
	return &table
}
