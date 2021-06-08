package my

import (
	"github.com/jinzhu/gorm"
)

// User mode.
type User struct {
	gorm.Model
	Account  string
	Name     string
	Password string
	Message  string
}

// Post model.
type Post struct {
	gorm.Model
	Address string
	Message string
	UserId  int
	GroupId int
}

// Group model.
type Group struct {
	gorm.Model
	UserId  int
	Name    string
	Message string
}

// Comment model.
type Comment struct {
	gorm.Model
	UserId  int
	PostId  int
	Message string
}

// CommentJoin join model.
type CommentJoin struct {
	Comment
	User
	Post
}
