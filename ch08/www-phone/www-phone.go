package main

import (
	"encoding/csv"
	"os"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

// JSONFILE resides in the current directory
var CSVFILE = "./data.csv"

type PhoneBook []Entry

var data = PhoneBook{}
var index map[string]int

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// CSV file read all at once
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		// Storing to global variable
		data = append(data, temp)
	}

	return nil
}
