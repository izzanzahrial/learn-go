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
	father, mother := "rusli", "erna"
	fmt.Println(firstLanguage)
	fmt.Println(secondLanguage)
	fmt.Println(father)
	fmt.Println(mother)

	// Constant = variable with fixed value, and have to initialize the value
	const country string = "Indonesia"
	const city = "Jakarta"
	const rt = 01
	// If you build constant you dont have to use it, unlike variable
	fmt.Println(country)
	fmt.Println(city)
	// in constant you can't use ":="
	// example = "const city := "Bandung" => ERROR
	// Declare multiple const = multiple var
	const (
		favMovie   string = "Toy story"
		favMusic          = "Repeat until death"
		numOfMusic        = 2000
	)
	fmt.Println(favMovie)
	fmt.Println(favMusic)
	fmt.Println(numOfMusic)

	// Data conversion
	// almost like python, but becareful of integer overflow
	var numOfDays8 int8 = 7
	var numOfDays16 int16 = int16(numOfDays8)
	var numOfDays32 int32 = int32(numOfDays16)
	fmt.Println(numOfDays8)
	fmt.Println(numOfDays16)
	fmt.Println(numOfDays32)
	// byte into string
	var object = "Laptop"
	var byteObject = object[0]
	var stringObject = string(byteObject)
	fmt.Println(object)
	fmt.Println(byteObject)
	fmt.Println(stringObject)
}
