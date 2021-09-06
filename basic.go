// use package main for running package
package main

// import package fmt to use Println
import (
	"fmt"
	"log"
	"strconv"
)

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

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Variable")
	fmt.Println("===========================")

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

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Constant")
	fmt.Println("===========================")

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

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Data Conversion")
	fmt.Println("===========================")

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

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Type Declarations")
	fmt.Println("===========================")

	// Type declarations
	// make your own type, alias
	type IdNum int
	var id IdNum = 12345678
	fmt.Println(id)
	fmt.Println(IdNum(87654321))

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Math Operation")
	fmt.Println("===========================")

	// Math operation
	var a = 4
	var b = 2
	var c = a + b
	fmt.Println(c)

	var d = 20 - 20
	d = 21 * 2
	d = 3 / 3
	d = 10 % 2
	fmt.Println(d)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Augmented Assignment")
	fmt.Println("===========================")

	// Augmented assignment
	var e = 21
	e += 2 // e = e + 2, the rest is same
	e -= 2
	e /= 2
	e *= 3
	e %= 2
	fmt.Println(e)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Urnary Operation")
	fmt.Println("===========================")

	// Urnary operation
	var f = 4
	f++ // f = f + 1
	f-- // f = f - 1
	fmt.Println(f)
	var negative = -4
	var positif = +4 // for positif int you dont have to use the "+" symbol
	fmt.Println(negative)
	fmt.Println(positif)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Comparison Operation")
	fmt.Println("===========================")

	// Comparison operation
	// same as python
	// it doesnt have to be int, it can be used for string and else
	var g = 5
	var j = 6
	fmt.Println(g > j)
	fmt.Println(g < j)
	fmt.Println(g >= j)
	fmt.Println(g <= j)
	fmt.Println(g == j)
	fmt.Println(g != j)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Boolean")
	fmt.Println("===========================")

	// Boolean operation
	// && = and
	// || = or
	// ! = opposite
	fieldGoal := 30
	block := 10
	assist := 13

	var mvpFieldGoal bool = fieldGoal > 20
	var mvpBlock bool = block > 7
	var mvpAssist bool = assist > 8

	var mvp bool = mvpFieldGoal && mvpBlock && mvpAssist
	fmt.Println("Are you the mvp :", mvp)

	var mvp2 bool = mvpFieldGoal || mvpBlock || mvpAssist
	fmt.Println("Are you the mvp :", mvp2)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Array")
	fmt.Println("===========================")

	// Array or list in python
	//
	// array used to create a set amount of variable with the same data type
	// the difference in go, array have fixed set of data
	// unlike python where you dont have fixed set of data
	var kota [4]string
	kota[0] = "Jakarta"
	kota[1] = "Bandung"
	kota[2] = "Bogor"
	kota[3] = "Depok"

	fmt.Println(kota)
	fmt.Println(kota[0])
	fmt.Println(kota[3])

	// Declare and set value array at the same time
	var angka = [4]int{
		4, 5, 6, 7,
	}
	fmt.Println(angka)
	// length array, same as lenght string
	fmt.Println("Length array :", len(angka))

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Slice")
	fmt.Println("===========================")

	// Slice
	// almost same like slicing in python
	// slice is like a snapshot or a copy part of the data from array
	// slice have 3 artibutes, pointer, length, and capacity
	// pointer = the first data from the array
	// length = length of the slice
	// capacity = capacity of the slice that you get from the start of the slice(pointer)
	// 			and end at the last of array
	// example = array[10]
	// slice = array[4:9]
	// pointer = array[4]
	// length = array[4] until array[9] so = 5, from 4 until 8, you dont count the last array in slice
	// capacity = 7, from array[4] until array[10]
	// if you change the value of the array, the value of the slice also changes
	// because slice reference from array

	// you can use [...] if dont know the capaticy of your array
	var kota2 = [...]string{
		"Jakarta",
		"Bandung",
		"Bogor",
		"Depok",
		"Tangerang",
		"Padang",
		"Bali",
		"Lombok",
		"Medan",
	}

	var slice = kota2[4:6]
	var slice2 = kota2[2:]
	var slice3 = kota2[:5]
	var slice4 = kota2[:]
	fmt.Println(slice, "\n", slice2, "\n", slice3, "\n", slice4)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Slice function")
	fmt.Println("===========================")

	// Slice function
	// len(slice) = length of the slice
	// cap(slice) = capacity of the slice
	// append(slice, data) = add a new data to the end of slice, if the capacity already max, it will create a new array
	// make([]TypeData, length, capacity) = create a new slice without data
	// copy(destination, source) = copy slice from the source to the destination
	fmt.Println("Length of the slice :", len(slice))
	fmt.Println("Capacity of the slice :", cap(slice))
	var slice5 = kota2[5:]
	fmt.Println(slice5)

	// since we append to the max capacity array, we create a new array
	var slice6 = append(slice5, "Ujung Kulong")
	fmt.Println(slice6)
	// when we edit the new array, it doesnt effect the old array(reference array)
	slice6[0] = "Tokyo"
	fmt.Println(slice6)
	// reference array
	fmt.Println(slice5)
	fmt.Println(kota2)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Make slice")
	fmt.Println("===========================")

	// Make slice function
	newSlice := make([]string, 4, 6)
	newSlice[0] = "abc"
	newSlice[1] = "def"
	newSlice[2] = "ghi"
	newSlice[3] = "jkl"

	fmt.Println(newSlice)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Copy slice function")
	fmt.Println("===========================")

	// Copy slice function
	fromSlice := kota2[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Array vs Slice")
	fmt.Println("===========================")

	// Array vs Slice
	array := [3]int{1, 2, 3}    // This is array
	array2 := [...]int{1, 2, 3} // This is also array
	slice1 := []int{1, 2, 3}    // This is slice

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(slice1)

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Map")
	fmt.Println("===========================")

	// Map
	// is like dictionary in python, or hash
	// map[KeyType]ValueType = map[string]int

	// var person map[string]string = map[string]string{
	// 		"name": "Izzan",
	// 		"country": "Indonesia",
	// }

	person := map[string]string{
		"name":    "Izzan",
		"country": "Indonesia",
	}

	person["age"] = "26"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["country"])
	fmt.Println(person["age"])

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Map fucntion")
	fmt.Println("===========================")

	// Map function
	// len(map), length of the map
	// map[key], get the value of the data using key
	// map[key] = value, change the value of that key or create new key and value
	// make(map[TypeKey]TypeValue) = create new map wihtout data
	// delete(map, key) = delete data with that key value

	var weather map[string]string = make(map[string]string)
	weather["weather"] = "rain"
	weather["temperature"] = "26"
	weather["unit"] = "celcius"
	weather["city"] = "jakarta"

	delete(weather, "city")
	fmt.Println(weather)
	fmt.Println(len(weather))

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("If Else")
	fmt.Println("===========================")

	// If Else expression
	if weather["weather"] == "rain" {
		fmt.Println("raining outside!")
	}

	if weather["unit"] == "farenheit" {
		fmt.Println("You're in USA")
	} else {
		fmt.Println("You're not in USA")
	}

	// turn weather temperature value from string into int using strconv.Atoi and check if there is an error
	// import strconv and log
	temperature, err := strconv.Atoi(weather["temperature"])
	if err != nil {
		log.Fatal(err)
	}

	if temperature > 35 {
		fmt.Println("Hot outside!")
	} else if temperature > 20 && temperature < 35 {
		fmt.Println("Sunny outside!")
	} else {
		fmt.Println("Freezing outside!")
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("If without short statement")
	fmt.Println("===========================")

	// if without short statement
	var length2 = len(weather)
	if length2 < 4 {
		fmt.Println("weather is not complete")
	} else {
		fmt.Println("weather is complete")
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("If with short statement")
	fmt.Println("===========================")

	// if with short statement
	if length := len(weather); length < 4 { // <-- short statement
		fmt.Println("weather is not complete")
	} else {
		fmt.Println("weather is complete")
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Switch case")
	fmt.Println("===========================")

	// Switch case Expression
	role := "dev"

	switch role {
	case "dev":
		fmt.Println("Welcome")
	case "CEO":
		fmt.Println("Welcome chief")
	default:
		fmt.Println("You're not a dev")
		fmt.Println("You're not welcome here")
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Switch case with short statement")
	fmt.Println("===========================")

	// swtich case with short statement
	switch bounty := 80; bounty > 70 { // <-- short statement
	case true:
		fmt.Println("High value target")
	case false:
		fmt.Println("Low value target")
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Switch case without condition")
	fmt.Println("===========================")

	// swtich case without condition in swtich, only condition in case
	target := "important"
	switch {
	case target == "important":
		fmt.Println("Priority this target")
	case target == "not important":
		fmt.Println("Skip this target")
	default:
		fmt.Println("Queue this target")
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("For loops")
	fmt.Println("===========================")

	// For loops
	counter := 1
	for counter <= 4 {
		fmt.Println("Counter at :", counter)
		counter++
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("For with statement")
	fmt.Println("===========================")

	// for with statement
	// statement at the start = init statement
	// statement at the end = post statement
	// [init statement]; [condition]; [post statement]
	for count := 0; count <= 10; count++ {
		fmt.Println("Count ", count)
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("For with slice")
	fmt.Println("===========================")

	// for with slice
	slice7 := []string{"Ahmad", "Izzan", "Zahrial"}

	for i := 0; i < len(slice7); i++ {
		fmt.Println(slice7[i])
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("For range")
	fmt.Println("===========================")

	// For range
	for index, name := range slice7 {
		fmt.Println("Index : ", index, "\n", "name : ", name)
	}

	// you can use "_" if you dont need the value
	// for _, name := range slice7

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("For in map")
	fmt.Println("===========================")

	// For in map
	for key, value := range weather {
		fmt.Println(key, " = ", value)
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Break and Continue")
	fmt.Println("===========================")

	// Break
	for i := 0; i < 99; i++ {
		if i == 4 {
			break // Stop the function(for loop)
		}

		fmt.Println("i count: ", i)
	}

	fmt.Println("========================")

	// Continue
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue // Skip the condition
		}

		fmt.Println("i count: ", i)
	}

	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println("Function")
	fmt.Println("===========================")

	// Function
	testFunc() // <-- calling the function

	// Function with parameter
	sectionDivider("Function Parameter") // <-- calling the funciton and input the parameter
	fmt.Println("Using fuction parameter")

	sectionDivider("Function Return")
	// Function return
	result := getResult(4, 4)
	fmt.Println(result)
}

// Function
func testFunc() {
	fmt.Println("This is a test!")
}

// Function Parameter
// create section divider using function paramter
func sectionDivider(section string) {
	fmt.Println("")
	fmt.Println("===========================")
	fmt.Println(section)
	fmt.Println("===========================")
}

// Function Return
func getResult(value1 int, value2 int) int {
	return value1 + value2
}
