// ParamsHelper
package main

import (
	"github.com/lxn/walk"
)

type MainForm struct {
	*walk.MainWindow
}

type GraphModel struct {
	walk.TableModelBase
	items [][]int
}

func NewGraphModel() *GraphModel {
	model := new(GraphModel)

	model.ResetRows()

	return model
}

func (model *GraphModel) RowCount() int {
	return len(model.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (model *GraphModel) Value(row, col int) interface{} {
	item := model.items[row][col]

	return item
}

// Called by the TableView to retrieve if a given row is checked.
func (model *GraphModel) Checked(row int) bool {
	return false
}

// Called by the TableView when the user toggled the check box of a given row.
func (model *GraphModel) SetChecked(row int, checked bool) error {
	return nil
}

// Called by the TableView to sort the model.
func (model *GraphModel) Sort(col int, order walk.SortOrder) error {
	return nil
}

func (model *GraphModel) ResetRows() {
	// Create some random data.
	model.items = make([][]int, 2)

	for i := range model.items {
		model.items[i] = make([]int, 2)
	}

	for i := range model.items {
		for j := range model.items[i] {
			model.items[i][j] = 0
		}
	}

	// Notify TableView and other interested parties about the reset.
	model.PublishRowsReset()
}

var InitDirectory string = ""
