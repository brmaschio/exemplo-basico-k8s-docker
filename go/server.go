package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
  
	http.ListenAndServe(":8000", nil) 

}

func index(w http.ResponseWriter, req *http.Request) {

  message := greeting("Code.education Rocks!!!")

  t, _ := template.ParseFiles("index.html")

  data := map[string]string{
    "Title": message,
  }

  w.WriteHeader(http.StatusOK)
  
  t.Execute(w, data)

}

func greeting(str string)string{
	return str 
}