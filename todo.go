package main

import "time" 


type Todo struct{
	Title string
	Complated bool
	CreateAt time.Time
	ComplatetAt *time.Time
}

