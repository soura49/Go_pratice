package main

import (
	"fmt"
	"reflect"
)

func main(){
	name := "Bala"
	course := "GO Fundamantels"
	charge := 36
	fmt.Println("Name is",name,"is type of",reflect.TypeOf(name))
	fmt.Println("Course is",course,"is type of",reflect.TypeOf(course))
	fmt.Println("Charge is",charge,"is type of",reflect.TypeOf(charge))
}