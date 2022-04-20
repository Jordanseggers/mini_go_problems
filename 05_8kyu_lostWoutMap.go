/* Beginner - Lost Without a Map

Given an array of integers, return a new array with each value doubled.

For example:

[1, 2, 3] --> [2, 4, 6]
*/

package kata

func Maps(x []int) []int {
	var doubled []int = make([]int, len(x))

	for idx, elem := range x {
		doubled[idx] = elem * 2
	}

	return doubled
}
