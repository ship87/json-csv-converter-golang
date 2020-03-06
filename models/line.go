package models

type Line struct {
	Number  int      `json:"number"`
	Columns []string `json:"columns"`
}
