package gpt

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

// структура для работы с API GPT
type Gpt struct {
	token  string
	client *openai.Client
}

func New(token string) *Gpt {
	return &Gpt{
		token:  token,
		client: openai.NewClient(token),
	}
}

// Функция отправки текстового сообщения GPT API
// Если будет указан параметр fileUrl, то бот отправит реакцию на фото
func (gpt *Gpt) Send(msg string, fileUrl string) (string, error) {
	model := openai.GPT3Dot5Turbo
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: msg,
		},
	}

	if fileUrl != "" {
		model = openai.GPT4o
		message = []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    fileUrl,
							Detail: "high",
						},
					},
				},
			},
		}
	}

	resp, err := gpt.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    model,
		Messages: message,
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
