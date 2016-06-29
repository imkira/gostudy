package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// START OMIT
func RedirectServer(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "http://127.0.0.1:12345", 301)
}

func main() {
	http.HandleFunc("/", RedirectServer)
	go http.ListenAndServe("127.0.0.1:12345", nil)
	time.Sleep(500 * time.Millisecond)
	resp, err := http.Get("http://127.0.0.1:12345")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// END OMIT
