package main
import(
	"fmt"

)

func main(){
	var B int = 35
	A := &B
	fmt.Println(A) 
	fmt.Println(*A)
	
}