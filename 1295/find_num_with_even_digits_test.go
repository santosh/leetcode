package numeven

import "testing"

func TestNumWithEvenDigits(t *testing.T) {
	t.Run("test with [12,345,2,6,7896]", func(t *testing.T) {
		arr := []int{12, 345, 2, 6, 7896}

		want := 2
		got := NumWithEvenDigits(arr)

		assertEquality(t, got, want)
	})

	t.Run("test with [555,901,482,1771]", func(t *testing.T) {
		arr := []int{555,901,482,1771}

		want := 1
		got := NumWithEvenDigits(arr)

		assertEquality(t, got, want)
	})

	t.Run("test with [555,901,482,111]", func(t *testing.T) {
		arr := []int{555,901,482,111}

		want := 0
		got := NumWithEvenDigits(arr)

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
		arr := []int{555,901,482,1771}
		NumWithEvenDigits(arr)
	}
}
