package models

// Line - line in csv file
type Line struct {
	Number  int      `json:"number"`
	Columns []string `json:"columns"`
}
