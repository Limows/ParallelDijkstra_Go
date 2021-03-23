// Algorithms
package main

import (
	"log"
)

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
	log.Println("MinDist: ")

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
	log.Println("MinDist: ")
}
