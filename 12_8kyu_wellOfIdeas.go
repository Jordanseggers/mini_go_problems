/* Well of Ideas - Easy Version

For every good kata idea there seem to be quite a few bad ones!

In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. If there are one or two good ideas, return 'Publish!', if there are more than 2 return 'I smell a series!'. If there are no good ideas, as is often the case, return 'Fail!'.
*/

package kata

func Well(x []string) string {
	tally := 0
	for _, elem := range x {
		if elem == "good" {
			tally += 1
			if tally > 2 {
				return "I smell a series!"
			}
		}
	}
	if tally > 0 {
		return "Publish!"
	} else {
		return "Fail!"
	}
}
