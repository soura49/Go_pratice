package main

import (
	"fmt"
)

func main() {
	name := "bala"
	course := "DOCKER DEEP DIVE"
	fmt.Println("\nHi",name,"you're currently watching",course)
	changeCourse(course)
}
func changeCourse(course string) string{
	course = "First name is bla"
	fmt.Println("Trying to change it",course)
	return course
}