// -*- coding: utf-8 -*-
// @Date    : 2022-03-25 23:24:15
// @Author  : Elijahxb (xbelijah@gmail.com)
// @Link    : www.junyipan.top
// @Version : 1.0.0

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("你好，世界！--Time: %s", time.Now().String())
	log.Printf("%v\n", s)
	//打印请求主机地址
	log.Println(r.Host)
	//打印请求方法
	log.Println(r.Method)
	//打印请求头信息
	log.Printf("header content:[%v]\n", r.Header)
	//打印post请求form数据
	log.Printf("form content:[username: %v, password: %v]\n", r.PostFormValue("username"), r.PostFormValue("password"))
	//读取请求体
	bodyContent, err := ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		fmt.Printf("read body content failed, errr: [%v]\n", err)
		return
	}
	fmt.Printf("body content:[%v]\n", string(bodyContent))
	fmt.Fprintf(w, "%v\n", s)
}

func main() {
	fmt.Println("Please visit http://127.0.0.1:8000/api/v1/")
	http.HandleFunc("/", HttpHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
