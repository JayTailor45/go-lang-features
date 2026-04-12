package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"jaytailor.com/note/note"
	"jaytailor.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()

	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo text:")
	return text
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	fmt.Println("Data saved successfully")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return data.Save()
}

func saveAnyData(data any) {
}

func saveAnyDataWithInterface(data interface{}) {
}
