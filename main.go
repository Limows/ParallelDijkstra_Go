// Dijkstra
package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var OpenAction, NewAction, SaveAction, CalculateAction, AboutBoxAction *walk.Action
	var FileMenu, ActionsMenu, HelpMenu *walk.Menu
	var CalculateButton *walk.PushButton
	var GraphTableView *walk.TableView
	var MinDistLabel *walk.Label
	GraphTableModel := NewGraphModel()
	NewMainForm := new(MainForm)

	MainWindow{
		Title:    "GraphMinDist",
		Size:     Size{400, 600},
		Layout:   VBox{},
		AssignTo: &NewMainForm.MainWindow,
		MenuItems: []MenuItem{
			Menu{
				Text:     "&File",
				AssignTo: &FileMenu,
				Items: []MenuItem{
					Action{
						AssignTo: &NewAction,
						Text:     "&New",
					},
					Action{
						AssignTo:    &OpenAction,
						Text:        "&Open",
						OnTriggered: NewMainForm.OpenAction_Triggered,
					},
					Action{
						AssignTo:    &SaveAction,
						Text:        "&Save",
						OnTriggered: NewMainForm.SaveAction_Triggered,
					},
					Separator{},
					Action{
						Text:        "&Exit",
						OnTriggered: func() { NewMainForm.Close() },
					},
				},
			},

			Menu{
				Text:     "&Actions",
				AssignTo: &ActionsMenu,
				Items: []MenuItem{
					Action{
						AssignTo:    &CalculateAction,
						Text:        "&Calculate minimum distance",
						OnTriggered: NewMainForm.CalculateAction_Triggered,
					},
				},
			},

			Menu{
				Text:     "&Help",
				AssignTo: &HelpMenu,
				Items: []MenuItem{
					Action{
						AssignTo:    &AboutBoxAction,
						Text:        "&About",
						OnTriggered: NewMainForm.AboutBoxAction_Triggered,
					},
				},
			},
		},

		Children: []Widget{
			Label{
				Name: "&GraphTableLable",
				Text: "&Graph Adjacency Table",
			},

			TableView{
				AssignTo:         &GraphTableView,
				ColumnsOrderable: true,
				Columns:          []TableViewColumn{},
				Model:            GraphTableModel,
			},

			Label{
				Text: "&Number of rows",
			},

			NumberEdit{},

			Label{
				Text: "&Number of columns",
			},

			NumberEdit{},

			PushButton{
				AssignTo:  &CalculateButton,
				Text:      "&Calculate minimum distance",
				OnClicked: NewMainForm.CalculateButton_Clicked,
			},

			Label{
				AssignTo: &MinDistLabel,
				Name:     "&MinDistLabel",
				Text:     "&Minimum distance: ",
			},
		},
	}.Run()
}

func OpenFile(Owner walk.Form, InitDirectory string) (FileName string, Error error) {
	FileDialog := new(walk.FileDialog)

	FileDialog.Title = "Open graph file"
	FileDialog.Filter = "CSV files(*.csv)|*.csv"
	FileDialog.InitialDirPath = InitDirectory

	if OK, Error := FileDialog.ShowOpen(Owner); Error != nil {
		log.Println(Error)
		return "", Error
	} else if !OK {
		return "", nil
	}

	FileName = FileDialog.FilePath

	log.Println("File Chosed")

	return FileName, nil
}

func SaveFile(Owner walk.Form, InitDirectory string) (FileName string, Error error) {
	FileDialog := new(walk.FileDialog)

	FileDialog.Title = "Save graph file"
	FileDialog.Filter = "CSV files(*.csv)|*.csv"
	FileDialog.InitialDirPath = InitDirectory

	if OK, Error := FileDialog.ShowSave(Owner); Error != nil {
		log.Println(Error)
		return "", Error
	} else if !OK {
		return "", nil
	}

	FileName = FileDialog.FilePath

	return FileName, nil
}

func (owner *MainForm) OpenAction_Triggered() {

	if FileName, Error := OpenFile(owner, InitDirectory); Error == nil {
		Error = ReadFile(FileName)

		if Error != nil {
			walk.MsgBox(owner, "Error", "Can't read file!", walk.MsgBoxIconError)
			log.Println("Can't read file!")
		}
	}
}

func (owner *MainForm) SaveAction_Triggered() {

	if FileName, Error := SaveFile(owner, InitDirectory); Error == nil {
		WriteFile(FileName)
	}
}

func (owner *MainForm) CalculateAction_Triggered() {

}

func (owner *MainForm) CalculateButton_Clicked() {

}

func (owner *MainForm) AboutBoxAction_Triggered() {
	walk.MsgBox(owner, "About", "This is parallel modification of Dijkstra algorithm\n\nAuthor: Limows\n\nVersion: 0.3", walk.MsgBoxIconInformation)
}
