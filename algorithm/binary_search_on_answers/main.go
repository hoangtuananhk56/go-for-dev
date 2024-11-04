package main

import (
	"fmt"
)

// Hàm kiểm tra liệu có tồn tại khoảng con có độ dài `length` mà tổng lớn hơn hoặc bằng target
func IsValid(arr []int, length, target int) bool {
	sum := 0
	for i := 0; i < length; i++ {
		sum += arr[i]
	}
	if sum >= target {
		return true
	}

	for i := length; i < len(arr); i++ {
		sum += arr[i] - arr[i-length]
		if sum >= target {
			return true
		}
	}
	return false
}

// Hàm thực hiện tìm kiếm nhị phân trên đáp án
func MinSubarrayLength(arr []int, target int) int {
	left, right := 1, len(arr)
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if IsValid(arr, mid, target) {
			result = mid    // Nếu tìm được độ dài thỏa mãn, ghi nhận kết quả
			right = mid - 1 // Tiếp tục tìm độ dài ngắn hơn
		} else {
			left = mid + 1 // Tăng độ dài nếu không thỏa mãn
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 15

	minLength := MinSubarrayLength(arr, target)

	if minLength != -1 {
		fmt.Printf("Độ dài ngắn nhất của khoảng con có tổng >= %d là %d\n", target, minLength)
	} else {
		fmt.Printf("Không tồn tại khoảng con nào có tổng >= %d\n", target)
	}
}
