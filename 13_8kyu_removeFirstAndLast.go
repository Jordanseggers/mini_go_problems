/* Remove First and Last Character

It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string. You're given one parameter, the original string. You don't have to worry with strings with less than two characters.

*/

package kata

func RemoveChar(word string) string {
	if len(word) < 3 {
		return word
	} else {
		return word[1 : len(word)-1]
	}
}
