// Copyright 2013 Rcsoft. All rights reserved.
package web

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/rrsoft/guestbook/core"
	"github.com/rrsoft/guestbook/logger"
	"github.com/rrsoft/guestbook/utils"
)

func serve404(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusNotFound)
	// w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// w.Write([]byte("Not Found"))
	http.NotFound(w, r)
}

func serveStatus(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func serveError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func loadPage(title string) (*template.Template, error) {
	filename := "template/" + title + ".html"
	return template.ParseFiles(filename)
	/*body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return template.Must(template.New(title).Parse(body))*/
}

// HandleMainPage page starting from 1
func HandleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		serveStatus(w, http.StatusMethodNotAllowed)
		return
	}

	var page int
	if r.URL.Path != "/" {
		arr := strings.Split(r.URL.Path, "/")
		page, _ = strconv.Atoi(arr[len(arr)-1])
	}
	if page < 1 {
		page = 1
	}
	list, err := core.GetList(page-1, 5)
	if err != nil {
		serveError(w, err)
		return
	}
	count := core.Count()
	pager := utils.NewDataPager(page, 5, 5, count, list)
	pager.LinkHook = func(page int) string {
		return fmt.Sprintf("/%d", page)
	}
	listPage, err := loadPage("book")
	if err != nil {
		serveError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := listPage.Execute(w, pager); err != nil {
		serveError(w, err)
	}
}

// HandleDetail show message details
func HandleDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		serveStatus(w, http.StatusMethodNotAllowed)
		return
	}
	arr := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(arr[len(arr)-1])
	if err != nil {
		serve404(w, r)
		return
	}
	details, err := core.GetDetails(id)
	if err != nil {
		serveError(w, err)
		return
	}
	detailsPage, err := loadPage("details")
	if err != nil {
		serveError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := detailsPage.Execute(w, details); err != nil {
		serveError(w, err)
	}
}

// HandleCommit submit a message
func HandleCommit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		serveStatus(w, http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		serveError(w, err)
		return
	}
	info := &core.Greeting{
		Author:   strings.TrimSpace(r.PostFormValue("author")), // FormValue
		Content:  strings.TrimSpace(r.PostFormValue("content")),
		PostDate: time.Now(),
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if len(info.Author) == 0 || len(info.Content) == 0 {
		w.Write([]byte("author or content can't be empty"))
		return
	}
	if err := core.Comment(info); err != nil {
		serveError(w, err)
	} else {
		logger.Write(strconv.Itoa(info.Id))
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// HandleDelete delete a message
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		serveStatus(w, http.StatusMethodNotAllowed)
		return
	}
	arr := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(arr[len(arr)-1])
	if err != nil {
		serveError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Without the permission of the delete %d", id)
	/*if err := book.Delete(id); err != nil {
		serveError(w, err)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}*/
}
