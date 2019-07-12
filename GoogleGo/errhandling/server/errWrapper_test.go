package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"goLearn/GoogleGo/errhandling/server/fileHandler"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return fileHandler.UserError("User Error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unKnown Error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "No Error")
	return nil
}

func verifyResponse(response *http.Response, code int, msg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != code {
		t.Errorf("expect (%d, %s); got (%d %s)\n", code, msg, response.StatusCode, body)
	}

}

// 表格驱动测试
var data = []struct {
	handler appHandler
	code    int
	msg     string
}{
	{errPanic, 500, "Internal Server Error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{errUserError, 400, "User Error"},
	{noError, 200, "No Error"},
}

// 直接调用函数，使用假的 Request/Response
func TestErrWrapper(t *testing.T) {
	for _, d := range data {
		handler := errWrapper(d.handler)
		request := httptest.NewRequest(http.MethodGet, "localhost:8888", nil)
		response := httptest.NewRecorder()
		handler(response, request)

		verifyResponse(response.Result(), d.code, d.msg, t)
	}
}

// 启动一个server，使用http请求
func TestErrWrapperInServer(t *testing.T) {
	for _, d := range data {
		handler := http.HandlerFunc(errWrapper(d.handler))
		server := httptest.NewServer(handler)

		response, _ := http.Get(server.URL)

		verifyResponse(response, d.code, d.msg, t)

	}
}
