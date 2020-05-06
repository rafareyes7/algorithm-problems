package main

import (
	"fmt"
	"reflect"
)

/*
	Imagine that you're a teacher who's just graded the final exam in a class. You
	have a list of student scores on the final exam in a particular order (not
	necessarily sorted), and you want to reward your students. You decide to do so
	fairly by giving them arbitrary rewards following two rules:

	1) All students must receive at least one reward.

	2) Any given student must receive strictly more rewards than an adjacent
	student (a student immediately to the left or to the right) with a lower
	score and must receive strictly fewer rewards than an adjacent student with
	a higher score.

	Write a function that takes in a list of scores and returns the minimum number
	of rewards that you must give out to students to satisfy the two rules.

	You can assume that all students have different scores; in other words, the
	scores are all unique.

	expample:

		input: [8, 4, 2, 1, 3, 6, 7, 9, 5]
		output: 25 // rewards: [4, 3, 2, 1, 2, 3, 4, 5, 1]
*/

// Time: O(n), Space: O(n)
func minRewards(scores []int) int {
	if len(scores) < 2 {
		return 1
	}

	// Initialize all the rewards' array with 1
	// since every student must receive at leat 1 reward.
	rewards := make([]int, len(scores))
	fill(rewards, 1)

	// Get local mins indexes
	localMinIndexes := getLocalMinIndexes(scores)
	for _, idx := range localMinIndexes {
		expandFromLocalMin(idx, scores, rewards)
	}

	return sum(rewards)
}

func fill(arr []int, commonValue int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = commonValue
	}
}

func getLocalMinIndexes(arr []int) []int {
	// Local min is a number which is lower than both numbers on its side.
	if len(arr) < 1 {
		return []int{0}
	}
	localMinIndexes := []int{}
	for i := 0; i < len(arr); i++ {
		// At the start of the array, only right side can be compared.
		if i == 0 && arr[i] < arr[i+1] {
			localMinIndexes = append(localMinIndexes, i)
		}
		// At the end of the array, only left side can be compared.
		if i == len(arr)-1 && arr[i] < arr[i-1] {
			localMinIndexes = append(localMinIndexes, i)
		}
		// To avoid index out of range
		if i == 0 || i == len(arr)-1 {
			continue
		}

		if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
			localMinIndexes = append(localMinIndexes, i)
		}
	}

	return localMinIndexes
}

func expandFromLocalMin(localMinIdx int, scores []int, rewards []int) {
	leftIdx := localMinIdx - 1
	for leftIdx >= 0 && scores[leftIdx] > scores[leftIdx+1] {
		rewards[leftIdx] = max(rewards[leftIdx], rewards[leftIdx+1]+1)
		leftIdx--
	}

	rightIdx := localMinIdx + 1
	for rightIdx < len(scores) && scores[rightIdx] > scores[rightIdx-1] {
		rewards[rightIdx] = rewards[rightIdx-1] + 1
		rightIdx++
	}
}

func max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}

	return value2
}

func sum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}

	return sum
}

func main() {

	want := 25
	got := minRewards([]int{8, 4, 2, 1, 3, 6, 7, 9, 5})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = 1
	got = minRewards([]int{1})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = 3
	got = minRewards([]int{5, 10})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = 3
	got = minRewards([]int{10, 5})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	want = 52
	got = minRewards([]int{2, 20, 13, 12, 11, 8, 4, 3, 1, 5, 6, 7, 9, 0})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}
}
