/*
	Testign the delete function and the add function

	Authored by: Nicholas Lockhart
	Course: CST 8333 Program Language Reseach Project
	 Professor: Stan Pieda
*/package main

import "testing"

// I generate an two empty Row structs in a slice and use the delete
// funciton, testing the length after.
func testDelete(t *testing.T) {
	var test = []Row{
		{"", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", ""},
	}

	// delete the 0th index
	test = deleteFromSlice(test, 0)

	if len(test) != 1 {
		t.Error("We have a failure!")
	}
}

// I generate a slice of two empty Row structs, and add an
// additional one, testing the length to ensure that it is finished
func testadd(t *testing.T) {
	var test = []Row{
		{"", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", ""},
	}

	// add a record
	test = addRecord(test, Row{"", "", "", "", "", "", "", "", ""})

	if len(test) != 3 {
		t.Error("We have a failure!")
	}
}
