package model

import "time"

type Todo struct {
	Id          int
	Description string
	Done        bool
	CreatedAt  time.Time
}
