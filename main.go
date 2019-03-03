/*
	Simple CLI program for taking in a csv file and placing the lines
	into a slice to be edited and changed

	Authored by: Nicholas Lockhart
	Course: CST 8333 Program Language Reseach Project
	 Professor: Stan Pieda
*/package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Row A struct specifically for a row of data. This is what will be pulled from the csv
type Row struct {
	species         string
	year            string
	dayOfYear       string
	id              string
	numberOfBuds    string
	numberOfFlowers string
	numberMature    string
	initials        string
	comments        string
}

// Global variables. These will be used through out the program.
var (
	pathToFile = ""
	reader     = bufio.NewReader(os.Stdin)
	columns    = ""
	records    []Row
)

// Main function, prints the menu out to the console for the user
func main() {

	// How we determine if we are going to stop the program
	run := true
	initial := true

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(`Welcome to the File Reader by Nicholas Lockhart!
		Choose one of the following options:
		a. load/reload data from a file
		b. Display all records 
		c. Create a new record
		d. Display specific record
		e. Delete a record
		f. Quit the awesome File Reader
		Choice: `)

		//Grabs the choice as a char and switch from there
		choice, _, err := reader.ReadRune()
		check(err)

		switch choice {
		// Read the file in. Will overwrite the in memory structure if necessary
		case 'a':
			records = loadFile(initial)
			initial = false
		// checks to see if the file is in memory, displays them otherwise
		case 'b':
			if initial {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			displayRecords(records)
		// Add a record to the in memory structure
		case 'c':
			if initial {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			records = createRecord(records)
		// change a member of the structure
		case 'd':
			if initial {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			records = workWithRecord(records)
		// Remove a record from the structure
		case 'e':
			if initial {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			records = deleteRecord(records)
		case 'f':
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

/*
	A simple check for errors when processing a file or read.
	applying DRY methodology
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
	Loads the file from disk into memory. Takes the first line as the columns
	and removes them from the in memory structure.
*/
func loadFile(initial bool) []Row {

	var columns = ""
	if initial {
		fmt.Println("Let's read some dope files! Enter the path here:")
		pathToFile, _ = reader.ReadString('\n')
	} else {
		fmt.Println("Let's reload the data! Enter the path to the file:")
		pathToFile, _ = reader.ReadString('\n')
	}
	pathToFile = strings.TrimSpace(pathToFile)

	fmt.Println("Nicholas Lockhart")

	f, err := os.Open(pathToFile)
	check(err)
	defer f.Close()

	lineNo := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if lineNo == 1 {
			columns = line
			continue
		}
		if lineNo == 2 {
			continue
		}
		row := strings.Split(line, ",")
		records = append(records, Row{row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8]})
		if len(records) == 12 {
			break
		}
		lineNo++
	}

	// Prepare the columns for printing.
	columns = strings.Replace(columns, ",", " \t", 0)

	fmt.Println("Data from file retrieved! Nicholas Lockhart")

	return records
}

/*
	Display all of the records. Prints the columns out first for easy reading.
*/
func displayRecords(records []Row) {
	fmt.Println("Congrats, you get to read the data that we have! Nicholas Lockhart")

	for i, s := range records {
		if i == 0 {
			fmt.Println(columns)
		}
		fmt.Println(i+1, "\t",
			s.species, "\t",
			s.year, "\t",
			s.dayOfYear, "\t",
			s.id, "\t",
			s.numberOfBuds, "\t",
			s.numberOfFlowers, "\t",
			s.numberMature, "\t",
			s.initials, "\t",
			s.comments)
	}
	fmt.Println("Nicholas Lockhart")

}

/*
	Add a record to the in memory structure
*/
func createRecord(records []Row) []Row {
	//TODO add a new record to slice
	fmt.Println("Awesome, let's make a new record!")

	fmt.Println("What species of Flower is it?")
	reader := bufio.NewReader(os.Stdin)
	species, err := reader.ReadString('\n')
	check(err)
	species = strings.TrimSuffix(species, "\n")

	fmt.Println("What year is it?")
	year, err := reader.ReadString('\n')
	check(err)
	year = strings.TrimSuffix(year, "\n")

	fmt.Println("What is the julian day of the year?")
	day, err := reader.ReadString('\n')
	check(err)
	day = strings.TrimSuffix(day, "\n")

	fmt.Println("What is the id of the flower")
	id, err := reader.ReadString('\n')
	check(err)
	id = strings.TrimSuffix(id, "\n")

	fmt.Println("How many buds are there?")
	bud, err := reader.ReadString('\n')
	check(err)
	bud = strings.TrimSuffix(bud, "\n")

	fmt.Println("How many flowers?")
	flower, err := reader.ReadString('\n')
	check(err)
	flower = strings.TrimSuffix(flower, "\n")

	fmt.Println("How many are mature?")
	mature, err := reader.ReadString('\n')
	check(err)
	mature = strings.TrimSuffix(mature, "\n")

	fmt.Println("Place your initials")
	initials, err := reader.ReadString('\n')
	check(err)
	initials = strings.TrimSuffix(initials, "\n")

	fmt.Println("Any Comments?")
	comment, err := reader.ReadString('\n')
	check(err)
	comment = strings.TrimSuffix(comment, "\n")

	fmt.Println("Placing your info into the list! Nicholas Lockhart")

	records = addRecord(records, Row{species, year, day, id, bud, flower, mature, initials, comment})

	return records

}

// Separated out the add Records functionality
func addRecord(records []Row, row Row) []Row {
	records = append(records, row)
	return records
}

/*
	Change the in memory structure
*/
func workWithRecord(records []Row) []Row {
	//TODO work with record
	fmt.Println("Alright, tell me the index of the record you want to change: ")

	reader := bufio.NewReader(os.Stdin)
	in, _, err := reader.ReadRune()
	check(err)
	i := int(in)

	fmt.Println("tell me what you want to change: species, year, dayOfYear, id, numberOfBuds, numberOfFlowers, numberMature, initals, comments")

	field, err := reader.ReadString('\n')
	check(err)
	field = strings.TrimSuffix(field, "\n")

	fmt.Println("The new value:")

	value, err := reader.ReadString('\n')
	check(err)
	value = strings.TrimSuffix(value, "\n")

	switch field {
	case "species":
		records[i].species = value
	case "year":
		records[i].year = value
	case "dayOfYear":
		records[i].dayOfYear = value
	case "id":
		records[i].id = value
	case "numberOfBuds":
		records[i].numberOfBuds = value
	case "numberOfFlowers":
		records[i].numberOfFlowers = value
	case "numberMature":
		records[i].numberMature = value
	case "initials":
		records[i].initials = value
	case "comments":
		records[i].comments = value
	default:
		fmt.Println("whoops you made a mistake. Returning to main menu! Nicholas Lockhart")
	}

	return records
}

/*
	delete an in memory structure
*/
func deleteRecord(records []Row) []Row {
	fmt.Println("Alright, tell me the index of the record you want to get rid of: ")

	reader := bufio.NewReader(os.Stdin)
	in, _, err := reader.ReadRune()
	check(err)
	i := int(in)
	records = deleteFromSlice(records, i)
	fmt.Println("check it out, it should be deleted! Nicholas Lockhart")

	return records
}

// Separated out the deletion statement from delete records
func deleteFromSlice(records []Row, i int) []Row {
	// Golang is a bit funny when it comes to removing from a slice
	// This essentially creates a new slice from all records except for the
	// index we do not want anymore
	records = append(records[:i], records[i+1:]...)
	return records
}
