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