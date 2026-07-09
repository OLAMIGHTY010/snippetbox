package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}
		w.Write([]byte("welcome to the home page "))

}
func snippetview(w http.ResponseWriter, r *http.Request){
	id,err :=strconv.Atoi(r.URL.Query().Get("id"))
	if err !=nil || id < 1{
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w,"Display snippett view ID %d",id)

}

func snippetcreate(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Allow",http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
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