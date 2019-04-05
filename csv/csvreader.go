package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	//will pass filepath
	filepath := "files/credit_test.csv"

	//will open file
	openfile, err := os.Open(filepath)
	checkError("Error in opening the file\n", err)

	//read the data of file
	filedata, err := csv.NewReader(openfile).ReadAll()
	checkError("Error in reading the file\n", err)

	for i, value := range filedata {
		fmt.Printf("i= %d where type is %T, value %s where type is %T\n", i, i, value, value)
	}
}

func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
