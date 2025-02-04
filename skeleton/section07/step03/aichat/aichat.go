package aichat

import (
	"context"
	"net/http"
	"slices"

	"github.com/openai/openai-go"
)

type Post struct {
	Name    string
	Message string
}

type AIChat struct {
	openai *openai.Client
	posts  []*Post // NOTE: 本来であれば排他制御が必要
	mux    *http.ServeMux
	server *http.Server
}

func New(addr string) (*AIChat, error) {
	// TODO: *http.ServeMux型の値を生成しmuxという変数に代入

	return &AIChat{
		openai: openai.NewClient(),
		mux:    mux,
		server: /* TODO: AddrフィールドがaddrでHandlerフィールドがmuxの*http.Server型の値を生成 */,
	}, nil
}

func (chat *AIChat) Start() error {
	chat.initHandlers()
	if err := /* TODO: serverフィールドのListenAndServeを呼ぶ */; err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (chat *AIChat) Posts(ctx context.Context, limit int) ([]*Post, error) {
	return slices.Clone(chat.posts), nil
}

func (chat *AIChat) AddPost(ctx context.Context, p *Post) error {

	reply, err := chat.reply(ctx, p)
	if err != nil {
		return err
	}

	chat.posts = append(chat.posts, p, reply)

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
