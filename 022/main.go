package main
import "fmt"

func add(a,b int){
	fmt.Println("hay: i am global",a + b)
}

func main(){
	add(5,20)
	// Function Expression or assign function in variable
	add := func(a,b int)  {
		fmt.Println("hay i am in main",a + b)
	}
	add(5,10)
}

func init(){
	fmt.Println("Init function called")
}
