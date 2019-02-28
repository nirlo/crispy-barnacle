package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	pathToFile = ""
	reader     = bufio.NewReader(os.Stdin)
	records    []string
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

}

func displayRecords() {
	//TODO display all records
}

func createRecord() {
	//TODO add a new record to slice
}

func workWithRecord() {
	//TODO work with record
}

func deleteRecord() {

}
