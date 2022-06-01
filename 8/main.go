package main

import (
	"fmt"
)

func main() {
	book := map[string]map[string][]string{
		"Anna": {
			"books": []string{"Golang","1ASS","Three hundred bucks"},
		}, "Jhohan": {
			"books": []string{"Golang Ninja","Basic","Pascal"},
		}, "Vladislav": {
			"books":  []string{"Sis Admin","Nice cock","great view"},
		},
			"Raul": {
			"books": []string{"Tridsat pyati dney","great view on you"},
		},
	}
	for i:= range(book){
		for g:= range(book[i]){
			for j:= range (book[i][g]){
				fmt.Println(book[i][g][j])
			}
		}
	}

}
