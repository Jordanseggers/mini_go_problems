/*
Write a method, that will get an integer array as parameter and will process every number from this array.

Return a new array with processing every number of the input-array like this:

If the number has an integer square root, take this, otherwise square the number.

Example
[4,3,9,7,2,1] -> [2,9,3,49,4,1]
Notes
The input array will always contain only positive numbers, and will never be empty or null.
*/

/*
 - create new slice the length of the input slice
 - loop through input slice
		- if value has a square root, add sq root to new slice
				if math.Sqrt(x)
		- if it doesn't, add squared version to slice
*/

package main

import "math"

func squareRootOrSquare(arr []int) []int {
	sqrArr := make([]int, len(arr))

	for idx, elem := range arr {
		sqrt := math.Sqrt(float64(elem))

		if sqrt == math.Floor(sqrt) {
			sqrArr[idx] = int(sqrt)
		} else {
			sqrArr[idx] = elem * elem
		}
	}
	return sqrArr
}

func main() {

}
