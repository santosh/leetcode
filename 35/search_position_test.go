package search

import "testing"

func TestFindIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("test return -1 if not in index", func(t *testing.T) {
		want := -1
		got := FindIndex(arr, 12)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test return 2 for 3", func(t *testing.T) {
		want := 2
		got := FindIndex(arr, 3)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
