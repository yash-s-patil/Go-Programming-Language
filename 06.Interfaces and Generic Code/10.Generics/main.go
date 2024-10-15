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

// embedded interface
type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.54)
	printSomething("Hello")

	result := add(4, 5)
	fmt.Println("Result:", result)

	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	printSomething(todo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

// generic add function
func add[T int | float64 | string](a, b T) T {
	return a + b
}

// any value allowed
// interface{} or any keyword
func printSomething(value any) {
	// alternative
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float64:", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}

	/*
		switch value.(type) {
		case int:
			fmt.Println("Integer:", value)
		case float64:
			fmt.Println("Float64:", value)
		case string:
			fmt.Println("String:", value)
		}
	*/
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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
