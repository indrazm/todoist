package models

type Task struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	Body	  string `json:"body"`
	IsCompleted bool `json:"is_completed"`
}