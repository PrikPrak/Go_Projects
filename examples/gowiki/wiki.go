/*
Problem #
Author: Corey Prak
Date Created: 04/03/2014
Comments:

This file was created with guifance from www.golang.org, "Writing Web Applications" and is lovated at http://golang.org/doc/articles/wiki/
*/


package main

import (
	"fmt"
  "io/ioutil"
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

func main() {
  p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
  p1.save()

  p2, _ := loadPage("TestPage")
  fmt.Println(string(p2.Body))
}
