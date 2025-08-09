package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)
/*
type Note struct {
	Title     string // not accessible outside the package
	Content   string // same as above
	CreatedAt time.Time
}
*/
type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
	
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content %v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}