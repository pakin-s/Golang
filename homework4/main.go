package main

import (
	"fmt"
	"homework4/function"
	"homework4/model"
)

func main() {
	john := model.Student{
		Name: "John Doe",
		Age:  18,
	}

	function.AddGrade(&john, "Math", 85.5)
	function.AddGrade(&john, "English", 92.0)
  	function.AddGrade(&john, "Science", 78.0)

	average := function.GetAverageGrade(john)

	fmt.Printf("Student: %s\n", john.Name)
  	fmt.Printf("Age: %d\n", john.Age)
 	fmt.Printf("Grades: %v\n", john.Course)
  	fmt.Printf("Average Grade: %.2f\n", average)

}
