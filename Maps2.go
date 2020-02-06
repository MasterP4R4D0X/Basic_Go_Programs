package main 
import "fmt"
func main() {
	superhero := map[string]map[string] string{
		"Ironman" : map[string]string{
		"RealName" : "Tony Stark",
		"City" : "New York",
		},
		"Batman" : map[string]string{
		"RealName" : "Bruce Wayne",
		"City" : "Gotham City",
		},
	}
	if temp, hero := superhero["Ironman"]; hero{
		fmt.Println(temp["RealName"], temp["City"])
	}
}