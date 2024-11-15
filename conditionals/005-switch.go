// SWITCH

// Switch statements are used when for executing one condition from multiple conditions.

// It is similar to an if-else-if ladder.

package main

import "fmt"

func switchExample() {
	day := []string{"sunday", "saturday", "wednesday"}

	for _, v := range day {
		switch v {
		case "sunday":
			fmt.Println("Today is Sunday")
		case "saturday":
			fmt.Println("Today is Saturday")
		case "friday":
			fmt.Println("Today is Friday")
		case "thursday":
			fmt.Println("Today is Thursday")
		case "wednesday":
			fmt.Println("Today is Wednesday")
		case "tuesday":
			fmt.Println("Today is Tuesday")
		case "monday":
			fmt.Println("Today is Monday")
		default:
			fmt.Println("Invalid day")
		}
	}
}
