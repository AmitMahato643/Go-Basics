package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// keys should always begin with capital letter
type Task struct {
	Id          uint64     `sql:"AUTO_INCREMENT" json:"task_id"`
	UserId      uint64     `json:"userId"`
	Title       string     `json:"task_title"`
	Description string     `json:"task_description"`
	Completed   bool       `json:"task_completed"`
	DueDate     *time.Time `json:"task_dueDate"`
	Created_at  *time.Time `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at"`
}

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}
	filename := arguments[1]
	var task Task
	err := loadFromJSON(filename, &task)
	if err == nil {
		fmt.Println("Your task : ", task)
	} else {
		fmt.Println(err)
	}
}
