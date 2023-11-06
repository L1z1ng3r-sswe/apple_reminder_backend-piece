package models

import "time"

type TGetTask struct {
	ID          int       `json:"id"`	
	Text        string    `json:"text"`
	Timer       time.Time `json:"timer"`
	Result      string    `json:"result"`
	CreatedDate time.Time `json:"created_date"`
	IsDone      bool      `json:"isDone"`	
}


type TAddTask struct {
	Text        string     `json:"text" validate:"required,max=512"`
	Timer       time.Time  `json:"timer" validate:"required"`
}

type TUpdateTask struct {
	Result      string     `json:"result" validate:"required,max=512"`
}