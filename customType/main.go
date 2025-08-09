package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/custom-type/note"
	"example.com/custom-type/todo"
)

/*
type str string

func (text *str) log() {
	fmt.Print(*text)
}

func main() {
	var name str = "pratik"
	name.log()
}

*/

// interface
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	todoText := getUserInput("Todo text: ")
	
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = outputData(todo)
	if err != nil {
		return;
	}

	title, content:= getNoteData()
	userNote, err := note.New(title, content)
	outputData(userNote)
	
	if err != nil {
		return
	}


	printSomething(1)
	printSomething(1.5)
	printSomething("apple")
}

func printSomething(value any) {
	typedVal, ok := value.(int)
	
	if ok {
		ans := typedVal + 1
		fmt.Println(ans)
		return
	}
	typedValFloat, ok := value.(float64)

	if ok {
		fmt.Println(typedValFloat + 2.3)
	}
	/*
	// type swtich
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case string:
		fmt.Println("String: ", value) 
	case float64:
		fmt.Println("Float 64", value)
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
		fmt.Println("Saving failed")
		return err
	}
	
	fmt.Println("Saved successfully!")
	return nil
}



// input -----------------------------------
func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	/*
	var value string
	fmt.Scanln(&value) // leaves a \n in the buffer next scan consumes it
	*/
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

