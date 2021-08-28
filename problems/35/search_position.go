// Package search implements a binary search to find index of given
// element in a sorted array.
package search

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.
*/

// FindInsertIndex takes a sorted array arr and a target, and returns
// the position/index inthe array when position is found.
// Otherwise return the position where it wound be inserted
func FindInsertIndex(arr []int, target int) int {
	minIndex := 0
	maxIndex := len(arr) - 1

	return binarySearch(arr, minIndex, maxIndex, target)
}

func binarySearch(arr []int, minIndex, maxIndex, target int) int {
	if maxIndex >= minIndex {
		// most essential part is calculating the midIndex
		midIndex := minIndex + (maxIndex-minIndex)/2

		if arr[midIndex] == target {
			// if target appears to be at midIndex, return midIndex
			return midIndex
		} else if arr[midIndex] < target {
			// if midIndex is less than target, target will be at right
			return binarySearch(arr, midIndex+1, maxIndex, target)
		} else {
			// else target will be at left
			return binarySearch(arr, minIndex, midIndex-1, target)
		}
	}

	// this could have been minIndex or maxIndex,
	// depending upon how 'if conditions' is written
	return minIndex
}
