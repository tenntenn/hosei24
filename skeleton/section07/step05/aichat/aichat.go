package aichat

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/openai/openai-go"
	"github.com/tenntenn/sqlite"
)

type Post struct {
	Name    string
	Message string
}

type AIChat struct {
	openai *openai.Client
	db     *sql.DB
	mux    *http.ServeMux
	server *http.Server
}

func New(addr string) (*AIChat, error) {
	mux := http.NewServeMux()
	// TODO: ドライバ名にsqlite.DriverName、接続文字列に"aichat.db"を指定してデータベースを開く

	if err != nil {
		return nil, err
	}

	return &AIChat{
		openai: openai.NewClient(),
		mux:    mux,
		db:     db,
		server: &http.Server{Addr: addr, Handler: mux},
	}, nil
}

func (chat *AIChat) Start() error {
	if err := chat.initDB(context.Background()); err != nil {
		return err
	}
	chat.initHandlers()
	if err := chat.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (chat *AIChat) initDB(ctx context.Context) error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS posts(
		id	INTEGER PRIMARY KEY,
		name 	TEXT NOT NULL,
		message	TEXT NOT NULL
	);`

	if _, err := chat.db.ExecContext(ctx, sqlStr); err != nil {
		return err
	}

	return nil
}

func (chat *AIChat) Posts(ctx context.Context, limit int) ([]*Post, error) {
	const sqlStr = `SELECT name, message FROM posts LIMIT ?`
	rows, err := chat.db.QueryContext(ctx, sqlStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 関数終了時にCloseが呼び出される

	var ps []*Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Name, &p.Message)
		if err != nil {
			return nil, err
		}
		ps = append(ps, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ps, nil
}

func (chat *AIChat) AddPost(ctx context.Context, p *Post) error {

	reply, err := chat.reply(ctx, p)
	if err != nil {
		return err
	}

	const sqlStr = `INSERT INTO posts(name, message) VALUES (?,?), (?, ?);`
	_, err = chat.db.ExecContext(ctx, sqlStr, p.Name, p.Message, reply.Name, reply.Message)
	if err != nil {
		return err
	}

	return nil
}

func (chat *AIChat) reply(ctx context.Context, p *Post) (*Post, error) {
	completion, err := chat.openai.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(p.Message),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	})

	if err != nil {
		return nil, err
	}

	return &Post{
		Name:    "AI",
		Message: completion.Choices[0].Message.Content,
	}, nil
}
