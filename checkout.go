package checkout

import (
	"strings"
	"unicode"
)

func Run(input string) int {
	var answer int

	var prices = map[string]int{"A": 50, "B": 30, "C": 20, "D": 15}

	TotalAs := strings.Count(input, "A")
	TotalBs := strings.Count(input, "B")

	if TotalAs > 0 {
		offerAs := TotalAs / 3
		answer += (130 * offerAs)
		remainingAs := TotalAs - (offerAs * 3)
		answer += (remainingAs * 50)
		input = strings.ReplaceAll(input, "A", "")
	}
	if TotalBs > 0 {
		offerBs := TotalBs / 2
		answer += (45 * offerBs)
		remainingBs := TotalBs - (offerBs * 2)
		answer += (remainingBs * 30)
		input = strings.ReplaceAll(input, "B", "")
	}

	for _, letter := range input {
		if unicode.IsLower(letter) {
			answer = -1
		} else {
			for priceLetter, price := range prices {
				stringletter := string(letter)
				if priceLetter == stringletter {
					answer += price

				}
			}
		}
	}
	return answer
}
