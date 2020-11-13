package checkout

import (
	"strings"
	"unicode"
)

func Run(input string) int {
	var answer int

	var prices = map[string]int{"A": 50, "B": 30, "C": 20, "D": 15}

	TotalAs := strings.Count(input, "A")
	// BOffer := strings.Count(input, "BB")

	if TotalAs > 0 {
		offerAs := TotalAs / 3
		answer += (130 * offerAs)
		remainingAs := TotalAs - (offerAs * 3)
		answer += (remainingAs * 50)
		input = strings.ReplaceAll(input, "A", "")
	}
	// if BOffer > 0 {
	// 	answer += (45 * BOffer)
	// 	input = strings.ReplaceAll(input, "BB", "")
	// }

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
