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
	var CalculateButton *walk.PushButton
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
						AssignTo: &AboutBoxAction,
						Text:     "&About",
					},
				},
			},
		},

		Children: []Widget{
			TableView{},

			PushButton{
				AssignTo:  &CalculateButton,
				Text:      "&Calculate minimum distance",
				OnClicked: NewMainForm.CalculateButton_Clicked,
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

	GraphTable := make([][]int, 0)

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
			if Error != nil {
				log.Println(Error)
			}
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

func DijkstraSequential(Size int, GraphTable [][]int) {
	Distances := make([]int, Size) // минимальное расстояние
	Vertexes := make([]int, Size)  // посещенные вершины
	var temp, minindex, min int
	begin_index := 0

	//Инициализация вершин и расстояний
	for i := 0; i < Size; i++ {
		Distances[i] = 10000
		Vertexes[i] = 1
	}

	Distances[begin_index] = 0
	// Шаг алгоритма
	for minindex < 10000 {
		minindex = 10000
		min = 10000

		for i := 0; i < Size; i++ {

			if (Vertexes[i] == 1) && (Distances[i] < min) { // Если вершину ещё не обошли и вес меньше min
				// Переприсваиваем значения
				min = Distances[i]
				minindex = i
			}
		}
		// Добавляем найденный минимальный вес
		// к текущему весу вершины
		// и сравниваем с текущим минимальным весом вершины
		if minindex != 10000 {
			for i := 0; i < Size; i++ {
				if GraphTable[minindex][i] > 0 {
					temp = min + GraphTable[minindex][i]
					if temp < Distances[i] {
						Distances[i] = temp
					}
				}
			}
			Vertexes[minindex] = 0
		}
	}

	// Вывод кратчайших расстояний до вершин

	// Восстановление пути
	UsedVertexes := make([]int, Size) // массив посещенных вершин
	end := Size - 1                   // индекс конечной вершины = 5 - 1
	UsedVertexes[0] = Size            // начальный элемент - конечная вершина
	k := 1                            // индекс предыдущей вершины
	weight := Distances[end]          // вес конечной вершины

	for end != begin_index { // пока не дошли до начальной вершины
		for i := 0; i < Size; i++ { // просматриваем все вершины
			if GraphTable[i][end] != 0 { // если связь есть
				temp := weight - GraphTable[i][end] // определяем вес пути из предыдущей вершины
				if temp == Distances[i] {           // если вес совпал с рассчитанным
					// значит из этой вершины и был переход
					weight = temp           // сохраняем новый вес
					end = i                 // сохраняем предыдущую вершину
					UsedVertexes[k] = i + 1 // и записываем ее в массив
					k++
				}
			}
		}
	}

	// Вывод пути (начальная вершина оказалась в конце массива из k элементов)
}

func (owner *MainForm) CalculateAction_Triggered() {

}

func (owner *MainForm) CalculateButton_Clicked() {

}
