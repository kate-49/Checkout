package checkout

import "unicode"

func Run(input string) int {
	var answer int

	for _, letter := range input {
		if unicode.IsLower(letter) {
			answer = -1
		}
	}
	return answer
}
