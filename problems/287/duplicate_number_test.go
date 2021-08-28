package dupnum

import "testing"

func TestFindDuplicate(t *testing.T) {
	t.Run("one duplicate", func(t *testing.T) {
		nums := []int{1, 3, 4, 2, 2}

		want := 2
		got := FindDuplicate(nums)

		assertNum(t, want, got)
	})
	t.Run("two duplicate", func(t *testing.T) {
		nums := []int{1, 3, 4, 3, 2, 3, 7}

		want := 3
		got := FindDuplicate(nums)

		assertNum(t, want, got)
	})
}

func assertNum(t *testing.T, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}
