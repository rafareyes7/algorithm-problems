package main

import (
	"fmt"
	"reflect"
)

/*
	https://leetcode.com/problems/remove-duplicates-from-sorted-array/ && https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/submissions/
	Given a sorted array nums, remove the duplicates in-place such that each element appear at most {appearancesAllowed} and return the new length.
	Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
	Given nums = [1,1,2],
	Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
	It doesn't matter what you leave beyond the returned length.

	example:
		if appearancesAllowed = 1
			input: [0,0,1,1,1,2,2,3,3,4]
			output: 5 // [0, 1, 2, 3, 4]
*/

func removeDuplicatesInPlace(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Since the given array is sorted, the duplicated values are next to each other, then
	// we ignore the first {allowedAppearances} values because there are not more than
	// {allowedAppearances} appearances for each element, so we start checking the array
	// at {allowedAppearances} index.
	allowedAppearances := 1
	// 'slow' pointer will mark the current advance of the result array.
	// 'slow' pointer will mark the index which value is immediately greater than {allowedAppearances}
	// to be replaced for the next value in the order (since the array is sorted)
	// e.g: allowedAppearances = 1; then [0, 1, slow:1, fast:2, 3]
	// At the begining both pointer starts the same but when the check condition (line:42) is not
	// meet, the 'slow' pointer will not grow serving as the mark we mentioned before.
	slow := allowedAppearances
	for fast := slow; fast < len(nums); fast++ {
		// At each iteration the 'fast' pointer will check if the current value is different
		// from {allowedAppearances} indexes behind the 'slow' pointer.
		// It's like keep checking the 'slow' pointer tail that represent a interval of
		// {allowedAppearances} spots.
		if nums[slow-allowedAppearances] != nums[fast] {
			// when slow's pointer tail is different of the current number, we perform an in place edit,
			// based on the index that needs to be edited marked by the 'slow' pointer
			nums[slow] = nums[fast]
			// After the in place edition, 'slow' pointer grows because the backward element has been corrected.
			slow++
		}

		// At this point if the check condition (line:42) was not meet, it means that in this iteration we have
		// found more appearances of a number than the {allowedAppearances}, so 'slow' pointer did not grow
		// to mark the index that needs to be edited with its immediately greater number which will be found by
		// the 'pointer' eventually based on check condition (line:42)
	}
	// We return the 'slow' index because as we mentioned before, 'slow' pointer has been marked the advance or the
	// progress of the result array, while the 'fast' pointer was checking each value in the array agains the 'slow'
	// pointer tail
	fmt.Printf("%v\n", nums[:slow])
	return slow
}

func main() {

	// test cases for appearancesAllowed := 1
	want := 0
	got := removeDuplicatesInPlace([]int{})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = 2
	got = removeDuplicatesInPlace([]int{1, 1, 2})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = 5
	got = removeDuplicatesInPlace([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}
}
