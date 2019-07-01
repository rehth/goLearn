package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#foo"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("URL:", u.Scheme, u.User, u.Host, u.Path, u.RawQuery, u.Fragment)
	// url 用户信息解析
	username := u.User.Username()
	pwd, _ := u.User.Password()
	fmt.Println("User:", username, pwd)

	// url 主机信息解析
	fmt.Println("Host:", u.Hostname(), u.Port())

	// url query参数解析
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("Query:", m)
}
