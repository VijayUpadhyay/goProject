package main

import "fmt"

func main() {
	var grade = ""
	var marks = 80
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	case 50, 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Println("Going for grade check with mark and values as: ", marks, grade)
	switch {
	case grade == "A":
		fmt.Println("Excellent")
	case grade == "B":
		fmt.Println("Good")
	case grade == "C":
		fmt.Println("Passed")
	case grade == "D":
		fmt.Println("Improve")
	case grade == "E":
		fmt.Println("Failed")

	}
	fmt.Println(marks, grade)
}
