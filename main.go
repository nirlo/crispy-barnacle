package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

//Species,Year,Julian Day of Year,Plant Identification Number,Number of Buds,Number of Flowers,Number of Flowers that have Reached Maturity,Observer Initials,Observer Comments

// Row A struct specifically for a row of data. This is what will be pulled from the csv
type Row struct {
	species         string
	year            string
	dayOfYear       string
	id              string
	numberOfBuds    string
	numberOfFlowers string
	numberMature    string
	initals         string
	comments        string
}

var (
	pathToFile = ""
	reader     = bufio.NewReader(os.Stdin)
	records    []Row
)

func main() {
	run := true
	pathToFile := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(`Welcome to the File Reader by Nicholas Lockhart!
		Choose one of the following options:
		1. load/reload data from a file
		2. Display all records 
		3. Create a new record
		4. Display specific record
		5. Delete a record
		6. Quit the awesome File Reader
		Choice: `)

		choice, _ := reader.ReadString('\n')

		switch choice {
		case "1":
			loadFile()
		case "2":
			if pathToFile == "" {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			displayRecords()
		case "3":
			if pathToFile == "" {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			createRecord()
		case "4":
			if pathToFile == "" {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			workWithRecord()
		case "5":
			if pathToFile == "" {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			deleteRecord()
		case "6":
			fmt.Println("Sorry to see you go. Goodbye! Nicholas Lockhart")
			run = false
		default:
			fmt.Println("whoops, you did something very wrong here. GO BACK!")
		}
		if !run {
			break
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadFile() {
	//TODO get the file name
	if pathToFile == "" {
		fmt.Println("Let's read some dope files! Enter the path here:")
		pathToFile, _ = reader.ReadString('\n')
	} else {
		fmt.Println("Let's reload the data! Enter the path to the file:")
		pathToFile, _ = reader.ReadString('\n')
	}

	fmt.Println("Nicholas Lockhart")

	f, err := os.Open(pathToFile)
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, ",")
		records = append(records, Row{row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8]})
		if len(records) == 12 {
			break
		}
	}

	// Removing the column names, we won't need them were we are going
	records = append(records[:1], records[2:]...)
	records = append(records[:0], records[1:]...)

	fmt.Println("Data from file retrieved! Nicholas Lockhart")
}

func displayRecords() {
	fmt.Println("Congrats, you get to read the data that we have! Nicholas Lockhart")

	ref := reflect.ValueOf(&records[0]).Elem()
	typeOfT := ref.Type()

	for i, s := range records {
		if i == 0 {
			for i := 0; i < ref.NumField(); i++ {
				fmt.Println("%s\t",
					typeOfT.Field(i).Name)
			}
		}
		fmt.Println("%d: %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", i+1,
			s.species, s.year, s.dayOfYear, s.id, s.numberOfBuds, s.numberOfFlowers, s.numberMature, s.initals, s.comments)
	}
	fmt.Println("Nicholas Lockhart")
}

func createRecord() {
	//TODO add a new record to slice
}

func workWithRecord() {
	//TODO work with record
}

func deleteRecord() {

}
