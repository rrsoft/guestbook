package main

import (
	"log"
	"net/http"

	"github.com/rrsoft/guestbook/core"
	"github.com/rrsoft/guestbook/data"
)

func main() {
	defer data.Close()
	http.HandleFunc("/", core.HandleMainPage)
	http.HandleFunc("/sign", core.HandleSign)
	http.HandleFunc("/details/", core.HandleDetails)
	http.HandleFunc("/delete/", core.HandleDelete)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
