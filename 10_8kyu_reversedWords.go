/* Reversed Words
Complete the solution so that it reverses all of the words within the string passed in.

Example:

"The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"

-create empty string to hold the final sentence
-split string into words array
-loop backwards through words array
	-add each word to final sentence
-return final sentence
*/

package kata

import "strings"

func ReverseWords(str string) string {
	split := strings.Split(str, " ")
	reversed := split[len(split)-1]
	for i := len(split) - 2; i >= 0; i-- {
		reversed += " "
		reversed += split[i]
	}
	return reversed
}

// best practice solution below

package kata

import "strings"

func ReverseWords(str string) string {
  words := strings.Split(str, " ")
  for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
  }
  return strings.Join(words, " ")
} 