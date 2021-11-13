package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
)

// Format 1
type F1 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

// Format 2
type F2 struct {
	Name       string
	Surname    string
	Areacode   string
	Tel        string
	LastAccess string
}

type Book1 []F1
type Book2 []F2

// CSVFILE resides in the home directory of the current user
var CSVFILE = ""
var d1 = Book1{}
var d2 = Book2{}

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

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	var firstLine = true
	var format1 = true
	for _, line := range lines {
		if firstLine {
			if len(line) == 4 {
				format1 = true
			} else if len(line) == 5 {
				format1 = false
			} else {
				return errors.New("Unknown File Format!")
			}
			firstLine = false
		}

		if format1 {
			if len(line) == 4 {
				temp := F1{
					Name:       line[0],
					Surname:    line[1],
					Tel:        line[2],
					LastAccess: line[3],
				}
				d1 = append(d1, temp)
			}
		} else {
			if len(line) == 5 {
				temp := F2{
					Name:       line[0],
					Surname:    line[1],
					Areacode:   line[2],
					Tel:        line[3],
					LastAccess: line[4],
				}
				d2 = append(d2, temp)
			}
		}
	}
	return nil
}

// Implement sort.Interface for Book1
func (b1 Book1) Len() int {
	return len(b1)
}

// First based on surname. If they have the same
// surname take into account the name.
func (b1 Book1) Less(i, j int) bool {
	if b1[i].Surname == b1[j].Surname {
		return b1[i].Name < b1[j].Name
	}
	return b1[i].Surname < b1[j].Surname
}

func (b1 Book1) Swap(i, j int) {
	b1[i], b1[j] = b1[j], b1[i]
}

// Implement sort.Interface for Book2
func (b2 Book2) Len() int {
	return len(b2)
}

// First based on areacode. If they have the same
// areacode take into account the surname.
func (b2 Book2) Less(i, j int) bool {
	if b2[i].Areacode == b2[j].Areacode {
		return b2[i].Surname < b2[j].Surname
	}
	return b2[i].Areacode < b2[j].Areacode
}

func (b2 Book2) Swap(i, j int) {
	b2[i], b2[j] = b2[j], b2[i]
}

func list(d interface{}) {
	switch T := d.(type) {
	case Book1:
		data := d.(Book1)
		for _, v := range data {
			fmt.Println(v)
		}
	case Book2:
		data := d.(Book2)
		for _, v := range data {
			fmt.Println(v)
		}
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func sortData(d interface{}) {
	switch T := d.(type) {
	case Book1:
		data := d.(Book1)
		sort.Sort(data)
		list(data)
	case Book2:
		data := d.(Book2)
		sort.Sort(data)
		list(data)
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func main() {
	if len(os.Args) != 1 {
		CSVFILE = os.Args[1]
	} else {
		fmt.Println("No data file!")
		return
	}

	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println(CSVFILE, "does not exist!")
		return
	}

	fileInfo, err := os.Stat(CSVFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "not a regular file!")
		return
	}

	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(d1) != 0 {
		sortData(d1)
	} else {
		sortData(d2)
	}
}
