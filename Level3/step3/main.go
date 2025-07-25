package main

import (
	"errors"
	"fmt"
)

var (
	ErrorArgs = errors.New("Некорректые аргументы")
)

func UnderLimit(nums []int, limit, n int) ([]int, error) {
	if n <= 0 {
		return nil, ErrorArgs
	}
	if nums == nil {
		return nil, ErrorArgs
	}

	firstN := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] < limit {
			firstN = append(firstN, nums[i])
			n--
		}
		if n == 0 {
			break
		}
	}

	return firstN, nil
}

func Mix(nums []int) []int {
	length := len(nums)
	nums1 := nums[:length/2]
	nums2 := nums[length/2:]
	new := make([]int, length)

	var j int = 0
	for i := 0; i < length; i += 2 {
		new[i] = nums1[j]
		new[i+1] = nums2[j]
		j++
	}

	return new
}

func Join(nums1, nums2 []int) []int {
	newNums := make([]int, 0, len(nums1)+len(nums2))
	newNums = append(newNums, nums1...)
	newNums = append(newNums, nums2...)
	return newNums
}

func SliceCopy(nums []int) []int {
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	return newNums
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(Mix(nums))
}
