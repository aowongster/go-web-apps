package main

import (
  "fmt"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func main() {
  r := httprouter.New()
  r.GET("/", HomeHandler)

  // posts collection
  r.GET("/posts", PostsIndexHandler)
  r.POST("/posts", PostsCreateHandler)

  // Posts singular
  r.GET("/posts/:id", PostShowHandler)
  r.PUT("/posts/:id", PostUpdatehandler)
  r.GET("/posts/:id/edit", PostEditHandler)

  fmt.Println("starting server on :8080")
  http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request,
p httprouter.Params) {
  fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Reqest,
  p httprouter.Params) {
  fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Reqest,
  p httprouter.Params) {
  fmt.Fprintln(rw, "posts create")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Reqest,
  p httprouter.Params) {
  id := p.ByName("id")
  fmt.Fprintln(rw, "showing post", id)
}
func PostUpdateHandler(rw http.ResponseWriter, r *http.Reqest,
  p httprouter.Params) {
  fmt.Fprintln(rw, "post update")
}
func PostDeleteHandler(rw http.ResponseWriter, r *http.Reqest,
  p httprouter.Params) {
  fmt.Fprintln(rw, "post delete")
}
func PostEditHandler(rw http.ResponseWriter, r *http.Reqest,
  p httprouter.Params) {
  fmt.Fprintln(rw, "post edit")
}
