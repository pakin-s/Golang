package main

import (
	"circleArea/area"
	"fmt"
)

func main() {
	rad := area.CircleArea{
		Ra: 3.00,
	}

	fmt.Printf(`%v`,rad.AreaCal())
}
