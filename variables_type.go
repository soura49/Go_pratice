package main

import (
	"fmt"
	"reflect"
)

var (
	name2,course2,module2  = "Nigel","DockerDive",3.2
)

func main() {
	fmt.Println("Name is",name2,"and is of type",reflect.TypeOf(name2))
	fmt.Println("Module is",module2,"and is of type",reflect.TypeOf(module2))
}