package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"student/display"
	"student/model"
	"student/result"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("====================================")
	fmt.Println("   STUDENT RESULT PROCESSING SYSTEM")
	fmt.Println("====================================")

	fmt.Print("Enter Student Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Register Number: ")
	registerNo, _ := reader.ReadString('\n')
	registerNo = strings.TrimSpace(registerNo)

	fmt.Print("Enter Number of Subjects: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numSubjects, err := strconv.Atoi(input)

	if err != nil || numSubjects <= 0 {
		fmt.Println("Invalid number of subjects.")
		return
	}

	marks := make([]int, numSubjects)

	for i := 0; i < numSubjects; i++ {

		fmt.Printf("Enter marks for Subject %d: ", i+1)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		mark, err := strconv.Atoi(input)

		if err != nil || mark < 0 || mark > 100 {
			fmt.Println("Invalid mark. Enter a value between 0 and 100.")
			return
		}

		marks[i] = mark
	}

	student := model.Student{
		Name:       name,
		RegisterNo: registerNo,
		Marks:      marks,
	}

	total := result.CalculateTotal(student.Marks)
	average := result.CalculateAverage(student.Marks)
	grade := result.CalculateGrade(average)

	status := "PASS"

	if !result.CheckPass(student.Marks) {
		status = "FAIL"
	}

	display.ShowResult(
		student,
		total,
		average,
		grade,
		status,
	)
}
