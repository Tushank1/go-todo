package models

type Todo struct {
	ID        uint   `gorum:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
