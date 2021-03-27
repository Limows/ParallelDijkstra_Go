// Algorithms
package main

import (
	"log"
)

func DijkstraSequential(Size int, GraphTable [][]int) (Results []string) {
	var Distances = make([]int, Size) // минимальное расстояние
	var Vertexes = make([]int, Size)  // посещенные вершины
	var Temp, MinIndex, Min int
	var BeginIndex = 0

	//Инициализация вершин и расстояний
	for i := 0; i < Size; i++ {
		Distances[i] = 10000
		Vertexes[i] = 1
	}

	Distances[BeginIndex] = 0
	// Шаг алгоритма
	for MinIndex < 10000 {
		MinIndex = 10000
		Min = 10000

		for i := 0; i < Size; i++ {

			if (Vertexes[i] == 1) && (Distances[i] < Min) { // Если вершину ещё не обошли и вес меньше min
				// Переприсваиваем значения
				Min = Distances[i]
				MinIndex = i
			}
		}
		// Добавляем найденный минимальный вес
		// к текущему весу вершины
		// и сравниваем с текущим минимальным весом вершины
		if MinIndex != 10000 {
			for i := 0; i < Size; i++ {
				if GraphTable[MinIndex][i] > 0 {
					Temp = Min + GraphTable[MinIndex][i]
					if Temp < Distances[i] {
						Distances[i] = Temp
					}
				}
			}
			Vertexes[MinIndex] = 0
		}
	}

	// Вывод кратчайших расстояний до вершин
	log.Println("Mininal distances to vertexes: ")

	for i := 0; i < Size; i++ {
		log.Println(Distances[i], " ")
	}

	// Восстановление пути
	UsedVertexes := make([]int, Size) // массив посещенных вершин
	end := Size - 1                   // индекс конечной вершины = 5 - 1
	UsedVertexes[0] = Size            // начальный элемент - конечная вершина
	k := 1                            // индекс предыдущей вершины
	weight := Distances[end]          // вес конечной вершины

	for end != BeginIndex { // пока не дошли до начальной вершины
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
	log.Println("Shortest Way:")

	for i := k - 1; i >= 0; i-- {
		log.Println(Vertexes[i], " ")
	}

	return nil
}
