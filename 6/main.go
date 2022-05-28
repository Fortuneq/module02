package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	path := "6.txt"
	getFile(path)


}
func getFile(path string){
	f,_ := os.Open(path)
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan(){
		fmt.Println(s.Text())
	}
}