package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func processFile1() ([]string, error) {
	//will pass filepath
	file1 := "files/credit_test_data.csv"

	//will open file
	openfile1, err := os.Open(file1)
	checkError("Error in opening the file\n", err)

	//read the data of file
	filedata1, err := csv.NewReader(openfile1).Read()
	// filedata1, err := csv.NewReader(openfile1).ReadAll()
	checkError("Error in reading the file\n", err)

	// for i, value := range filedata1 {
	// 	fmt.Printf("i= %d where type is %T, value %s where type is %T\n", i, i, value, value)
	// }

	for index, value := range filedata1 {
		fmt.Println(index, value)
	}

	return filedata1, nil

}

func processFile2() ([]string, error) {
	//will pass filepath
	file2 := "files/credit_train_data.csv"

	openfile2, err := os.Open(file2)
	checkError("Error in opening the file\n", err)

	filedata2, err := csv.NewReader(openfile2).Read()
	// filedata2, err := csv.NewReader(openfile2).ReadAll()
	checkError("Error in reading the file\n", err)

	// for i, value := range filedata2 {
	// 	fmt.Printf("i= %d where type is %T, value %s where type is %T\n", i, i, value, value)
	// }

	for index, value := range filedata2 {
		fmt.Println(index, value)
	}

	return filedata2, nil
}

func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func main() {
	filedata1, err := processFile1()

	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	filedata2, err := processFile2()

	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	/*
		Checks if Headers match
		E.g: Filedata1 'Customer ID' and Filedata2 'Customer ID' match
	*/

	if len(filedata1) == len(filedata2) {
		for i := 0; i < len(filedata1); i++ {
			if filedata1[i] == filedata2[i] {
				fmt.Println("filedata1: ", filedata1[i], "\nfiledata2: ", filedata1[i])
				fmt.Println("MATCH!")
			}
		}
	}
}
