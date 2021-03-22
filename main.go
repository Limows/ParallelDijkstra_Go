// Dijkstra
package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MainForm struct {
	*walk.MainWindow
}

type DialogForm struct {
	*walk.Dialog
}

var InitDirectory string = ""

func main() {
	var OpenAction, NewAction, SaveAction, CalculateAction, AboutBoxAction *walk.Action
	var FileMenu, ActionsMenu, HelpMenu *walk.Menu
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
						AssignTo: &CalculateAction,
						Text:     "&Calculate minimum distance",
					},
				},
			},

			Menu{
				Text:     "&Help",
				AssignTo: &HelpMenu,
				Items: []MenuItem{
					Action{
						AssignTo: &AboutBoxAction,
						Text:     "&About",
					},
				},
			},
		},

		Children: []Widget{
			TableView{},

			PushButton{
				Text: "&Calculate minimum distance",
			},

			Label{
				Text: "&Minimum distance: ",
			},
		},
	}.Run()
}

func OpenFile(Owner walk.Form, InitDirectory string) (FileName string, Error error) {
	FileDialog := new(walk.FileDialog)

	FileDialog.Title = "Open graph file"
	FileDialog.Filter = "CSV files(*.csv)|*.csv"

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

func ReadFile(FileName string) (Error error) {

	File, Error := os.Open(FileName)
	if Error != nil {
		File.Close()
		log.Println(Error)
		return Error
	}

	log.Println("File Opened")

	Reader := csv.NewReader(File)

	Reader.Comma = ';'
	Reader.Comment = '#'

	//Record, Error := Reader.Read()
	//if Error != nil {
	//	log.Println(Error)
	//	return Error
	//}

	//n, Error := strconv.Atoi(Record[0])
	//m, Error := strconv.Atoi(Record[1])

	//log.Println("Table size: ", n, m)

	GraphTable := make([][]int, 0)

	//for i := range GraphTable {
	//	GraphTable[i] = make([]int, m)
	//}

	for {
		Record, Error := Reader.Read()
		if Error != nil {
			log.Println(Error)
			break
		}

		m := len(Record)

		GraphColumn := make([]int, m)

		for j := 0; j < m; j++ {
			GraphColumn[j], Error = strconv.Atoi(Record[j])
		}

		GraphTable = append(GraphTable, GraphColumn)

		log.Println("Value: ", GraphColumn)
	}

	return nil
}

func SaveFile(Owner walk.Form, InitDirectory string) (FileName string, Error error) {
	FileDialog := new(walk.FileDialog)

	FileDialog.Title = "Save graph file"
	FileDialog.Filter = "CSV files(*.csv)|*.csv"

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
	}
}

func WriteFile(FileName string) {

}

func (owner *MainForm) SaveAction_Triggered() {

	if FileName, Error := SaveFile(owner, InitDirectory); Error == nil {
		WriteFile(FileName)
	}
}
