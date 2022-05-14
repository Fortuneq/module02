package main

import ("fmt"
		"log"
		"strconv"
)
func main(){
	var (
	price string
	qty int
	)
price = "104"

qty = 35

bigprice, err := strconv.Atoi(price)
if err != nil {
	log.Fatal(err)
}
bigqty := string(qty)
fmt.Printf("price type %T\n", bigprice)
fmt.Printf("qty type %T", bigqty)
	
}


