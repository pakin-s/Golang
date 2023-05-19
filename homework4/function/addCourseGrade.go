package function

import (
	"homework4/model"
)

func AddGrade(s *model.Student, CourseName string, Grade float64) {
	s.Course = append(s.Course, model.Course{CourseName: CourseName, Grade: Grade})
}
