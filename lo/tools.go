package lo

import "fmt"

func printResult[T any, R any](input T, result R) {
	fmt.Println("input:", input)
	fmt.Println("result:", result)
}
