package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://g.cn/robots.txt")
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read error: %s", err)
		return
	}
	fmt.Printf("%s\n", string(contents))
}
