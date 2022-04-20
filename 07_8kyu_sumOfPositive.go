/* Sum of positive

You get an array of numbers, return the sum of all of the positives ones.

Example [1,-4,7,12] => 1 + 7 + 12 = 20

Note: if there is nothing to sum, the sum is default to 0.
*/

package kata

func PositiveSum(numbers []int) int {
	sum := 0
	for _, elem := range numbers {
		if elem > 0 {
			sum += elem
		}
	}
	return sum
}
