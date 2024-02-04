package main

import (
	"time"
	"fmt"
)

func main() {
	p := fmt.Println
	now := time.Now()
	then := time.Date(2021, 10, 1, 20, 34, 58, 651387237, time.UTC)
	diff := now.Sub(then)
	p(now) 

	fmt.Printf("The year is: %d\n", now.Year())
	fmt.Printf("The month is: %d\n", now.Month())
	fmt.Printf("The day is: %d\n", now.Day())
	fmt.Printf("The hour is: %d\n", now.Hour())
	fmt.Printf("The minute is: %d\n", now.Minute())
	fmt.Printf("The second is: %d\n", now.Second())
	fmt.Printf("The nanosecond is: %d\n", now.Nanosecond())
	fmt.Printf("The location is: %s\n", now.Location())
	fmt.Printf("The week day is: %s\n", now.Weekday())
	fmt.Printf("The year day is: %d\n", now.YearDay())
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700 MST"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700 MST Mon"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700 MST Mon Jan"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700 MST Mon Jan 2"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700 MST Mon Jan 2 15"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700 MST Mon Jan 2 15:04"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700 MST Mon Jan 2 15:04:05"))
	fmt.Printf("The time is: %s\n", now.Format("2006-01-02 03:04:05 PM -0700 MST Mon Jan 2 15:04:05 -0700"))
	fmt.Println("The difference is: ", diff.Hours()/24, " days")
}