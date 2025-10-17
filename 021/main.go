package main
import "fmt";

func main(){
	// anynomous function
	// IIFE

	func(a,b int){
		fmt.Println(a + b)
	}(10,20)
}

func init() {
	fmt.Println("Init function called")
}

// init function call first then call main function...