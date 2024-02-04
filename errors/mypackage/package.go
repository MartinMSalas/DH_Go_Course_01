package mypackage

import "fmt"
func Run() {
	defer func() {
		fmt.Println("start of defer run function")
		r:= recover()

		if r != nil {
			fmt.Println("Recovered from: ", r)
		}
	}()
	panic("panic in run function")
	fmt.Println("end of run function")
}