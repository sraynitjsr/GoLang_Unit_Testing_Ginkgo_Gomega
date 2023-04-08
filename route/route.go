package route

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var posts []Post

func init() {
	post := Post{
		ID:    1,
		Title: "Title One",
		Text:  "Text One",
	}
	posts = append(posts, post)
}

func Home(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(`{"Welcome Home"}`))
}

func GetPosts(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(`{"error":"error marshalling the posts slice"}`))
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(result)
}
