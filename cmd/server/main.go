package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("welcome to the home page "))
}
func snippetview(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this is the snippetview"))
}

func snippetcreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This creating a snippet view "))
}


func main(){
	r:= http.NewServeMux()
	r.HandleFunc("/", home)
	r.HandleFunc("/snippet/view",snippetview)
	r.HandleFunc("/snippet/create", snippetcreate)
	log.Println("server running on port http://localhost:4000")
	err := http.ListenAndServe(":4000",r)
	log.Fatal(err)

}