package main

import "testing"

func TestGradeA(t *testing.T) {
	// 80-100 => A, 70-79 => B, 60-69 => C, 50-59 => D, < 50 => F
	//Arrange
	grade := "A"
	score := 85
	//Act
	result := Grade(score)
	//Assert
	if grade != result {
		t.Errorf("Expected result is %v, but got %v." ,grade, result)
	}

}

func TestGradeB(t *testing.T) {
	// 80-100 => A, 70-79 => B, 60-69 => C, 50-59 => D, < 50 => F
	//Arrange
	grade := "B"
	score := 75
	//Act
	result := Grade(score)
	//Assert
	if grade != result {
		t.Errorf("Expected result is %v, but got %v." ,grade, result)
	}

}

func TestGradeC(t *testing.T) {
	// 80-100 => A, 70-79 => B, 60-69 => C, 50-59 => D, < 50 => F
	//Arrange
	grade := "C"
	score := 65
	//Act
	result := Grade(score)
	//Assert
	if grade != result {
		t.Errorf("Expected result is %v, but got %v." ,grade, result)
	}

}

func TestGradeD(t *testing.T) {
	// 80-100 => A, 70-79 => B, 60-69 => C, 50-59 => D, < 50 => F
	//Arrange
	grade := "D"
	score := 55
	//Act
	result := Grade(score)
	//Assert
	if grade != result {
		t.Errorf("Expected result is %v, but got %v." ,grade, result)
	}

}

func TestGradeF(t *testing.T) {
	// 80-100 => A, 70-79 => B, 60-69 => C, 50-59 => D, < 50 => F
	//Arrange
	grade := "F"
	score := 45
	//Act
	result := Grade(score)
	//Assert
	if grade != result {
		t.Errorf("Expected result is %v, but got %v." ,grade, result)
	}

}