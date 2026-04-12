package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"jaytailor.com/note/note"
	"jaytailor.com/note/todo"
)

func main() {
	title, content := getNoteData()

	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Error saving todo: ", err)
		return
	}
	fmt.Println("Saving todo success")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note: ", err)
		return
	}
	fmt.Println("Saving note success")
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
