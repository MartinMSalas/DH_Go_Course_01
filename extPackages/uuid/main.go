package main

import (
	"fmt"


	"github.com/google/uuid"
)

func main() {
	id1 := uuid.New()
	fmt.Println(id1)
	fmt.Println(id1.String())


}