package main

import (
	"log"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenNumber := []int{}
	oddNumber := []int{}
	// loop through list of numbers and sort numbers into even and odd
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumber = append(evenNumber, number)

		} else {
			oddNumber = append(oddNumber, number)
		}
	}
	if len(evenNumber) > 0 {
		log.Println(evenNumber)
	}
	if len(oddNumber) > 0 {
		log.Println(oddNumber)
	}

}
