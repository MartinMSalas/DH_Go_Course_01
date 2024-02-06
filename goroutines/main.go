package main

import (
	"fmt"
	"time"
)

func main(){
	go myProcess("A")
	go myProcess("B")
	myFirstChannel := make(chan string)
	
	go myProccessWithChannel("C", myFirstChannel)
	
	result := <- myFirstChannel
	fmt.Println(result)
	close(myFirstChannel)

	channelA := make(chan string)
	channelB := make(chan string)

	go myProccessWithChannel("D", channelA)
	go myOtherProccessWithChannel("E", channelB)

	resultA := <- channelA
	fmt.Println("A: ", resultA)
	resultB := <- channelB
	fmt.Println("B: ",resultB)

	time.Sleep(time.Second*2)
}
func myProccessWithChannel(p string, c chan string){
	i:=0
	for i<20{
		time.Sleep(time.Millisecond*100)
		i++
		fmt.Println(p, ":", i)
	}

	c <- "ok"
}

func myOtherProccessWithChannel(p string, c chan string){
	i:=0
	for i<50{
		time.Sleep(time.Millisecond*100)
		i++
		fmt.Println(p, ":", i)
	}
	
	c <- "ok"
}
func myProcess(p string){
	i:=0
	for i<15{
		time.Sleep(time.Millisecond*100)
		i++
		fmt.Println(p, ":", i)
	}

}