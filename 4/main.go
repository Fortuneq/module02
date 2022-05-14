package main
import (
	"fmt"
	"math"
)

func main(){
	var R float64 = 35 * math.Pi*2
	pntr := &R
	Ploshad := *pntr*(math.Round(math.Pow(math.Pi,2)))
	fmt.Println(Ploshad)
}