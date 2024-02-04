package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Display(myValue int) {
	display(myValue)
}

func display(myValue int) {
	fmt.Println("Hello, World!")
	fmt.Printf("My integer variable is %d , his type is %T, his addres is %v \n\n", myValue, myValue, &myValue)
}

func Add(a, b int) int {
	return a + b
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, errors.New("invalid operation")
	}

}

func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

func MSum(values ...float64) (result float64) {
	for _, value := range values {
		result += value
	}
	return
}

func MOperations(op Operation, values ...float64) (result float64, err error) {
	if len(values) == 0 {
		return 0, errors.New("no values to operate")
	}
	result = values[0]
	for _, value := range values[1:] {
		switch op {
		case SUM:
			result += value
		case SUB:
			result -= value
		case MUL:
			result *= value
		case DIV:
			if value == 0 {
				return 0, errors.New("division by zero")
			}
			result /= value
		default:
			return 0, errors.New("invalid operation")
		}
	}
	return
}
