// packages
package main

//

// libraries needed
import "fmt"

// main function
func main() {

	// variables
	var message string = "Hello World"
	easyMessage := "Hello World :="

	fmt.Println(message)
	fmt.Println(easyMessage)

	//fmt.Println("Hello World from Go")
	// float numbers
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)

	// integer numbers
	var c int = 10
	d := 3
	fmt.Println(c / d)

	// boolean
	var x bool = true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)

	// private logic
	// use lowercase to assig a private package
	private()

	// Publiclogic
	Public()

	// Defer function
	printHelloWorld()
}

// func types
func private() {
	fmt.Println("Execute logic that doenst need to be exported (belong to this package only)")
}

func Public() {
	fmt.Println("Logic which I want to export to others packages")
}

func printHelloWorld() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	fmt.Println("Defer execute the code till last")
}
