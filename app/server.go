package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"time"
	"io/ioutil"
)

var startedAt = time.Now()

func main(){
	http.HandleFunc("/healthz",  Healthz)
	http.HandleFunc("/configmap",  ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)

}

func Hello(w http.ResponseWriter, r *http.Request){
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello I'm %s and have %s year's old", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}


func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}

}