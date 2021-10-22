package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHi(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./artical.txt")
	fmt.Fprintln(w, string(b))
}
func main() {
	http.HandleFunc("/hello", sayHi)
	err:=http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
	}
}
