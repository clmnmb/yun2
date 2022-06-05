package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type Fs_in struct {
	Challenge string
	Token     string
	Ttype     string
}

type Fs_out struct {
	Challenge string
}

// type test struct {
// 	Uuid string
// 	Token string
// 	Ts string
// 	Type string
// 	Event test_event
// }

// type test_event struct {
// 	Type string
// 	App_id string
// 	Tenant_key string
// 	Root_id string
// 	Parent_id string
// 	Open_chat_id string
// 	Chan_type string
// 	Msg_type string
// 	Open_id string
// 	Open_message_id string
// 	Is_mention string
// 	Text string
// 	Text_without_at_bot string
// 	Chat_type string
// 	Mg_type string
// }

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
