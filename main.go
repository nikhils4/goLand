package main

import "fmt"
import "math"
import "time"
//import "strings"


func main() {


	fmt.Println("hello world") // normal string
	fmt.Println("go lang" + " is love") // string concatenation
	fmt.Println(!true) // boolean testing

	var a int = 60 // proper declaration
	b := 90.9 // shorthand notation

	fmt.Println(a,b) // printing the value to the console

	const c float64 = 90  // declaration of const entity type

	fmt.Println(math.Sin(c))  // sin value using the math object to access it

	i := 1

	for i <=3 {
		fmt.Println("loop", i)    // by explicitly defining counter
		i = i + 1
	}

	for {
		fmt.Println("loop")    // breaking the loop using break keyword // can also use continue instead (whenever needed)
		break
	}


	for k := 0 ; k < 6 ; k++ {   // iterating over using the in counter variable
		fmt.Println(k)
	}

	if num := 9; num < 10 {
		fmt.Println("The num is smaller than 10")     // start else if from the line where bracket closes
	} else if num > 10 {
		fmt.Println("The num is greater than 10")		// upper similar rules apply for all the lines in the if else or if else if statements
	} else {
		fmt.Println("No guesses about the number")
	}

	i = 2     // switch statements
	switch i {
		case 1 :
			fmt.Println("One")         // cases
		case 2 :
			fmt.Println("Two")
		default :						// default cases
			fmt.Println("This is the default value")
	}

	t := time.Now()     // time package
	fmt.Println(t.Hour())


	switch t.Weekday() {					// switch statement --> with time package
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	fmt.Println(t.Weekday())
	//fmt.Println(t.Weekday() == "Saturday")


	whatAmI := func(i interface{}) {     /// kinda function declaration taking as input the data
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")


	//Array

	var arr [5]int
	fmt.Println(arr)
	arr[3] = 10
	fmt.Println(arr)

	// to know the data type

	x := 29
	fmt.Printf("x is of type %T\n", x) // %T is the key here

	// lvalue is something that refers to a memory location in terms of Go lang
	// rvalue is something that refers to the value being stored there

	// x = 20 ; x --> lvalue and y --> rvalue


	max(20, 25)

	//using function with multiple return values
	//a, b := swap("Mahesh", "Kumar")
	//fmt.Println(a, b)


	// call by value doesn't change the main var value
	// call by reference changes the main var value in the memory location


	//function as values

	/* declare a function variable */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* use the function */
	fmt.Println(getSquareRoot(9))

	// to make var global just put ot outside main function
	// and use it anywhere in the code


	//%s for string
	//%x for hex code of the string for each character

	//len(string)  returns the length of thr string



	// two ways of concatenating string



	//fmt.Println(strings.Join("Hello ", "Nikhil"))

	fmt.Println("Hello " + "world")


	// Arrays

	//var balance [10] float32
	var index =  [2] int {1,2} // can leave the square bracket blank as well when don't know the exact size of the array
	fmt.Println(index[1])

	index[0] = 3.0

	// pointers are var that are used to store any var address
	// to know the address of any var use & operator but keep in mind that this particular address can only be stored in pointers

	var addOfVar int = 22

	var ip *int /* points to an integer float */
	ip  = &addOfVar

	// to get the vakue of ip i.e. what exactly it has stored go with *ip

	fmt.Println(*ip)


	var nikhil student

	nikhil.name = "Nikhil Singh"
	nikhil.registerNumber = "17BIT0192"

	fmt.Println(nikhil.name)
	fmt.Println(&nikhil)

}


func max(num1, num2 int) int {
	fmt.Println(num1, num2)
	return num1
}


// returning multiple values from the function

func swap(x, y string) (string, string) {
	return y, x
}


// anonymous functions :: first is the main function while second is the returned function

//func getSequence() func() int {
//	i:=0
//	return func() int {
//		i+=1
//		return i
//	}
//}
//
//func main(){
//	/* nextNumber is now a function with i as 0 */
//	nextNumber := getSequence()
//
//	/* invoke nextNumber to increase i by 1 and return the same */
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//
//	/* create a new sequence and see the result, i is 0 again*/
//	nextNumber1 := getSequence()
//	fmt.Println(nextNumber1())
//	fmt.Println(nextNumber1())
//}



// function as a method   ----------> important to unsderstand as a method

///* define a circle */
//type Circle struct {
//	x,y,radius float64
//}
//
///* define a method for circle */
//func(circle Circle) area() float64 {
//	return math.Pi * circle.radius * circle.radius
//}
//
//func main(){
//	circle := Circle{x:0, y:0, radius:5}
//	fmt.Printf("Circle area: %f", circle.area())
//}



//passing array as parameter to the function

//func main() {
//	/* an int array with 5 elements */
//	var  balance = []int {1000, 2, 3, 17, 50}
//	var avg float32
//
//	/* pass array as an argument */
//	avg = getAverage( balance, 5 ) ;
//
//	/* output the returned value */
//	fmt.Printf( "Average value is: %f ", avg );
//}
//func getAverage(arr []int, size int) float32 {
//	var i,sum int
//	var avg float32
//
//	for i = 0; i < size;i++ {
//		sum += arr[i]
//	}
//
//	avg = float32(sum / size)
//	return avg;
//}



// multidimensional array
//
//func main() {
//	/* an array with 5 rows and 2 columns*/
//	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
//	var i, j int
//
//	/* output each array element's value */
//	for  i = 0; i < 5; i++ {
//		for j = 0; j < 2; j++ {
//			fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
//		}
//	}
//}



// all about pointers -----> executed inside the main function


//// pointers are var that are used to store any var address
//// to know the address of any var use & operator but keep in mind that this particular address can only be stored in pointers
//
//var addOfVar int = 22
//
//var ip *int /* points to an integer float */
//ip  = &addOfVar
//
//// to get the vakue of ip i.e. what exactly it has stored go with *ip
//
//fmt.Println(ip)


// nil pointers are the pointers those are declared but not are initialised with address of any variable
// such state leads to nil pointer -----> ptr == nil
//
//func main() {
//	a := []int{10,100,200}
//	var i int
//	var ptr [MAX]*int;
//
//	for  i = 0; i < MAX; i++ {
//		ptr[i] = &a[i] /* assign the address of integer. */
//	}
//	for  i = 0; i < MAX; i++ {
//		fmt.Printf("Value of a[%d] = %d\n", i,*ptr[i] )
//	}
//}

//here the ptr is an array of pointers which contain the adsress of all the elements of another array


// variable that is a pointer to a pointer should be declared as such  **addrOfPtr

//func main() {
//	var a int
//	var ptr *int
//	var pptr **int
//
//	a = 3000
//
//	/* take the address of var */
//	ptr = &a
//
//	/* take the address of ptr using address of operator & */
//	pptr = &ptr
//
//	/* take the value using pptr */
//	fmt.Printf("Value of a = %d\n", a )
//	fmt.Printf("Value available at *ptr = %d\n", *ptr )
//	fmt.Printf("Value available at **pptr = %d\n", **pptr)
//}


//passing pointers to the function in GO

//func main() {
//	/* local variable definition */
//	var a int = 100
//	var b int = 200
//
//	fmt.Printf("Before swap, value of a : %d\n", a )
//	fmt.Printf("Before swap, value of b : %d\n", b )
//
//	/* calling a function to swap the values.
//	 * &a indicates pointer to a ie. address of variable a and
//	 * &b indicates pointer to b ie. address of variable b.
//	 */
//	swap(&a, &b);
//
//	fmt.Printf("After swap, value of a : %d\n", a )
//	fmt.Printf("After swap, value of b : %d\n", b )
//}
//func swap(x *int, y *int) {
//	var temp int
//	temp = *x    /* save the value at address x */
//	*x = *y      /* put y into x */
//	*y = temp    /* put temp into y */
//}


// in arrays in go you cannot hold data items of different data type for that we need struct


// struct

//type struct_variable_type struct {
//	member definition;
//	member definition;
//	...
//	member definition;
//}

type student struct {
	name string
	registerNumber string
}
//
//func main () {
//	var nikhil student
//
//	nikhil.name = "Nikhil Singh"
//	nikhil.registerNumber = "17BIT0192"
//
//	fmt.Println(nikhil.name)
//	fmt.Println(&nikhil)
//}

// struct as function arguments

//func someFunc ( hi student) {
//	// some code here
//
//	// here the struct would be reffered as hi
//}

//var struct_pointer *Books
//struct_pointer = &Book1
//struct_pointer.title


//func some (hello *student) {
// hello.name to access inside the function
// using pointer to structure inside the code
//}

// len() and cap() of a slice ----> length and capacity of a slice
//
//func main() {
//	/* create a slice */
//	numbers := []int{0,1,2,3,4,5,6,7,8}
//	printSlice(numbers)
//
//	/* print the original slice */
//	fmt.Println("numbers ==", numbers)
//
//	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
//	fmt.Println("numbers[1:4] ==", numbers[1:4])
//
//	/* missing lower bound implies 0*/
//	fmt.Println("numbers[:3] ==", numbers[:3])
//
//	/* missing upper bound implies len(s)*/
//	fmt.Println("numbers[4:] ==", numbers[4:])
//
//	numbers1 := make([]int,0,5)
//	printSlice(numbers1)
//
//	/* print the sub slice starting from index 0(included) to index 2(excluded) */
//	number2 := numbers[:2]
//	printSlice(number2)
//
//	/* print the sub slice starting from index 2(included) to index 5(excluded) */
//	number3 := numbers[2:5]
//	printSlice(number3)
//
//}
//func printSlice(x []int){
//	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x),x)

//for integer place holder %d and for array we use %v

// cannot change the size of array so slice is used most frequently using which we can make array of variable lenght


//func main() {
//	var numbers []int
//	printSlice(numbers)
//
//	/* append allows nil slice */
//	numbers = append(numbers, 0)
//	printSlice(numbers)
//
//	/* add one element to slice*/
//	numbers = append(numbers, 1)
//	printSlice(numbers)
//
//	/* add more than one element at a time*/
//	numbers = append(numbers, 2,3,4)
//	printSlice(numbers)
//
//	/* create a slice numbers1 with double the capacity of earlier slice*/
//	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
//
//	/* copy content of numbers to numbers1 */
//	copy(numbers1,numbers)
//	printSlice(numbers1)
//}
//func printSlice(x []int){
//	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}

// range and key-value pair all declartion and iteration

//func main() {
//	/* create a slice */
//	numbers := []int{0,1,2,3,4,5,6,7,8}
//
//	/* print the numbers */
//	for i:= range numbers {
//		fmt.Println("Slice item",i,"is",numbers[i])
//	}
//
//	/* create a map*/
//	countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}
//
//	/* print map using keys*/
//	for country := range countryCapitalMap {
//		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
//	}
//
//	/* print map using key-value*/
//	for country,capital := range countryCapitalMap {
//		fmt.Println("Capital of",country,"is",capital)
//	}
//}

// slicing to array
// mapping to key-value pair

// deleting a key-value pair
//    /* delete an entry */
//   delete(countryCapitalMap,"France");
//   fmt.Println("Entry for France is deleted")

// even while deleting a key-value pair use key to identify a particular key-value pair
// and get that deleted using the above syntax


// recursion

//func fibonaci(i int) (ret int) {
//	if i == 0 {
//		return 0
//	}
//	if i == 1 {
//		return 1
//	}
//	return fibonaci(i-1) + fibonaci(i-2)
//}
//func main() {
//	var i int
//	for i = 0; i < 10; i++ {
//		fmt.Printf("%d ", fibonaci(i))
//	}
//}

// type casting ------> int(someVariable)



//interfaces represent set of method signatures

//The struct data type implements these interfaces to have method definitions for the method signature of the interfaces.

///* define an interface */
//type Shape interface {
//	area() float64
//}
//
///* define a circle */
//type Circle struct {
//	x,y,radius float64
//}
//
///* define a rectangle */
//type Rectangle struct {
//	width, height float64
//}
//
///* define a method for circle (implementation of Shape.area())*/
//func(circle Circle) area() float64 {
//	return math.Pi * circle.radius * circle.radius
//}
//
///* define a method for rectangle (implementation of Shape.area())*/
//func(rect Rectangle) area() float64 {
//	return rect.width * rect.height
//}
//
///* define a method for shape */
//func getArea(shape Shape) float64 {
//	return shape.area()
//}
//
//func main() {
//	circle := Circle{x:0,y:0,radius:5}
//	rectangle := Rectangle {width:10, height:5}
//
//	fmt.Printf("Circle area: %f\n",getArea(circle))
//	fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
//}


// error handling

//func Sqrt(value float64)(float64, error) {
//	if(value < 0){
//		return 0, errors.New("Math: negative number passed to Sqrt")
//	}
//	return math.Sqrt(value), nil
//}
//func main() {
//	result, err:= Sqrt(-1)
//
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(result)
//	}
//
//	result, err = Sqrt(9)
//
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(result)
//	}
//}