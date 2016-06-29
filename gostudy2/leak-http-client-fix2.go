package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// START OMIT
func main() {
	resp, err := http.Get("https://www.random.org/integers/?num=5&min=1&max=10&col=1&base=10&format=plain")
	if resp != nil {
		defer resp.Body.Close()
		defer io.Copy(ioutil.Discard, resp.Body)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
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
