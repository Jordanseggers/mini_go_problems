/* Simple string characters

In this Kata, you will be given a string and your task will be to return a list of ints detailing the count of uppercase letters, lowercase, numbers and special characters, as follows.

Solve("*'&ABCDabcde12345") = [4,5,5,3].
--the order is: uppercase letters, lowercase, numbers and special characters.
More examples in the test cases.

Good luck!

-make tally array and initialize to 4 zeros
-split string into chars
-range over chars
	-tally each
-return array

*/

package kata

import "unicode"

func Solve(s string) []int {
	var tally []int = []int{0, 0, 0, 0}

	for _, char := range s {
		if unicode.IsUpper(char) {
			tally[0] += 1
		} else if unicode.IsLower(char) {
			tally[1] += 1
		} else if unicode.IsDigit(char) {
			tally[2] += 1
		} else {
			tally[3] += 1
		}
	}

	return tally
}
