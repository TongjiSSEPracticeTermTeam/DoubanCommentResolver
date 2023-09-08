package resolver

type Comment struct {
	Cid          string `json:"cid"`
	AuthorId     string `json:"author_id"`
	AuthorName   string `json:"author_name"`
	AuthorAvatar string `json:"author_avatar"`
	Rate         int    `json:"rate"`
	Date         string `json:"date"`
	Content      string `json:"content"`
	Vote         int    `json:"upVote"`
}
