package aichat

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

var (
	// TODO: aichat/_template/*.htmlをテンプレートとしてパースする

)

func (chat *AIChat) initHandlers() {
	chat.mux.HandleFunc("/", chat.HandleIndex)
	chat.mux.HandleFunc("/add", chat.HandleAdd)
}

func (chat *AIChat) HandleIndex(w http.ResponseWriter, r *http.Request) {
	ps, err := chat.Posts(r.Context(), 10)
	if err != nil {
		chat.error(w, err, http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(/* TODO: 結果をレスポンスとして返す */, "index", ps); err != nil {
		chat.error(w, err, http.StatusInternalServerError)
		return
	}
}

func (chat *AIChat) HandleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("MethodがPOSTではありません")
		chat.error(w, err, http.StatusMethodNotAllowed)
		return
	}

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

	http.Redirect(w, r, "/", http.StatusFound)
}

func (chat *AIChat) error(w http.ResponseWriter, err error, code int) {
	log.Println("Error:", err)
	http.Error(w, http.StatusText(code), code)
}
