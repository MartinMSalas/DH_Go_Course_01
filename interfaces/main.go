package main

import (
	"fmt"
	"DH_Go_Course_01/interfaces/vehicles"
)

func main(){
	vSlice := []string{"CAR", "MOTORCYCLE", "TRUCK","CAR","CAR","FOOT","CAR", "PLANE"}

	totalDistance := 0.0
	var vDistance = float64(0)
	for _, v := range vSlice {
		fmt.Println("The vehicle is: ", v)
		vehicle, err := vehicles.New(v, 600)
		if err != nil {
			fmt.Println("An error occurred: ", err)
			fmt.Println()
			continue
		}
		vDistance = vehicle.Distance()
		fmt.Printf("The distance is: %.2f\n\n", vDistance)
		totalDistance += vDistance

	}
	fmt.Printf("The total distance is: %.2f\n", totalDistance)

}

func Display(value interface{}){

	fmt.Println(value)
}