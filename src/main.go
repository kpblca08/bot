package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Lectures struct {
	Lectures []Lecture `json:"Lecture"`
}

type Lecture struct {
	Name         string `json:"name"`
	IsWeekly     uint8  `json:"isWeekly"`
	TimeStartMin int    `json:"timeStart"`
	ClassNumber  string `json:"classNumber"`
	TeacherName  string `json:"teacherName"`
}

func newLecture() Lecture {
	var L Lecture
	fmt.Print("type name: ")
	fmt.Scanf("%s", &L.Name)
	fmt.Print("\n1 - every week, 2 - every 2 weeks, 0 - once: ")
	fmt.Scanf("%d", &L.IsWeekly)
	var hour, min int
	var check rune
	for {
		fmt.Print("when it starts (hh:mm): ")
		scanned, _ := fmt.Scanf("%d%c%d", &hour, &check, &min)
		if (check == ':' || check == '-' || check == ' ') && (scanned == 3) {
			break
		}
		fmt.Println("cringe input, try again")
	}
	fmt.Print("where banana: ")
	fmt.Scanf("%s", &L.ClassNumber)
	fmt.Print("who: ")
	fmt.Scanf("%s", &L.TeacherName)
	return L
}

func readInstructions(File *os.File) Lectures {
	byteValue, _ := ioutil.ReadAll(File)
	var Lectures Lectures
	err := json.Unmarshal(byteValue, &Lectures)
	if err != nil {
		fmt.Println(err)
	}
	return Lectures
}

func main() {
	File, err := os.Open("instructions.json")
	if err != nil {
		File, err = os.Create("instructions.json")
	}
	defer File.Close()
	Lectures := readInstructions(File) //make([]Lecture, 10, 10)
	lec := Lecture{
		"Русский",
		1,
		720,
		"123",
		"Vestyak",
	}
	Lectures.Lectures = append(Lectures.Lectures, lec)
	Lectures.Lectures = append(Lectures.Lectures, lec)
	a, _ := json.MarshalIndent(Lectures, "", " ")
	_, err = File.WriteString(string(a))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(a))
}
