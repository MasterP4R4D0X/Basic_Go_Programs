package main 
import "fmt"
func main() {
	var g string="A"
	switch{
	case g=="A":
		fmt.Println("Good grade")
	
	case g=="B":
		fmt.Println("Ok grade")
	
	case g=="C":
		fmt.Println("Not so good grade")
	
	case g=="D":
		fmt.Println("Bad grade")
	
    default:
	fmt.Println("Default")
	}
}