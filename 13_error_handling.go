// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return -1, errors.New("division by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	// un used import for e.g. is a compile time error

// 	defer fmt.Println("defer") // <- will run at the end of any function even if there is a runtime error it will execute just before
// 	fmt.Println("hello")
// 	// panic("helppp") //<- runtime error

// 	res, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}
// 	fmt.Println(res, err)
// }
