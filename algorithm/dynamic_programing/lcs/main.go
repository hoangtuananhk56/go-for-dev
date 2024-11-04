// package main

// import (
// 	"fmt"
// )

// // Hàm tìm độ dài và nội dung của chuỗi con chung dài nhất
// func LCS(X, Y string) (int, string) {
// 	m, n := len(X), len(Y)
// 	// Khởi tạo ma trận DP để lưu độ dài của LCS
// 	dp := make([][]int, m+1)
// 	for i := range dp {
// 		dp[i] = make([]int, n+1)
// 	}

// 	// Điền các giá trị cho ma trận dp
// 	for i := 1; i <= m; i++ {
// 		for j := 1; j <= n; j++ {
// 			if X[i-1] == Y[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 			} else {
// 				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
// 			}
// 		}
// 	}

// 	// Truy vết ngược lại để tìm chuỗi LCS
// 	lcs := ""
// 	i, j := m, n
// 	for i > 0 && j > 0 {
// 		if X[i-1] == Y[j-1] {
// 			lcs = string(X[i-1]) + lcs
// 			i--
// 			j--
// 		} else if dp[i-1][j] > dp[i][j-1] {
// 			i--
// 		} else {
// 			j--
// 		}
// 	}

// 	return dp[m][n], lcs
// }

// // Hàm tiện ích để tìm giá trị lớn hơn giữa hai số nguyên
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	X := "AGGTAB"
// 	Y := "GXTXAYB"
// 	length, lcs := LCS(X, Y)
// 	fmt.Printf("Độ dài của chuỗi con chung dài nhất là %d\n", length)
// 	fmt.Printf("Chuỗi con chung dài nhất là: %s\n", lcs)
// }

// Chuỗi con chung dài nhất liên tiếp
// package main

// import (
// 	"fmt"
// )

// // Hàm tìm độ dài và nội dung của chuỗi con chung dài nhất liên tiếp
// func LongestCommonSubstring(X, Y string) (int, string) {
// 	m, n := len(X), len(Y)
// 	// Khởi tạo ma trận DP để lưu độ dài của chuỗi con chung liên tiếp
// 	dp := make([][]int, m+1)
// 	for i := range dp {
// 		dp[i] = make([]int, n+1)
// 	}

// 	maxLength := 0 // Độ dài của chuỗi con chung dài nhất liên tiếp
// 	endIndex := 0  // Vị trí kết thúc của chuỗi con chung dài nhất

// 	// Điền các giá trị cho ma trận dp
// 	for i := 1; i <= m; i++ {
// 		for j := 1; j <= n; j++ {
// 			if X[i-1] == Y[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 				if dp[i][j] > maxLength {
// 					maxLength = dp[i][j]
// 					endIndex = i
// 				}
// 			} else {
// 				dp[i][j] = 0
// 			}
// 		}
// 	}

// 	// Trích xuất chuỗi con chung dài nhất liên tiếp
// 	longestSubstring := X[endIndex-maxLength : endIndex]
// 	return maxLength, longestSubstring
// }

// func main() {
// 	X := "ABABC"
// 	Y := "BABCA"
// 	length, substring := LongestCommonSubstring(X, Y)
// 	fmt.Printf("Độ dài của chuỗi con chung dài nhất liên tiếp là %d\n", length)
// 	fmt.Printf("Chuỗi con chung dài nhất liên tiếp là: %s\n", substring)
// }

package main

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	result := 0
	for _, v := range dp {
		result = max(result, v)
	}

	return result
}

func main() {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	println(maxEnvelopes(envelopes))
}
