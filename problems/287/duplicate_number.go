package dupnum

/* Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one. */

// FindDuplicate uses binary search to find duplicate numbers.
func FindDuplicate(nums []int) int {
	start, end := 1, len(nums)

	for start <= end {
		mid := start + (end-start)/2
		cnt := 0
		for _, val := range nums {
			if val <= mid {
				cnt++
			}
		}

		// update end or start to divide searching
		if cnt > mid {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}
