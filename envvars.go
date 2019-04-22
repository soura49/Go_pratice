package main
import (
	"fmt"
	"os"
)

func main(){
	name := os.Getenv("USER")
	for _, env := range os.Environ(){
		fmt.Println(env)
	}
	fmt.Println(name)
}