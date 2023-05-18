package main

func Grade(num int) string {
	if num < 0 {
		return "You just can't have a negative score, who are you!!?"
	} 

	switch {
	case (num > 79 && num < 101) :
		return "A"
	case (num > 69 && num < 80) :
		return "B"
	case (num > 59 && num < 70) :
		return "C"
	case (num > 49 && num < 60) :
		return "D"
	case (num < 50) :
		return "F"
	}

	return "You just can't have a score morethan 100, who are you!!?"
}