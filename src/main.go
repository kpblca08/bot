package main

import (
	"encoding/json"
	"fmt"
)

type Lecture struct {
	Name          string `json:"name"`
	IsWeekly      bool   `json:"isWeekly"`
	DaysOfWeek    uint8  `json:"dayOfWeek"`
	TimeStartMin  int    `json:"timeStart"`
	TimeFinishMin int    `json:"timeFinish"`
	ClassNumber   string `json:"classNumber"`
	TeacherName   string `json:"teacherName"`
}

func main() {
	str := Lecture{"kek", true, 2, 5, 8}
	// fmt.Print(str)
	a, _ := json.MarshalIndent(str, "", " ")

	fmt.Println(string(a))
}
