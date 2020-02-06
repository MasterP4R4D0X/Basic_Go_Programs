package main 
import "fmt"
func main() {
	arr:=[5]int{1,2,3,4,5}
	slc:=arr[0:3]
	fmt.Println(slc,len(slc))
}