package main	
import "fmt"

// func main(){
// 	fmt.Println("Hello, World on Golang!")
// }

//variables

var nombre string = "Juan"
var edad int = 20
var pi_1 float64 = 3.14

// constants
const pi float64 = 3.14


// functions
func name() string {
	fmt.Println(nombre)
	return nombre
}

func main(){
	name()
}

