package main 
import "fmt"
func main() {
	var b int=15
	var c int=0
	for i := 0; i <10; i++ {
		fmt.Printf("value of i is %d\n",i)
	}
	for c<b{
		c++
		fmt.Printf("Value of c is %d\n",c)
	}
}