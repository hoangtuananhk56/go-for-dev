package main

import (
	"fmt"
)

// Hàm kiểm tra liệu có thể hoàn thành công việc trong `days` ngày
// với số lượng công việc tối đa một ngày là `maxWork`
func CanComplete(jobs []int, days, maxWork int) bool {
	currentWork := 0
	dayCount := 1

	for _, job := range jobs {
		// Nếu một công việc lớn hơn `maxWork`, không thể hoàn thành
		if job > maxWork {
			return false
		}

		// Nếu cộng công việc này vượt quá `maxWork`, tăng ngày và reset công việc
		if currentWork+job > maxWork {
			dayCount++
			currentWork = job
			if dayCount > days {
				return false
			}
		} else {
			currentWork += job
		}
	}

	return true
}

// Hàm thực hiện tìm kiếm nhị phân trên khoảng giá trị
func MinimizeMaxWork(jobs []int, days int) int {
	left, right := 1, 0
	for _, job := range jobs {
		right += job // Tổng công việc là giới hạn trên
	}

	result := right

	for left <= right {
		mid := left + (right-left)/2

		// Kiểm tra liệu có thể hoàn thành công việc với `mid` là số lượng công việc tối đa trong một ngày
		if CanComplete(jobs, days, mid) {
			result = mid    // Nếu có thể, ghi nhận kết quả và thử tìm giá trị nhỏ hơn
			right = mid - 1 // Tìm kiếm trong nửa trái
		} else {
			left = mid + 1 // Nếu không thể, tăng `mid`
		}
	}

	return result
}

func main() {
	jobs := []int{10, 20, 30, 40, 50}
	days := 3

	minMaxWork := MinimizeMaxWork(jobs, days)

	fmt.Printf("Số lượng công việc lớn nhất ít nhất mà một ngày phải làm là %d\n", minMaxWork)
}
