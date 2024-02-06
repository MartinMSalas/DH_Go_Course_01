package main

import (
	
	"fmt"
	"log"
	"time"
	"os"
)

func main() {

	date := time.Now()
	file , err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))

	if err != nil {
		log.Fatal("An error occurred: ", err)
	}
	l := log.New(file, "LOG: ", log.LstdFlags|log.Lshortfile)

	l.Println("This is a log message: test 1")
	l.Println("This is a log message: test 2")
	l.Println("This is a log message: test 3")

	l.SetPrefix("errors: ")
	l.Println("This is a log message: test 4")
	l.Println("This is a log message: test 5")
	l.Println("This is a log message: test 6")

	stdoutLog := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	stdoutLog.Println("This is a log message: test 1")
	stdoutLog.Println("This is a log message: test 2")
	stdoutLog.Println("This is a log message: test 3")
	stdoutLog.SetPrefix("errors: ")
	stdoutLog.Println("This is a log message: test 4")
	stdoutLog.Println("This is a log message: test 5")
	stdoutLog.Println("This is a log message: test 6")
	
	
}