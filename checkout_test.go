package checkout

import "testing"

func TestRun(t *testing.T) {

	t.Run("Test for first acceptance criteria", func(t *testing.T) {
		got := Run("aBc")
		want := -1
		assertEquals(t, got, want)
	})

	t.Run("Test for second acceptance criteria", func(t *testing.T) {
		got := Run("-B8x")
		want := -1
		assertEquals(t, got, want)
	})
}

func assertEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
