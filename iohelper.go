// IOHelper
package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

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

func WriteFile(FileName string) {

}
