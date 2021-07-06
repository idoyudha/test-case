package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sumArray([]int{1, 2, 3}))
	fmt.Println(dateFormatConversion("Tue 28 Feb 2019"))
	fmt.Println(calculateFactorial(0))
	fmt.Println(timeConvert(1, 40))
	fmt.Println(minimumDeletion([]int{3, 3, 2, 1, 3}))
}

// Case 1
func sumArray(array []int) int {
	result := 0
	for i := 0; i < len(array); i++ {
		result += array[i]
	}
	return result
}

// Case 2
func dateFormatConversion(date string) string {
	var word = strings.Split(date, " ")
	var monthNumber = 1
	var monthsName = [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	var monthsDate = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Check month string format
	if notContainMonth(monthsName[:], word[2]) {
		return "Please provide the right month format, like ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']"
	}

	for i := 0; i < len(monthsName); i++ {
		if monthsName[i] == word[2] {
			monthNumber += i
		}
	}

	// Check if date above the monthsDate array
	dateNumberInt, _ := strconv.Atoi(word[1])
	if dateNumberInt > monthsDate[monthNumber-1] || dateNumberInt < 1 {
		return "Please provide the rigth date number"
	}

	monthNumberString := strconv.Itoa(monthNumber)
	return word[3] + "-" + monthNumberString + "-" + word[1]
}

func notContainMonth(array []string, str string) bool {
	for _, month := range array {
		if month == str {
			return false
		}
	}
	return true
}

// Case 3
func calculateFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFactorial(n-1)
}

// Case 4
func timeConvert(h int, m int) string {
	hour := findWord(h)
	minutes := findWord(m)
	var result string
	if m <= 30 && m >= 0 && h >= 0 && h <= 24 {
		if m == 0 {
			result = hour + " o'clock"
		} else if m < 30 && m != 15 {
			result = minutes + " minutes past " + hour
		} else if m == 15 {
			result = minutes + " past " + hour
		} else if m == 30 {
			result = "half past " + hour
		}
	} else if m > 30 && m < 60 && h >= 0 && h <= 24 {
		convertMinutes := 60 - m
		hourConverted := findWord(h + 1)
		minutesConverted := findWord(convertMinutes)
		if convertMinutes == 15 {
			result = minutesConverted + " to " + hourConverted
		} else {
			result = minutesConverted + " minutes to " + hourConverted
		}
	} else {
		result = "Please provide rigth number input for hour or minutes"
	}
	return result
}

func findWord(t int) string {
	var timeString = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen", "twenty",
		"twenty-one", "twenty-two", "twenty-three", "twenty-four", "twenty-five", "twenty-six", "twenty-seven", "twenty-eight", "twenty-nine", "thirty"}
	word := ""
	for i := 0; i < len(timeString); i++ {
		if t == 0 {
			continue
		} else if t == i {
			word = timeString[i-1]
		}
	}
	return word
}

// Case 5
func minimumDeletion(array []int) int {

	// Create data of number from array with map
	numbersCounter := make(map[int]int)
	for _, i := range array {
		if _, exist := numbersCounter[i]; exist {
			numbersCounter[i] += 1
		} else {
			numbersCounter[i] = 1
		}
	}

	// Find max counter from data
	maxCount := 0
	for _, value := range numbersCounter {
		if value > maxCount {
			maxCount = value
		}
	}

	return len(array) - maxCount
}
