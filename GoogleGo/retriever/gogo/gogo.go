package gogo

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	Contents  string
	UserAgent string
	TimeOut   time.Duration
}

// Reader/Writer Stringer
func (r *Retriever) String() string {
	return fmt.Sprintf("<Retriever:Contents=%s>", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) {
	r.Contents = form["contents"]
}

func (r *Retriever) Connect(host string) string {
	return r.Contents
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	res, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(res)
}
