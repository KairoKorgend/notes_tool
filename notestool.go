package sprint

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Println("Usage: ./notestool [TAG]")
		os.Exit(1)
	}

	collectionName := os.Args[1]
	notes := loadNotes(collectionName)

	fmt.Printf("Welcome to the notes tool!\n")

	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		var operation string
		fmt.Scanln(&operation)

		switch operation {
		case "1": //Shows existing notes if there are any
			showNotes(collectionName, notes)
		case "2": //Creates the inputted note
			notes = addNote(notes)
		case "3": //Deletes the selected note
			notes = deleteNote(notes)
		case "4": //Saves changes and exits the tool
			saveNotes(collectionName, notes)
			os.Exit(0)
		}
	}
}

func showNotes(collectionName string, notes []string) {
	if len(notes) <= 0 {
		fmt.Printf("\nCollection '%s' is empty. Please select operation 2 to add a note.\n", collectionName)
		return
	}
	fmt.Println("\nNotes:")
	for i, note := range notes { //Adds note numbers in front of notes
		if note != "" {
			fmt.Printf("%03d - %s\n", i+1, note)
		} else {
			fmt.Println()
		}
	}
}

func addNote(notes []string) []string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEnter the note text:\n")
		scanner.Scan()
		newNote := scanner.Text()
		newNote = strings.TrimSpace(newNote) //trims white spaces and makes sure the note cannot be created empty
		if newNote != "" {
			return append(notes, newNote)
		} else { //if the note is empty return error
			fmt.Println("Note text cannot be empty. Please try again.")
		}
	}
}

func deleteNote(notes []string) []string {
	fmt.Print("\nEnter the number of the note to remove or 0 to cancel:\n")
	scanner := bufio.NewScanner(os.Stdin) //Takes the input for the note
	scanner.Scan()
	noteNumberStr := scanner.Text()

	noteNumber, err := strconv.Atoi(noteNumberStr) //Changes the note from string into int
	if err != nil {
		fmt.Println("\nNote does not exist. Please try again.")
		return notes
	}

	if noteNumber > 0 && noteNumber <= len(notes) { //Checks if the noteNumber is higher than 0 and lower than the length of the notes
		notes = append(notes[:noteNumber-1], notes[noteNumber:]...)
	} else if noteNumber != 0 { //If it number isn't within the parameters, returns an error message
		fmt.Println("\nNote does not exist. Please try again.")
	}

	return notes
}

func loadNotes(collectionName string) []string {
	var notes []string
	file, err := os.Open(collectionName + ".txt")
	if err != nil {
		return notes
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	return notes
}

func saveNotes(collectionName string, notes []string) {
	file, err := os.Create(collectionName + ".txt")
	if err != nil {
		fmt.Println("Error saving notes:", err)
		return
	}
	defer file.Close()

	for _, note := range notes {
		fmt.Fprintln(file, note)
	}
}
