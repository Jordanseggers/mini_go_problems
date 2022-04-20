/* Count of positives / sum of negatives

Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

Example
For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65]

*/

package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int = []int{0, 0}

	if len(numbers) < 1 {
		return make([]int, 0)
	}

	for _, num := range numbers {
		if num > 0 {
			res[0] += 1
		} else if num < 0 {
			res[1] += num
		}
	}
	return res // your code here
}
