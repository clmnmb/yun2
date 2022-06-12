package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func a(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//fmt.Printf("%T",a1)
		dataType , _ := json.Marshal(r.Header)
		dataString := string(dataType)
		fmt.Fprintf(w, dataString)

	}
}

func b(w http.ResponseWriter, r *http.Request)  {
	cmd := exec.Command("go","version")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprintf(w, string(stdout))

}

func healthz(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "200")
	}
}

func main() {
	http.HandleFunc("/a", a)
	http.HandleFunc("/b", b)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
