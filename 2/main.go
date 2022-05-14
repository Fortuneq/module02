package main
import(
	"fmt"
	"math"	
)

func main(){
	var AmericanVelocity = 120.4 * 3.6
	var EuropeanVelocity = (130 *3.6)/1.6

	fmt.Println(math.Round((AmericanVelocity*100)/100))
	fmt.Println(math.Round((EuropeanVelocity*100)/100))

}
