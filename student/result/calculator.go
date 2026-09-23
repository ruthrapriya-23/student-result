package result

func CalculateTotal(marks []int) int {
	total := 0

	for _, mark := range marks {
		total += mark
	}

	return total
}

func CalculateAverage(marks []int) float64 {
	if len(marks) == 0 {
		return 0
	}

	total := CalculateTotal(marks)

	return float64(total) / float64(len(marks))
}

func CalculateGrade(average float64) string {

	if average >= 90 {
		return "A+"
	} else if average >= 80 {
		return "A"
	} else if average >= 70 {
		return "B"
	} else if average >= 60 {
		return "C"
	} else if average >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func CheckPass(marks []int) bool {

	for _, mark := range marks {
		if mark < 40 {
			return false
		}
	}

	return true
}
