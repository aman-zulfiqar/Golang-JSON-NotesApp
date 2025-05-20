package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

var notes []Note

const dataFile = "notes.json"

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})
	loadNotes()
}

func loadNotes() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("No existing notes file, starting fresh")
			notes = []Note{}
		}
		log.WithError(err).Error("Failed to read notes file")
		return
	}

	err = json.Unmarshal(file, &notes)
	if err != nil {
		log.WithError(err).Error("Failed to parse notes data")
		notes = []Note{}
	}
	log.Info("Successfully loaded notes")
}

func addNote() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter note title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("Enter note content: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	notes = append(notes, Note{
		Title:     title,
		Content:   content,
		Timestamp: time.Now(),
	})
	saveNotes()
	log.WithFields(log.Fields{
		"title": title,
	}).Info("Note added")
}

func saveNotes() {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		log.WithError(err).Error("Failed to marshal notes")
		return
	}
	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		log.WithError(err).Error("Failed to save notes")
		return
	}
	log.Info("Successfully saved notes")
}

func listNotes() {
	if notes == nil {
		log.Error("Notes data is not loaded. Please load notes before listing.")
		return
	}
	if len(notes) == 0 {
		fmt.Println("No notes available")
		return
	}
	for i, note := range notes {
		fmt.Printf("\nNote #%d\n", i+1)
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)
		fmt.Printf("Date: %s\n", note.Timestamp.Format(time.RFC822))
	}
}

func deleteNote() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter title of note to delete: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	for i, note := range notes {
		if strings.EqualFold(note.Title, title) {
			notes = append(notes[:i], notes[i+1:]...)
			saveNotes()
			log.WithFields(log.Fields{
				"title": title,
			}).Info("Note deleted")
			fmt.Println("Note deleted successfully")
			return
		}
	}
	fmt.Println("Note not found")
}

func main() {
	for {
		fmt.Println("\nNotes Application")
		fmt.Println("1. Add Note")
		fmt.Println("2. List Notes")
		fmt.Println("3. Delete Note")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addNote()
		case 2:
			listNotes()
		case 3:
			deleteNote()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
