package data

type Post struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

var latestPost = &Post{
	Id:   "1",
	Text: "Hello world!",
}

func GetLatestPost() *Post {
	return latestPost
}

func init() {
}
