package checkout

import "unicode"

func Run(input string) int {
	var answer int

	var prices = map[string]int{"A": 50, "B": 30, "C": 20, "D": 15}

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
