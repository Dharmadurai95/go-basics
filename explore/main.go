package main

import "fmt"

func main() {
	var numberOne = 13
	var numberTwo = 4
	fmt.Printf("%d + %d = %d \n", numberOne, numberTwo, numberOne+numberTwo)
	fmt.Printf("%d * %d = %d\n", numberOne, numberTwo, numberOne*numberTwo)
	fmt.Printf("%d / %d = %d\n", numberOne, numberTwo, numberOne/numberTwo)
	fmt.Printf("%d remainder %d = %d\n", numberOne, numberTwo, numberOne%numberTwo)

	//type conversion

	var ebeAge int
	var dharmaAge int
	ebeAge = 10
	dharmaAge = 12
	var totalAge float32
	totalAge = float32(ebeAge + dharmaAge)
	fmt.Println(fmt.Sprintf("your total age is around %0.2f", totalAge))

	// findWeekend
	var day = "monday"
	switch day {
	case "sunday", "saturday":
		fmt.Println("weekend")
	case "monday", "Tuesday", "wednesday", "Thursday", "friday":
		fmt.Println("week day")
	default:
		fmt.Println("invalid date")
	}

	switch {
	case 1 > 2:
		fmt.Println("1 > 2")
	default:
		fmt.Println("nothing will print here")
	}

}
