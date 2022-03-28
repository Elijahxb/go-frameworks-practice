package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{ "data": [ "books", "movies" ] }`)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/api/v1/", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Apipost client Runtime/+https://www.apipost.cn/")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
