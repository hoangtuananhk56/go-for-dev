package main

import (
	"fmt"
)

// Hàm tìm tổng của dãy con có tổng lớn nhất và trả về cả dãy con
func MaxSubArray(arr []int) (int, []int) {
	currentMax := arr[0]
	globalMax := arr[0]
	start, end, tempStart := 0, 0, 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > currentMax+arr[i] {
			currentMax = arr[i]
			tempStart = i
		} else {
			currentMax += arr[i]
		}

		// Cập nhật globalMax và vị trí bắt đầu/kết thúc của dãy con
		if currentMax > globalMax {
			globalMax = currentMax
			start = tempStart
			end = i
		}
	}

	// Trích xuất dãy con có tổng lớn nhất
	maxSubArray := arr[start : end+1]
	return globalMax, maxSubArray
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum, subArray := MaxSubArray(arr)
	fmt.Printf("Tổng của dãy con có tổng lớn nhất là %d\n", maxSum)
	fmt.Printf("Dãy con có tổng lớn nhất là: %v\n", subArray)
}
