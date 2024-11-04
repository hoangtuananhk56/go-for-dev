package main

import (
	"fmt"
	"sort"
)

// Hàm tính độ dài của dãy con tăng dài nhất
func LongestIncreasingSubsequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// Mảng phụ lưu các phần tử kết thúc của các dãy con tăng
	lis := []int{}

	for _, num := range arr {
		// Tìm vị trí thích hợp để thay thế hoặc chèn phần tử
		idx := sort.SearchInts(lis, num)

		// Nếu không tìm thấy, thêm phần tử vào cuối mảng
		if idx == len(lis) {
			lis = append(lis, num)
		} else {
			// Thay thế phần tử tại vị trí tìm thấy
			lis[idx] = num
		}
	}

	// Độ dài của mảng phụ là độ dài của LIS
	fmt.Println(lis)
	return len(lis)
}

func main() {
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	fmt.Printf("Độ dài của dãy con tăng dài nhất là %d\n", LongestIncreasingSubsequence(arr))
}
