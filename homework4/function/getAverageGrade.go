package function

import (
	"homework4/model"
)

func GetAverageGrade(s model.Student) float64 {
	grade := 0.0

	for _, c := range s.Course {
		switch {
		case c.Grade >= 80 && c.Grade <= 100:
			grade += 4
		case c.Grade >= 70 && c.Grade <= 79:
			grade += 3
		case c.Grade >= 60 && c.Grade <= 69:
			grade += 2
		case c.Grade >= 50 && c.Grade <= 59:
			grade += 1
		default:
			grade += 0
		}
	}
	return grade / float64(len(s.Course))
}
