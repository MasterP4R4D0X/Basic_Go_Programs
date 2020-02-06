package main 
import "fmt"
func main() {
	x,y := 10,20
	fmt.Println(add(x,y))
}
func add(num1,num2 int) int{
	return num1+num2
}