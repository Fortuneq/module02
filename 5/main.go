package main

import (
	"fmt"
)

func main(){
	var s[] string
	s = []string{"monday","tusday","wednesday","thursday","friday","sunday","saturday"}
	s1 := s[0:5]
	s2 := s[5:]
	fmt.Println(s1)
	fmt.Println(s2)
	
	s3 :=append(s1,s2...)
	fmt.Println(s3)
}