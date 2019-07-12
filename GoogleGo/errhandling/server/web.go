package main

import (
	"log"
	"net/http"
	"os"

	"goLearn/GoogleGo/errhandling/server/fileHandler"
)

//const prefix = "/"
const prefix = fileHandler.Prefix

type appHandler func(http.ResponseWriter, *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic:", err)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Println("Error handling request:", err.Error())
			code := http.StatusOK

			switch err.(type) {
			case *os.PathError:
				if os.IsNotExist(err) {
					code = http.StatusNotFound
				} else if os.IsPermission(err) {
					code = http.StatusForbidden
				} else {
					code = http.StatusBadRequest
				}

			case fileHandler.UserError:
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc(prefix, errWrapper(fileHandler.Handler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
