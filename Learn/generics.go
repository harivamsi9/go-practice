// package main

// import (
// 	"fmt"
// )

// // Create a generic function max
// // max() should accept int, float 64, string and returns max(slice)

// type ValidTypes interface {
// 	int | float64 | string
// }

// func max[T ValidTypes](values []T) (T, error) {
// 	if len(values) == 0 {
// 		var zero T
// 		return zero, fmt.Errorf("is the max for empty slice")
// 	}

// 	maxi := values[0]
// 	for _, val := range values {
// 		if val > maxi {
// 			maxi = val
// 		}
// 	}
// 	return maxi, nil

// }

// func main() {
// 	fmt.Println(max([]string{"A", "Z", "B", "C"}))
// 	fmt.Println(max([]int{1, 2, 3, 4, 2, 3, 5, 4, 5}))
// }
