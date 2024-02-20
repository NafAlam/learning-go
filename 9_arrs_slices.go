// package main

// import "fmt"

// func main() {
// 	var arr [2]int
// 	arr[0] = 100
// 	fmt.Println(len(arr))
// 	fmt.Println(arr[0])

// 	arr2 := [2]int{1, 2}
// 	fmt.Println(len(arr2))

// 	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
// 	fmt.Println(len(arr3))

// 	arr4 := [2][2]string{{"hello", "there"}, {"code", "go"}}
// 	test(arr4) // this will not change the arr as in goland the arr is passed as a copy not a reference unlike python
// 	for _, nestedArr := range arr4 {
// 		for _, word := range nestedArr {
// 			fmt.Println(word)
// 		}
// 	}

// 	//-------- slices
// 	// uses the underlying array to make the slice
// 	// you can find out about the length, capacity, pointer

// 	list := [2]int{1, 2}
// 	slicedList := list[1:] // pointer is at list[1], but the length is now changed to the slice and the capcity to 1
// 	fmt.Println(len(slicedList))
// 	// if you change the arr first then the slice will be effected with the change of the underlying arr

// 	// this works as the same as dynamic list in python as the capacity is doubled everytime the elements has reached the capcity
// 	arr10 := []string{}
// 	for i := 0; i < 10; i++ {
// 		arr10 = append(arr10, "hello")
// 		fmt.Println(arr10, len(arr10), cap(arr10))
// 	}

// 	arr11 := make([][]int, 5, 10)
// 	fmt.Println(arr11, len(arr), cap(arr11))

// 	arrBool := []bool{true, false, false}
// 	arrBool2 := []bool{true, true}
// 	arrBool = append(arrBool, arrBool2...)

// 	test2(arrBool)
// 	fmt.Println(arrBool) // this will change the underlying arr as we are passing in a slice as an argument to the func parameter
// }

// func test(x [2][2]string) {
// 	x[0][0] = "new string"
// }

// func test2(x []bool) {
// 	x[0] = false
// }
