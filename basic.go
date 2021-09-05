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

	// Type declarations
	// make your own type, alias
	type IdNum int
	var id IdNum = 12345678
	fmt.Println(id)
	fmt.Println(IdNum(87654321))

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

	// Augmented assignment
	var e = 21
	e += 2 // e = e + 2, the rest is same
	e -= 2
	e /= 2
	e *= 3
	e %= 2
	fmt.Println(e)

	// Urnary operation
	var f = 4
	f++ // f = f + 1
	f-- // f = f - 1
	fmt.Println(f)
	var negative = -4
	var positif = +4 // for positif int you dont have to use the "+" symbol
	fmt.Println(negative)
	fmt.Println(positif)

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

	// Slice function
	// len(slice) = length of the slice
	// cap(slice) = capacity of the slice
	// append(slice, data) = add a new data to the end of slice, if the capacity already max, it will create a new array
	// make([]TypeData, length, capacity) = create a new slice
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

	// Make slice function
	newSlice := make([]string, 4, 6)
	newSlice[0] = "abc"
	newSlice[1] = "def"
	newSlice[2] = "ghi"
	newSlice[3] = "jkl"

	fmt.Println(newSlice)

	// Copy slice function
	fromSlice := kota2[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	// Array vs Slice
	array := [3]int{1, 2, 3}    // This is array
	array2 := [...]int{1, 2, 3} // This is also array
	slice1 := []int{1, 2, 3}    // This is slice

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(slice1)
}
