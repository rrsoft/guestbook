package main

import (
	"log"
	"net/http"

	"github.com/rrsoft/guestbook/data"
	"github.com/rrsoft/guestbook/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defer data.Close()
	http.HandleFunc("/", web.HandleMainPage)
	http.HandleFunc("/sign", web.HandleSign)
	http.HandleFunc("/details/", web.HandleDetails)
	http.HandleFunc("/delete/", web.HandleDelete)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
