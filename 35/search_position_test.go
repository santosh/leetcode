package search

import "testing"

func TestFindIndex(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 10}

	t.Run("test return index 4 for target 8", func(t *testing.T) {
		want := 4
		got := FindInsertIndex(arr, 8)

		assertIndex(t, got, want)
	})

	t.Run("test return index 1 for target 3", func(t *testing.T) {
		want := 1
		got := FindInsertIndex(arr, 3)

		assertIndex(t, got, want)
	})

	t.Run("test return index 0 for target 0", func(t *testing.T) {
		want := 0
		got := FindInsertIndex(arr, 0)

		assertIndex(t, got, want)
	})
}

func assertIndex(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
