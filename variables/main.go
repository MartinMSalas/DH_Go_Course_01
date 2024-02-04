package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	/*
		var myIntVar int
		//myIntVar = -42
		fmt.Println("My variable is: ", myIntVar)
		var myUintVar uint
		myUintVar = 42
		fmt.Println("My Unsigned int variable is: ", myUintVar)

		var myStringVar string

		fmt.Println("My string variable is: ", myStringVar)
		myStringVar = "Hello, World!"
		fmt.Println("My string variable is: ", myStringVar)

		var myBoolVar bool
		fmt.Println("My boolean variable is: ", myBoolVar)
		myBoolVar = true
		fmt.Println("My boolean variable is: ", myBoolVar)

		myIntVar2 := 42
		fmt.Println("My variable is: ", myIntVar2)

		myStringVar3 := "test1234"
		fmt.Printf("mi valor es %s \n", myStringVar3)

		floatVar := 3.14
		fmt.Printf("my value is %f \n and my type is %T \n", floatVar, floatVar)

		floatStringVar := fmt.Sprintf("%f", floatVar)
		fmt.Printf("my value is %s \n and my type is %T \n and my size is %d \n", floatStringVar, floatStringVar, unsafe.Sizeof(floatStringVar))

		intVar2 := 342
		intStrVart2 := strconv.Itoa(intVar2)
		fmt.Printf("my value is %s \n and my type is %T \n and my size is %d \n", intStrVart2, intStrVart2, unsafe.Sizeof(intStrVart2))
	*/
	strIntVar := "42"
	intVar, err := strconv.Atoi(strIntVar)
	fmt.Println("Error: ", err)

	fmt.Printf("my value is %d \n and my type is %T \n and my size is %d \n", intVar, intVar, unsafe.Sizeof(intVar))

	strIntVar3, err := strconv.ParseInt(strIntVar, 16, 64)
	fmt.Println("Error: ", err)
	fmt.Printf("my value is %d \n and my type is %T \n and my size is %d \n", strIntVar3, strIntVar3, unsafe.Sizeof(strIntVar3))

	number := 1

	switch number {
	case 1 + intVar:
		fmt.Println("My number is 1")
		fmt.Println("My number is 1")
	case 2:
		fmt.Println("My number is 2")
	case 3:
		fmt.Println("My number is 3")
	default:
		fmt.Println("My number is not 1, 2 or 3")

	}

	for i := 0; i < 10; i++ {
		fmt.Println("My number is: ", i)
	}
	sum := 1

	fmt.Println("My sum is: ", sum)

	s := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(s)
	myArrayVar1 := [10]int{}
	fmt.Println("My array is: ", myArrayVar1)
	for i := 0; i < 10; i++ {
		myArrayVar1[i] = rand.Intn(100)
	}
	for i, v := range myArrayVar1 {
		fmt.Println("My array is: ", i, v)
	}

	myArrayVar2 := [10]int{}
	for i := 0; i < 10; i++ {
		myArrayVar2[i] = rand.Intn(100)
	}
	fmt.Println("My array is: ", myArrayVar2, " and my size is: ", len(myArrayVar2))

	mySliceVar1 := []int{}

	for i := 0; i < 10; i++ {
		mySliceVar1 = append(mySliceVar1, rand.Intn(100))
	}
	fmt.Println("My slice is: ", mySliceVar1, " and my size is: ", len(mySliceVar1))
	mySliceVar1 = append(mySliceVar1, 12, 34, 54, 31, 12)
	fmt.Println("My slice is: ", mySliceVar1, " and my size is: ", len(mySliceVar1))

	mySliceVar2 := mySliceVar1[10:]
	fmt.Println("My slice is: ", mySliceVar2, " and my size is: ", len(mySliceVar2))

	mySliceVar3 := make([]int, 10)

	fmt.Println("My slice is: ", mySliceVar3, " and my size is: ", len(mySliceVar3))

	map1 := make(map[int]string)

	fmt.Println("My map is: ", map1, " and my size is: ", len(map1))
	map1[1] = "one"
	map1[2] = "two"
	map1[3] = "three"
	map1[4] = "four"
	map1[5] = "five"
	fmt.Println("My map is: ", map1, " and my size is: ", len(map1))

	map2 := make(map[int][]string)
	map2[1] = []string{"one", "uno"}
	map2[2] = []string{"two", "dos"}
	fmt.Println("My map is: ", map2, " and my size is: ", len(map2))
	fmt.Println("Map2[1]: ", map2[1])

	map3 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	fmt.Println("My map is: ", map3, " and my size is: ", len(map3))

	for k, v := range map3 {
		fmt.Println("My map is: ", k, v)
	}
	delete(map3, 1)
	delete(map3, 2)
	delete(map3, 3)
	fmt.Println("My map is: ", map3, " and my size is: ", len(map3))
	myFunction(5)
	myIntVar := 5
	fmt.Printf("My int is: %d and my address %v\n", myIntVar, &myIntVar)
}

func myFunction(myValue int64) string {
	fmt.Println("Hello, World!")
	return "Hello, World!"
}
