// use package main for running package
package main

// import package fmt to use Println
import "fmt"

// main function
func main() {
	fmt.Println("Hello world")

	// use go build [name of the go file] to compile the file into executable
	// example = "go build basic.go"
	// for development you can use go run [name of the go file] to run the file
	// without compile it into executable
	// example = "go run basic.go"

	// Boolean
	fmt.Println("Right = ", true)
	fmt.Println("Wrong = ", false)

	// String
	fmt.Println("") // still a string even tho the value is empty
	fmt.Println("Ahmad Izzan Zahrial")
	// String function
	// length of the string
	fmt.Println(len("Ahmad Izzan Zahrial"))
	// byte value of the character
	fmt.Println("Ahmad Izzan Zahrial"[0]) // byte value of "A"

	// Variable
	var name string

	name = "Ahmad Izzan Zahrial"
	fmt.Println(name)

	// In go you have to use every variable that created

	// You can't use different type of variable, example
	// name = 100

	// if you declare the value of the variable, you dont have to declare the type
	var age = 26
	fmt.Println(age)
	// but if you want you still can
	var hobby string = "thinking"
	fmt.Println(hobby)
	// you can use spesific type by declaring it
	var favNum int8 = 4
	fmt.Println(favNum)

	// You dont have to use var to declare variable
	roomColor := "Black" // You can use :=
	fmt.Println(roomColor)
	// After the declaring it for the first time, you dont have to use the ":"
	// just use the "="
	roomColor = "Blue"
	fmt.Println(roomColor)

	// Declare multiple variable
	var (
		firstLanguage  = "Bahasa Indonesia"
		secondLanguage = "English"
	)
	fmt.Println(firstLanguage)
	fmt.Println(secondLanguage)
}
