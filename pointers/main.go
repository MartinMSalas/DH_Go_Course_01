package main

import "fmt"


func main() {
	/*
	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println("The value of i is: ", i)
	fmt.Println("The value of i by Pointer is: ", *iP)
	fmt.Println("The address of i is: ", &i)
	fmt.Println("The address of i by Pointer is: ", iP)


	*iP = 21
	fmt.Println("The value of i is: ", i)
	fmt.Println("The value of i by Pointer is: ", *iP)

	myVar := 30

	fmt.Printf("The value of myVar is: %d and the direction is %p \n ", myVar, &myVar)
	*/
	//myVar := 30
	//increment(myVar)

	//fmt.Printf("The value of myVar is: %d and the direction is %p \n ", myVar, &myVar)
/*
	incrementP(&myVar)
	fmt.Printf("The value of myVar is: %d and the direction is %p \n ", myVar, &myVar)
	*/

	/*
	mySlice := []int {3,4,2}
	fmt.Printf("The value of mySlice is: %v and the direction is %p \n ", mySlice, &mySlice)
	fmt.Printf("The value of mySlice[0] is: %v and the direction is %p \n ", mySlice[0], &mySlice[0])
	fmt.Printf("The value of mySlice[1] is: %v and the direction is %p \n ", mySlice[1], &mySlice[1])
	fmt.Printf("The value of mySlice[2] is: %v and the direction is %p \n ", mySlice[2], &mySlice[2])
	updateSlice(mySlice)
*/
	myStruct := &MyStruct{ID: 1, Name: "John"}

	fmt.Printf("The address of myStruct is: %p \n", myStruct)
	fmt.Printf("The value of myStruct is: %d , name %s \n", myStruct.ID, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("The address of myStruct is: %p \n", myStruct)
	fmt.Printf("The value of myStruct is: %d , name %s \n", myStruct.ID, myStruct.Name)
}

func increment(i int) {
	i++
	fmt.Printf("The value of i is: %d \n", i)
	fmt.Printf("The address of i is: %p \n", &i)
	
}

func incrementP(i *int) {

	*i++
	fmt.Printf("The value of i is: %d \n", *i)
	fmt.Printf("The address of i is: %p \n", i)
	
}

func updateSlice(v []int){
	v[0] = 100
	fmt.Printf("The value of mySlice is: %v and the direction is %p \n ", v, &v)
	
	fmt.Printf("The value of mySlice[0] is: %v and the direction is %p \n ", v[0], &v[0])
	fmt.Printf("The value of mySlice[1] is: %v and the direction is %p \n ", v[1], &v[1])
	fmt.Printf("The value of mySlice[2] is: %v and the direction is %p \n ", v[2], &v[2])
}

type MyStruct struct{
	ID int
	Name string
}

func (m MyStruct) SetName(name string){
	fmt.Printf("The address of m is: %p \n", m)
	m.Name = name
}

func (m *MyStruct) SetP(name string){
	fmt.Printf("The address of m is: %p \n", m)
	m.Name = name
}

func updateStruct(m *MyStruct){
	fmt.Printf("The address of m is: %p \n", m)
	m.ID = 999
	m.Name = "Jane"
}

