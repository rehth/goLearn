package main

import (
	"fmt"
	"time"

	"goLearn/GoogleGo/retriever/gogo"
)

const url = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string)
}

type Connector interface {
	Poster
	Retriever
	Connect(host string) string
}

func session(c Connector) string {
	c.Post(url, map[string]string{
		"contents": "This is Post Form",
	})
	return c.Connect(url)
}

func download(r Retriever) string {
	return r.Get(url)
}

func main() {
	var r Retriever
	r = &gogo.Retriever{Contents: "this go learning...", UserAgent: "Mozilla/5.0", TimeOut: time.Second}
	req := gogo.Retriever{Contents: "this is Retriever..."}
	fmt.Printf("%T %v\n", r, r)

	// Type Assertion
	if v, ok := r.(*gogo.Retriever); ok {
		fmt.Println("*gogo.Retriever OK:", v)

	}

	// Type Switch
	switch v := r.(type) {
	case *gogo.Retriever:
		fmt.Println("*gogo.Retriever Instance:", v)
	default:
		fmt.Println("Not Find Type")
	}
	//fmt.Println(download(r))
	fmt.Println("Try a session:", session(&req))
	fmt.Println(req, &req)
}
