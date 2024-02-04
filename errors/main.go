package main

import (
	"DH_Go_Course_01/errors/mypackage"
	"errors"
	"fmt"
	

	
)

type MyCustomError struct {
	ID string
}
func (m MyCustomError) Error() string {
	return fmt.Sprintf("The error with id %s ", m.ID)
}

func main() {
	err := errors.New("an error occurred")

	fmt.Printf("The error is: %v, his type is %T\n", err,err)
	fmt.Printf("The error converted in string is: %s his type is %T\n", err.Error(),err.Error())

	error3 := TestError(1)
	fmt.Printf("The error3 is: %v, his type is %T\n", error3,error3)

	error4 := TestError(2)
	fmt.Printf("The error4 is: %v, his type is %T\n", error4,error4)

	//Validate is an error is from a struct
	error5 := TestError(1)  
	fmt.Println(errors.As(error5, &MyCustomError{}))

	//errors join
	error6 := errors.Join(error3,error5)
	fmt.Printf("The error6 is: %v, his type is %T\n", error6,error6)
	fmt.Println()

	//errors is
	fmt.Println(errors.Is(error6, error3))
	fmt.Println(errors.Is(error6, error4))
	fmt.Println(errors.Is(error6, error5))

	error7 := errors.New("an error occurred")
	error8 := fmt.Errorf("an error occurred: %w", error7)
	error9 := fmt.Errorf("an error occurred: %w", error8)
	fmt.Println()
	fmt.Println(error9)
	fmt.Println(errors.Unwrap(error9))

	//errors unwrap
	fmt.Println(errors.Unwrap(error6))

	defer func() {
		fmt.Println("start main defer")
		r:= recover()
		if r != nil {
			fmt.Println("Wolo Recovered from wolo: ", r)
		}
	}()
	
	mypackage.Run()

	fmt.Println("end main function")
}

func TestError(v int) error {
	if v ==1 {
		return MyCustomError{"4"}
	}else{
		return errors.New("an error occurred")
	}
	
}