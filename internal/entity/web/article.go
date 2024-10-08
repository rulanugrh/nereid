package web

type ArticleCreate struct {
	Title   string   `json:"title" form:"title"`
	Content string   `json:"content" form:"content"`
	Tags    []string `json:"tags" form:"tags"`
	Author  string   `json:"author" form:"author"`
}

type GetArticle struct {
	ID          string   `json:"id" form:"id"`
	Title       string   `json:"title" form:"title"`
	Content     string   `json:"content" form:"content"`
	Description string   `json:"description" form:"description"`
	Author      string   `json:"author" form:"author"`
	Tags        []string `json:"tags" form:"tags"`
	Banner      string   `json:"banner" form:"banner"`
	CreatedAt   string   `json:"create_at" form:"create_at"`
	Conclusion  string   `json:"conclusion" form:"conclusion"`
}

type ArticleDelete struct {
	ID    string `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
}
