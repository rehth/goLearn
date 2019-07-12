package fileHandler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const Prefix = "/list/"

type UserError string

func (e UserError) Error() string {
	return string(e)
}

func Handler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, Prefix) != 0 {
		return UserError("path must start with " + Prefix)
	}
	path := request.URL.Path[len(Prefix):]
	log.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(data)
	return nil
}
