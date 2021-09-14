// package is like combination of helper function
// you can't create the same function in one package, example : other-count.go
// but you can create the same function in different package, example : other/count.go
package helper

// You need to use Upper case Title for helper function
func CountNum(num int, times int) int {
	result := num * times
	return result
}
