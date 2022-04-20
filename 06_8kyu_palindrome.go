/* Is it a palindrome?

Write a function that checks if a given string (case insensitive) is a palindrome.

*/

package kata

import "strings"

func IsPalindrome(str string) bool {
	lowercase := strings.ToLower(str)
	letters := strings.Split(lowercase, "")
	var reversed string = ""
	for x := len(letters) - 1; x >= 0; x-- {
		reversed += letters[x]
	}
	return reversed == lowercase
}
