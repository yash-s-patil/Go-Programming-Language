package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving the note failed")
		return err
	}
	fmt.Printf("saving note succeeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	// to read long text string
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
