package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("main.go")

	if err != nil {
		fmt.Println("An error occurred: ", err)
		os.Exit(1)
	}

	fmt.Println(file)
/*
	v, _ := file.Stat()

	fmt.Println(v.Name(), v.Size(), v.Mode(), v.ModTime(), v.IsDir())

	data := make([]byte, 1024)
	//fmt.Println(data)

	c, err := file.Read(data)
	if err != nil {
		fmt.Println("An error occurred: ", err)
		os.Exit(1)
	}
	fmt.Println(string(data[:c]))

	fmt.Printf("read: %d bytes: %q\n", c, data[:c])
*/
	//fmt.Println(os.Getenv("USERNAME"))

	env, ok := os.LookupEnv("USERNAME")
	fmt.Println(env)
	fmt.Println(ok)
}
