// package main

// import "fmt"

// type Clothing struct {
// 	shoeSize   uint
// 	shirtColor string
// }

// // this struct is 'public' while the properties arent as they are not capitalised
// type Person struct {
// 	name     string
// 	age      uint
// 	clothing Clothing
// }

// func (p Person) GetName() string {
// 	return p.name
// }

// func (p Person) SetName(name string) {
// 	p.name = name
// }

// // first CAPITAL letter named exports this function outside of the package name e.g. main
// func Test() string {
// 	return ""
// }

// func main() {
// 	p1 := Person{name: "nafizul", age: 23, clothing: Clothing{12, "blue"}}
// 	p1.GetName()
// 	p1.SetName("nafizul alam") // this will not change the name as we are passing in a copy of this person struct not reference
// 	fmt.Println(p1.name, p1.age, p1.clothing.shirtColor, p1.clothing.shoeSize)
// }
