package model

import "errors"

type Comment struct {
	ID       int    `json:"id"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
	UserID   int    `json:"userID"`
	RecipeID int    `json:"recipeID"`
}

var (
	ErrSubjectInvalid  = errors.New("subject is invalid")
	ErrContentInvalid  = errors.New("content is invalid")
	ErrRecipeIDInvalid = errors.New("recipeID is invalid")
)

type AddComment struct {
	Subject  string `json:"subject"`
	Content  string `json:"content"`
	UserID   int    `json:"userID"`
	RecipeID int    `json:"recipeID"`
}

func (a AddComment) Validation() error {
	UserExists := false
	RecipeExists := false
	for _, user := range users {
		if user.ID == a.UserID {
			UserExists = true
		}
	}
	for _, recipe := range recipes {
		if recipe.ID == a.RecipeID {
			RecipeExists = true
		}
	}
	switch {
	case a.Subject == "":
		return ErrSubjectInvalid
	case a.Content == "":
		return ErrContentInvalid
	case !UserExists:
		return ErrUserIDInvalid
	case !RecipeExists:
		return ErrRecipeIDInvalid
	}
	return nil
}

func CommentOne(id int) (Comment, error) {
	for _, comment := range comments {
		if comment.ID == id {
			return comment, nil
		}
	}
	return Comment{}, ErrNotFound
}

type UpdateComment struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (u UpdateComment) Validation() error {
	switch {
	case u.ID == 0:
		return ErrIDInvalid
	case u.Subject == "":
		return ErrSubjectInvalid
	case u.Content == "":
		return ErrContentInvalid
	}
	return nil
}

func CommentsAll(q string) ([]Comment, error) {
	if q != "" {
		var commentsFiltered []Comment
		for _, comment := range comments {
			if comment.Subject == q {
				commentsFiltered = append(commentsFiltered, comment)
			}
		}
		return commentsFiltered, nil
	}
	return comments, nil
}

func (c Comment) InsertComment() (int, error) {
	commentMaxID++
	c.ID = commentMaxID
	comments = append(comments, c)
	return c.ID, nil
}

func (c Comment) UpdateComment() error {
	for i, comment := range comments {
		if comment.ID == c.ID {
			comments[i].Subject = c.Subject
			comments[i].Content = c.Content
			return nil
		}
	}
	return ErrNotFound
}

func DeleteComment(id int) error {
	for i, comment := range comments {
		if comment.ID == id {
			comments = append(comments[:i], comments[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

var commentMaxID = 5
var comments = []Comment{
	{ID: 1, Subject: "Authentic and delicious", Content: "I tried this recipe and it was delicious. I will definitely cook it again.", UserID: 1, RecipeID: 1},
	{ID: 2, Subject: "Great recipe", Content: "I tried this recipe and it was great. I will definitely cook it again.", UserID: 2, RecipeID: 2},
	{ID: 3, Subject: "Amazing", Content: "I tried this recipe and it was amazing. I will definitely cook it again.", UserID: 3, RecipeID: 3},
	{ID: 4, Subject: "Delicious", Content: "I tried this recipe and it was delicious. I will definitely cook it again.", UserID: 1, RecipeID: 4},
	{ID: 5, Subject: "Authentic and delicious", Content: "I tried this recipe and it was delicious. I will definitely cook it again.", UserID: 2, RecipeID: 1},
}
