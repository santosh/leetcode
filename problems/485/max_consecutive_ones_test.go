package maxconsec

import "testing"

func TestMaxConsecutiveOnes(t *testing.T) {
	t.Run("test with 1101010111", func(t *testing.T) {
		arr := []int{1, 1, 0, 1, 0, 1, 0, 1, 1, 1}

		want := 3
		got := MaxConsecutiveOnes(arr)

		assertEquality(t, got, want)
	})

	t.Run("test with 1010110010", func(t *testing.T) {
		arr := []int{1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0}

		want := 2
		got := MaxConsecutiveOnes(arr)

		assertEquality(t, got, want)
	})

	t.Run("test with 1110111101", func(t *testing.T) {
		arr := []int{0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1}

		want := 4
		got := MaxConsecutiveOnes(arr)

		assertEquality(t, got, want)
	})
}

func assertEquality(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func BenchmarkMaxConsecutiveOnes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1}
		MaxConsecutiveOnes(arr)
	}
}
