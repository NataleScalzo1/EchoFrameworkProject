package main

import "time"

type Film struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	Title   string    `json:"titlename"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Expires string    `json:"expires"`
}
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}
