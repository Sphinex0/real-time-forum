package models

import (
	"errors"
)

type Comment struct {
	ID        int64  `json:"id"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt int    `json:"created_at"`
}

type CommentWithUser struct {
	Comment
	GetLikes
	User Members `json:"user"`
}

// BeforCreateComment validates a comment's content length and ensures
// it is associated with a post. Returns an error if validation fails.
func (comment *Comment) BeforCreateComment() error {
	if length(1, 2000, comment.Content) || comment.PostID == 0 {
		return errors.New("size not allowed")
	}
	return nil
}
