package web

type CommentCreate struct {
	Author string `json:"author"`
	Content string `json:"content"`
}

type GetComment struct {
	ID string `json:"id"`
	Author string `json:"author"`
	Content string `json:"content"`
}

type DeleteComment struct {
	ID string `json:"id"`
	Content string `json:"content"`
}