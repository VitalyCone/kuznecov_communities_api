package model

type Post struct{
	ID int
	Text string
	Likes int
	Files [][]byte
}