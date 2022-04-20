/*
If you can't sleep, just count sheep!!

Task:
Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.
*/

/*
	- create a current round int
	- create an empty string
	- make a for loop that iterates from zero to the input num by increasing the value of current round
		-increase value of current round
		-push "current round sheep..." to string
	- return string
*/

package kata

import (
	"strconv"
)

func countSheep(num int) string {
	var finalString string = ""
	var currentCount int = 0

	for currentCount < num {
		currentCount++
		current := strconv.Itoa(currentCount)
		finalString += current
		finalString += " sheep..."
	}

	return finalString
}
