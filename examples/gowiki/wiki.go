/*
Problem #
Author: Corey Prak
Date Created: 04/03/2014
Comments:

This file was created with guifance from www.golang.org, "Writing Web Applications" and is lovated at http://golang.org/doc/articles/wiki/
*/


package main

import (
  "html/template"
  "io/ioutil"
  "net/http"
)

//Body is a byte slice because it's the type expected by io libraries.
type Page struct {
  Title string 
  Body []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"

  // func WriteFile(filename string, data []byte, perm os.FileMode) error
  // Source: golang.org/pkg/io/ioutil/
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"

  // func ReadFile(filename string) ([]byte, error)
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  //Return a page literal.
  return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)

  renderTemplate(w, "view", p)
}

func edithandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }

  renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  t, _ := template.ParseFiles(tmpl + ".html")
  t.Execute(w, p)
}

func main() {
  http.HandleFunc("/view/", viewHandler)
//  http.HandleFunc("/edit/", editHandler)
//  http.HandleFunc("/save/", saveHandler)
  http.ListenAndServe(":8080", nil)
}
