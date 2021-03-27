// IOHelper
package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ReadFile(FileName string) (GraphTable [][]int, Error error) {

	File, err := os.Open(FileName)
	if err != nil {
		File.Close()
		log.Println(Error)
		return nil, Error
	}

	log.Println("File Opened")

	Reader := csv.NewReader(File)

	Reader.Comma = ';'
	Reader.Comment = '#'

	GraphTable = make([][]int, 0)

	for {
		Record, err := Reader.Read()
		if err != nil {
			log.Println(Error)
			break
		}

		m := len(Record)

		GraphColumn := make([]int, m)

		for j := 0; j < m; j++ {
			GraphColumn[j], err = strconv.Atoi(Record[j])
			if err != nil {
				log.Println(Error)
			}
		}

		GraphTable = append(GraphTable, GraphColumn)

		log.Println("Value: ", GraphColumn)
	}

	return GraphTable, nil
}

func WriteFile(FileName string) (Error error) {
	return nil
}
