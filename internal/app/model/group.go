package model

type Group struct{
	ID int
	Name string
	Discription string
	Avatar []byte
	Posts []Post
	Photos [][]byte
	Videos [][]byte
	Followers []Follower
}

type Follower struct{

}
