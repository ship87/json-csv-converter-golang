package models

// Line in csv file
type Line struct {
	Number  int      `json:"number"`
	Columns []string `json:"columns"`
}
