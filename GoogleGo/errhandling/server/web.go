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

			if e, ok := err.(fileHandler.UserError); ok {
				http.Error(writer, e.Error(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK

			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
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
