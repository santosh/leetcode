package reverseint

import "testing"

func TestReverse(t *testing.T) {
	t.Run("test 123 for 321", func(t *testing.T) {
		want := 321
		got := Reverse(123)

		assertNumber(t, got, want)
	})

	t.Run("test -123 for -321", func(t *testing.T) {
		want := -321
		got := Reverse(-123)

		assertNumber(t, got, want)
	})

	t.Run("test 1200 for 21", func(t *testing.T) {
		want := 21
		got := Reverse(1200)

		assertNumber(t, got, want)
	})
}

func assertNumber(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
