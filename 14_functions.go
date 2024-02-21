// package main

// import "fmt"

// // this does not need to return x and y, goland just knows that the x and y in the return has been declared
// func namedReturnTypes(a int, b int) (x int, y int) {
// 	x = a
// 	y = b
// 	return
// }

// // variadic function (...)
// func unlimitedNumsSum(nums ...int) int {
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	return sum
// }

// func returnFuncFunction(x int) func(int) int {
// 	return func(y int) int {
// 		return x + y
// 	}
// }

// func callFunc(callable func(string) string, arg string) {
// 	res := callable(arg)
// 	fmt.Println(res)
// }

// func main() {
// 	res := unlimitedNumsSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
// 	fmt.Println(res)

// 	nestedFunc := returnFuncFunction(10)
// 	fmt.Println(nestedFunc(50))

// 	myFunc := func(str string) string {
// 		return str + "hello!"
// 	}
// 	callFunc(myFunc, "world ")
// }
