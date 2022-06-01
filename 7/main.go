package main

import (
	"fmt"; "net/http"; "html/template"
)
type Gym struct{
	slave_name string
	age uint8
	cock float32 
}
func (g Gym) Getallinfo() string{
	return fmt.Sprintf("Name of slave %s his age %d and most important his cock is %f",g.slave_name, g.age,g.cock)
}
func (g *Gym) Change_name(Newname string) {
	g.slave_name = Newname
}
func home_page(w http.ResponseWriter, r *http.Request){
	Master := Gym {}
	tmpl,_ := template.ParseFiles("templates/site2.html")
	tmpl.ExecuteTemplate(w)
}
func contact_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"awsome balls")
}

func HandleRequest(){
	http.HandleFunc("/",home_page)
	http.HandleFunc("/contacts/",contact_page)
	http.ListenAndServe(":8080",nil)

}
func main() {
	HandleRequest()
}
