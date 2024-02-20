// package main

// import "fmt"

// func main() {
// 	a := 1.2
// 	b := 1
// 	c := a < float64(b) && b < 1.0
// 	fmt.Println(c)

// 	if a < 2 {
// 		fmt.Println("hi")
// 	} else if a == 3 {
// 		fmt.Println("hi")
// 	} else {
// 		fmt.Println("hi")
// 	}

// 	switch a {
// 	case 1:
// 		fmt.Println("hi")
// 	case 2:
// 		fmt.Println("hi")
// 	case 3:
// 		fmt.Println("hi")
// 	default:
// 		fmt.Println("default")
// 	}

// 	switch x := 2; {
// 	case x < 1:
// 		fmt.Println("hi")
// 		fallthrough // this executes the case below without chekcing the conditional
// 	case x < 2:
// 		fmt.Println("hi")
// 	case x < 3:
// 		fmt.Println("hi")
// 	default:
// 		fmt.Println("default")
// 	}

// 	y := 'h'
// 	switch y {
// 	case 'h', 'a', 'j':
// 		fmt.Println("it has worked")
// 	}
// }
