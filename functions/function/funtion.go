package function
import (
	"fmt"
)
func Display(myValue int) {
	fmt.Println("Hello, World!")
	fmt.Printf("My integer variable is %d , his type is %T, his addres is %v \n\n", myValue, myValue, &myValue)
}