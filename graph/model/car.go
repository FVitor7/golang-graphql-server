package model

type Car struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand *Brand `json:"brand"`
}