package main

/**BinarySearch function searchs an item via binary search**/
func BinarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == item {
			return mid
		}
		if arr[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	item := 10
	result := BinarySearch(arr, item)
	if result == -1 {
		println("Item not found in array")
	} else {
		println("Item found at index: ", result)
	}
}
