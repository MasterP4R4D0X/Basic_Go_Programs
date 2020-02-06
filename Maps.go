package main
import "fmt"
func main() {
	var myMap map[string]int
	myMap=map[string]int{}
	myMap["a"]=12
	myMap["b"]=120
	myMap["c"]=1200
	myMap["d"]=12000
	myMap["e"]=120000
	delete(myMap,"d")
	fmt.Println(myMap)
}