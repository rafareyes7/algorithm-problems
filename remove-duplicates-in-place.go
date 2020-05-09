package main

import (
	"fmt"
	"reflect"
)

/*
	https://leetcode.com/problems/remove-duplicates-from-sorted-array/
	Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
	Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
	Given nums = [1,1,2],
	Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
	It doesn't matter what you leave beyond the returned length.

	example:
		input: [0,0,1,1,1,2,2,3,3,4]
		output: 5
*/

func removeDuplicatesInPlace(nums []int) int {
	slow := 0
	if len(nums) == 0 {
		return slow
	}

	for fast, num := range nums {
		if nums[slow] != num {
			// fast pointer grows at each iteration.
			slow++
			// if slow = fast -> the array remains the same, but
			// if slow < fast -> slow must be a duplicated element so we
			// assign the nums[fast] and execute the inplace remove.
			nums[slow] = nums[fast]
		}
		// at this point slow pointer does not grow to
		// flag the element that meet the condition we are looking for (duplicated element)
		// so we can process a logic (remove the flagged duplicated element)
		// based on the two pointers.
	}
	// return the slow current index + 1 to represent the length.
	return slow + 1
}

func main() {

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
