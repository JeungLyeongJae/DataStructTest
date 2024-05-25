package t_test

import (
	"fmt"
	"testing"
)

// 二分查找
func binary(arr []int, value int) int {
	var low, high int = 0, len(arr) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] < value {
			low = mid + 1
		} else if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] == value {
			return mid
		} else {
			return -1
		}
	}
	return -1
}

func TestBinary(t *testing.T) {
	fmt.Println(binary([]int{1, 2, 3, 4}, 8))
}
