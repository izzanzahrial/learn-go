// use package main for running package
package main

// import package fmt to use Println
import (
	"errors"
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

	sectionDivider("Multiple values return")
	// You can return more than one values in go, like python
	one, two := moreThanOne()
	fmt.Println(one, two)
	// You can use "_" if you dont need the return value
	three, _ := moreThanOne()
	fmt.Println(three)

	sectionDivider("Named return values")
	// Return values with assign key(named)
	firstName, middleName, lastName := fullName()
	fmt.Println(firstName, middleName, lastName)

	sectionDivider("Variadic Function(varargs)")
	// Variadic function(varargs) is like an array but you dont have to create it before hand
	// and variadic function located at the end of the parameter
	total := sumNum(1, 2, 3, 4, 5)
	fmt.Println(total)
	// you can input slice into variadic function, just add "..." at the of the slice that you assigned
	slice8 := []int{10, 9, 8, 7, 6}
	total = sumNum(slice8...) // <-- slice into variadic
	fmt.Println(total)

	sectionDivider("Function as Value")
	// Function as value
	// add a function into a variable
	hello := hello
	fmt.Println(hello("izzan"))

	sectionDivider("Function as Parameter")
	// make a function into a parameter
	sayHelloToMyLittleFriend("machine gun", littleFriendFilter)

	sectionDivider("Function type declarations")
	// using type declarations for function
	sayHelloToMyLittleFriend2("shoot gun", littleFriendFilter)

	sectionDivider("Anonymous Function")
	// like lambda in python, function without name
	// version 1
	inTheList := func(name string) bool { // <-- create variable and add function into that variable
		return name == "Izzan" // <-- check if the name equal "izzan"
	}
	whiteList("Izzan", inTheList) // <-- use variable that contain function, to check the input

	// version 2
	whiteList("Zahrial", func(name string) bool { // <-- without creating a variable, passing the anonymous function
		return name == "Izzan" // into the main function
	})

	sectionDivider("Recursive")
	// Recursive function
	fac := factorial(5)
	fmt.Println(fac)

	sectionDivider("Closure Function")
	// access data outside the scope, or some thing like that
	open := open()
	fmt.Println(open)

	sectionDivider("Defer function")
	// Function that will run after his ancestor/caller
	// defer function will run even though there is an error in the ancestor level if you place it at the start
	// of function
	runApp()

	sectionDivider("Panic function")
	// Function that will stop the application if there is an error in it
	// but if there is an defer in it, the defer function will still run
	start(false) // <-- true = call panic error, false = didn't call panic error

	sectionDivider("Recover function")
	// Function that will catch panic data.
	// with recover function, panic function will stop and the application will keep running
	start2(true)

	sectionDivider("Comments")
	// use "//" for single line comment
	// or use "/*" at the start of the line and "*/" at the of the line for multiple line comment

	sectionDivider("Struct")
	// Is like class in OOP
	var izzan Doctor
	izzan.Name = "Ahmad Izzan Zahrial"
	izzan.Address = "Indonesia"
	izzan.Age = 26
	izzan.Speciality = "Brain surgeon"

	fmt.Println(izzan)

	// another way to creat struct or struct literals
	ahmad := Doctor{
		Name:       "Ahmad",
		Address:    "Indonesia",
		Age:        25,
		Speciality: "Nerve",
	}
	fmt.Println(ahmad)

	zahrial := Doctor{"Zahrial", "Indonesia", "Eyes", 24}
	fmt.Println(zahrial)

	sectionDivider("Struct Method")
	// Is like class function in OOP
	// Is like struct function but not really
	rusli := Doctor{Name: "Rusli"}
	rusli.sayHi()

	sectionDivider("Interface")
	// Is abstract data = you can't implement it directly
	// Is a contract for a function
	// let say you have 2 object, baseball and football
	// that objects can have a similar function = throwAble() or kickAble()
	// based on that you can create contract using interface
	// in that interface you can store the throwAble() as a contract
	// where if you wanna use the interface your object have to have thorAble() function
	// https://www.youtube.com/watch?v=qJKQZKGZgf0&ab_channel=LearnGoProgramming
	var syakila Person
	syakila.Name = "Syakila"

	byeBye(syakila)

	sectionDivider("Interface 2")
	// antother example for interface
	wolf := Animal{
		Name: "Black",
	}
	byeBye(wolf)

	sectionDivider("Empty Interface")
	// is like data but you dont have to specify the type or something like that
	// so basically you can put any value inside it
	var data interface{} = Blank(2)
	fmt.Println(data)

	sectionDivider("Nil or Null or None")
	// In golang every data type has it's own default null value like 0 in int
	// so if you wanna specify it into null or nil, you gotta use "nil" as value
	var person1 map[string]string = nil
	fmt.Println(person1)
	person3 := newMap("Izzan")
	fmt.Println(person3)
	// you can use nil to check if the data is empty

	sectionDivider("Interface Error")
	// golang have build in library for error
	var errorExample error = errors.New("Example Error")
	fmt.Println(errorExample.Error())
	// another example
	value3, err := diveder1(25, 0)
	if err == nil {
		fmt.Println("Result :", value3)
	} else {
		fmt.Println("Error", err.Error())
	}

	sectionDivider("Type Assertions")
	// Type assertions is setting a type for empty interface
	// if you set wrong type, it will result panic(stop the program)
	value4 := testTypeAssertions()
	value4Int := value4.(int)
	fmt.Println(value4Int, value4)
	// if i use type assertions string, it will result panic
	// value4 := testTypeAssertions()
	// value4Int := value4.(string) <-- Panic
	// fmt.Println(value4Int, value4)

	sectionDivider("Type Assertions using Switch")
	value5 := testTypeAssertions()
	switch value6 := value5.(type) {
	case string:
		fmt.Println("String type", value6)
	case int:
		fmt.Println("Int type", value6)
	case bool:
		fmt.Println("Bool type", value6)
	default:
		fmt.Println("Unknown")
	}

	sectionDivider("Pointer")
	// Pointer in go is used for pass data by reference
	// basically like using the data without duplicating it
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1 // "&" symbol is the pointer, used to reference the data in address1
	address3 := &address1
	address4 := &address1

	// var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// var address2 *Address= &address1

	address2.Province = "Banten" // Because we using pointer, the data in address2 and address1 will change
	// it happens because the pointer, point to the same data memory
	// only can be used to change one field value
	fmt.Println(address1)
	fmt.Println(address2)
	line()

	address2 = &Address{"Lombok", "NTB", "Indonesia"} // You can't use "&" to change all the field value
	// instead changing all the field value, it will create a new data

	fmt.Println(address1)
	fmt.Println(address2)
	line()

	// You can use "*" to reference the all data that connect to the memory will change to the new data
	*address3 = Address{"Lombok", "NTB", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address3)
	fmt.Println(address4)

	sectionDivider("Pointer New")
	// You can create pointer using "new", but new can only return empty data, so basically without inital data
	address5 := new(Address)
	fmt.Println(address5)

	sectionDivider("Pointer in Function")
	address6 := Address{
		City:     "Yogyakarta",
		Province: "Daerah Istimewa Yogyakarta",
		Country:  "Indonesia",
	}
	changeCity(&address6, "Solo") // change the address6 type into a pointer using "&"
	fmt.Println(address6)

	sectionDivider("Pointer in struct or method")
	// pointer in struct will be usefull because you wont duplicate your data
	// and it wont overload your memory
	address7 := Address{
		City:     "London",
		Province: "Greater Londong",
		Country:  "United Kingdom",
	}
	address7.goToFrance()
	fmt.Println(address7)
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

func line() {
	fmt.Println("===========================")
}

// Function Return
func getResult(value1 int, value2 int) int {
	return value1 + value2
}

// Return multiple values
func moreThanOne() (string, string) {
	return "One", "Two"
}

// Named return values
func fullName() (firstName, middleName, lastName string) {
	firstName = "Ahmad"
	middleName = "Izzan"
	lastName = "Zahrial"

	// you can return it without assign the return key
	// you can use this if you already assign the named or the key of return data
	return
	// or the normal way
	// return firstName, middleName, lastName
}

// Variadic function(varargs)
func sumNum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function as value
func hello(name string) string {
	return "Hello " + name
}

// Function as parameter
func sayHelloToMyLittleFriend(name string, filter func(string) string) {
	filteredFriend := filter(name)
	fmt.Println("Say hello to my", filteredFriend)
}

func littleFriendFilter(name string) string {
	if name == "machine gun" {
		return "little friend"
	} else {
		return ", never mind"
	}
}

// Function type declarations

// Create the type
type Filter func(string) string

// Assign the type declaration
func sayHelloToMyLittleFriend2(name string, filter Filter) { // use the type declaration
	filteredFriend := filter(name)
	fmt.Println("Say hello to my", filteredFriend)
}

// Anonymous function
// create type function for ease of use
type InTheList func(string) bool

func whiteList(name string, inTheList InTheList) {
	if inTheList(name) {
		fmt.Println("This way please Mr.", name)
	} else {
		fmt.Println("Sorry you're not on the list")
	}
}

// Recursive function
func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}

// Closure function
func open() int {
	value := 0

	open := func() { // <-- closure func
		value++ // accessing data froum outside function
	} // the outside function can't access the data from inside the function

	open()

	return value
}

// Defer function
func logging() {
	fmt.Println("Application Done!")
}

func runApp() {
	defer logging()
	fmt.Println("Run Application")
}

// Panic function
func done() {
	fmt.Println("Done running")
}

func start(error bool) {
	defer done()
	if error {
		panic("Didn't start error!")
	}
	fmt.Println("Start!")
}

// Recover function
func done2() {
	message := recover()
	if message != nil {
		fmt.Println("ERROR :", message)
	}
	fmt.Println("Done running")
}

func start2(error bool) {
	defer done2()
	if error {
		panic("Didn't start error!")
	}
	fmt.Println("Start!")
}

// Struct
// in struct you use Upper case at first for the struct and the field
type Doctor struct {
	Name, Address, Speciality string
	Age                       int
}

// Struct function or method
func (doctor Doctor) sayHi() {
	fmt.Println("Hi, my name is Dr.", doctor.Name)
}

// interface
type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func byeBye(hasName HasName) {
	fmt.Println("Bye bye", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

// interface 2

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// Empty interface
func Blank(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return ""
	}
}

// ^ empty interface above is like interface below
// type Blank interface {

// }

// Nil
func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// Error interface
func diveder1(value int, diveder int) (int, error) {
	if diveder == 0 {
		return 0, errors.New("Can't use 0 as divider")
	} else {
		result := value / diveder
		return result, nil
	}
}

func testTypeAssertions() interface{} {
	return 0
}

// Pointer
type Address struct {
	City, Province, Country string
}

// Pointer in function
func changeCity(address *Address, value string) { // Because we use pointer in this function
	address.City = value // The value of the original data will change
}

// Pointer in struct
func (address *Address) goToFrance() {
	address.City = "Paris"
	address.Province = "Île-de-France"
	address.Country = "France"
}
