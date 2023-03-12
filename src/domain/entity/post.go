package entity

type Post struct {
	PostId      string `json:"postId"`
	Description string `json:"description"`
	PostDate    string `json:"postDate"`
	Receiver    string `json:"receiver"`
}

type posts []Post

var Posts posts
