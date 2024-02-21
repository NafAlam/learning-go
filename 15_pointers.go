// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func changeName(p *Person, name string) {
// 	// (*p).name = name -----> or you can explicitally de reference this way
// 	p.name = name // if you use a dot notation in a struct, goland will automatically de reference the pointer for you so you dont need to use (*)
// }

// func main() {
// 	//& is the address of the value
// 	//* de-reference the pointer to access the value

// 	x := 10
// 	y := &x
// 	*y = 999
// 	print(x)
// 	print("--------")
// 	print(*y)

// 	a := 0
// 	change(&a) // passing a pointer to a

// 	str := []string{"hello", "world"}
// 	pointerToStr := &str
// 	pointerToThePointer := &pointerToStr
// 	fmt.Println(str, *pointerToStr, **pointerToThePointer)

// 	person := Person{name: "naf", age: 23}
// 	changeName(&person, "nafizul")
// 	fmt.Println(person)

// 	hi := 0
// 	lol := 1
// 	arr := []*int{&hi, &lol} // you can hold a list of pointers + using the (&) is optional as without it goland will auto intiallise the pointers for you
// 	fmt.Println(arr)
// }

// // this func is accepting a pointer to an int
// func change(x *int) {
// 	*x = 4
// }
