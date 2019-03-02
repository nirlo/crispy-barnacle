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
	initials        string
	comments        string
}

var (
	pathToFile = ""
	reader     = bufio.NewReader(os.Stdin)
	records    []Row
)

func main() {
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

		choice, _, err := reader.ReadRune()
		check(err)

		switch choice {
		case 'a':
			loadFile(initial)
			initial = false
		case 'b':
			if initial {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			displayRecords()
		case 'c':
			if initial {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			createRecord()
		case 'd':
			if initial {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			workWithRecord()
		case 'e':
			if initial {
				fmt.Println("whoops, load a file first dum dum. Nicholas Lockhart")
				continue
			}
			deleteRecord()
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadFile(initial bool) {
	//TODO get the file name
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
				fmt.Print("\t" + typeOfT.Field(i).Name)
			}
			fmt.Println()
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

func createRecord() {
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

	records = append(records, Row{species, year, day, id, bud, flower, mature, initials, comment})

}

func workWithRecord() {
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
}

func deleteRecord() {
	fmt.Println("Alright, tell me the index of the record you want to get rid of: ")

	reader := bufio.NewReader(os.Stdin)
	in, _, err := reader.ReadRune()
	check(err)
	i := int(in)
	records = append(records[:i], records[i+1:]...)
	fmt.Println("check it out, it should be deleted! Nicholas Lockhart")
}
