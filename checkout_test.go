package checkout

import "testing"

func TestLettersHaveCorrectValues(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		// acceptance criteria tests
		{input: "aBc", want: -1},
		{input: "-B8x", want: -1},
		{input: "AA", want: 100},
		{input: "ABCD", want: 115},
		{input: "AAA", want: 130},
		{input: "AAAAAA", want: 260},

		//other tests for B offer
		{input: "BB", want: 45},
		{input: "BBBB", want: 90},

		//tests for integrated offers and non offer items
		{input: "ACACA", want: 170},
		{input: "BBDBDB", want: 120},
	}

	for _, tc := range tests {
		got := Run(tc.input)
		if tc.want != got {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
