package main

import (
	"DH_Go_Course_01/functions/function"
	"fmt"
	"encoding/json"
)
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name,omitempty"`
    Address  string `json:"address,omitempty"`
    Age      uint8  `json:"age,omitempty"`
    LastName string `json:"last_name"`
}

func (u User) Display() {
	v,err := json.Marshal(u)
	fmt.Println("The result of the json marshal error is: ", err)
	fmt.Println("The result of the json marshal value is: ", string(v))
}

func (u *User) SetName(name string) {
	u.Name = name
}
func main() {
	//myIntVar := 10

	
	//function.Display(myIntVar)
	//function.Display(function.Add(5, 6))
	//function.RepeatString(5, "Hello, World!")
	result, err := function.Calc(function.DIV, 50, 10)
	if err != nil {
		fmt.Println("An error occurred: ",err)
		
		
		return
	}
	fmt.Println("The result is: ", result)

	x,y := function.Split(10)
	fmt.Println("The result of the split is: ", x, y)

	fmt.Println("The result of the sum is: ", function.MSum(1,2,3,4,5))

	result,err = function.MOperations(function.SUM, 1,2,3,4,5)
	if err != nil {
		fmt.Println("An error occurred: ", err)
		return
	}
	fmt.Println("The result of the operations is: ", result)

	factOpFunc := function.FactoryOperation(function.SUM)
	fmt.Println("The result of the factory operation is: ", factOpFunc(5, 6))

	user := User{ID: 1, Name: "John", Address: "123 Main St", Age: 30, LastName: "Doe"}
	
	user.Display()
	user.SetName("Jane")
	user.Display()
}


