package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("return [0, 1] for 9 in [2, 7, 11, 15]", func(t *testing.T) {
		arr := []int{2, 7, 11, 15}
		want := []int{0, 1}
		got := TwoSum(arr, 9)

		assertArray(t, got, want)
	})

	t.Run("return [1, 2] for 6 in [3, 2, 4]", func(t *testing.T) {
		arr := []int{3, 2, 4}
		want := []int{1, 2}
		got := TwoSum(arr, 6)

		assertArray(t, got, want)
	})

	t.Run("return [0, 1] for 6 in [3, 3]", func(t *testing.T) {
		arr := []int{3, 3}
		want := []int{0, 1}
		got := TwoSum(arr, 6)

		assertArray(t, got, want)
	})
	
}


func assertArray(t *testing.T, got, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
