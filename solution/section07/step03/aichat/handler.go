package aichat

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (chat *AIChat) initHandlers() {
	chat.mux.HandleFunc("/", chat.HandleIndex)
}

func (chat *AIChat) HandleIndex(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name == "" {
		err := errors.New("名前が指定されていません")
		chat.error(w, err, http.StatusBadRequest)
		return
	}

	msg := r.FormValue("message")
	if msg == "" {
		err := errors.New("メッセージが指定されていません")
		chat.error(w, err, http.StatusBadRequest)
		return
	}

	p := &Post{
		Name:    name,
		Message: msg,
	}

	if err := chat.AddPost(r.Context(), p); err != nil {
		chat.error(w, err, http.StatusInternalServerError)
		return
	}

	for _, p := range chat.posts {
		if _, err := fmt.Fprintf(w, "%s:\n\t%s\n", p.Name, p.Message); err != nil {
			chat.error(w, err, http.StatusInternalServerError)
			return
		}
	}
}

func (chat *AIChat) error(w http.ResponseWriter, err error, code int) {
	log.Println("Error:", err)
	http.Error(w, http.StatusText(code), code)
}
