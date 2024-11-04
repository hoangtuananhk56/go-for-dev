package main

func TwoSum(arr []int, target int) (int, int) {
	low := 0
	high := len(arr) - 1
	for low < high {
		sum := arr[low] + arr[high]
		if sum == target {
			return low, high
		}
		if sum < target {
			low++
		} else {
			high--
		}
	}
	return -1, -1
}

func main() {
	arr := []int{2, 7, 11, 7, 15}
	target := 9
	result1, result2 := TwoSum(arr, target)
	if result1 == -1 {
		println("No two sum solution")
	} else {
		println("Two sum found at index: ", result1, result2)
	}
}
