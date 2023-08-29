package main

import (
	"fmt"
	"math/rand"
)

var rolled = func(slides int) int {
	return rand.Intn(slides) + 1
}

func main() {
	rolls, dice, slides := 2, 2, 6
	for i := 1; i <= rolls; i++ {
		sum := 0
		for j := 1; j <= dice; j++ {
			rolled := rolled(slides)
			sum += rolled
			fmt.Printf("Rolling # %d, Dice # %d: %d\n", i, j, rolled)

		}
		switch {
		case sum == 0:
			fmt.Printf(`Sum of dice: '%d' "Snake eyes"`+"\n\n", sum)
		case sum == 7:
			fmt.Printf(`Sum of dice: '%d' "Lucky 7"`+"\n\n", sum)
		case sum%2 == 0:
			fmt.Printf(`Sum of dice: '%d' "Even"`+"\n\n", sum)
		case sum%2 == 1:
			fmt.Printf(`Sum of dice: '%d' "Odd"`+"\n\n", sum)
		}

	}

}
