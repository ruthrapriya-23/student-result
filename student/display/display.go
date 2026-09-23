package display

import (
	"fmt"

	"student/model"
)

func ShowResult(
	student model.Student,
	total int,
	average float64,
	grade string,
	status string,
) {
	fmt.Println("\n====================================")
	fmt.Println("        STUDENT RESULT")
	fmt.Println("====================================")

	fmt.Println("Name        :", student.Name)
	fmt.Println("Register No :", student.RegisterNo)

	fmt.Println("------------------------------------")

	for i, mark := range student.Marks {
		fmt.Printf("Subject %d    : %d\n", i+1, mark)
	}

	fmt.Println("------------------------------------")

	fmt.Println("Total       :", total)
	fmt.Printf("Average     : %.2f\n", average)
	fmt.Println("Grade       :", grade)
	fmt.Println("Status      :", status)

	fmt.Println("====================================")
}
