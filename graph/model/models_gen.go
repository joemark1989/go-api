// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type NewBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}
