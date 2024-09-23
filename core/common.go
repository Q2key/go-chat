package core

import (
	chatservice "chat/services/chat-service"
	mockservice "chat/services/mock-service"
	"os"
)

const (
	OpenAICompletionsApiEndpoint = "https://api.openai.com/v1/chat/completions"
	TestApiEndpoint              = "https://postman-echo.com/post"
)

const (
	GptModel40 = "gpt-4o"
)

const (
	GptRoleUser   = "user"
	GptRoleSystem = "system"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string     `json:"model"`
	Messages *[]Message `json:"messages"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Refusal *bool  `json:"refusal",omitempty`
}

type ResponseChoice struct {
	Index        int             `json:"index"`
	Message      ResponseMessage `json:"message"`
	Logprobs     *bool           `json:"logprobs", omitempty`
	FinishReason string          `json:"finish_reason"`
}

type ChatResponse struct {
	Id      string           `json:"id"`
	Object  string           `json:"object"`
	Created int              `json:"created"`
	Model   string           `json:"model"`
	Choices []ResponseChoice `json:"choices"`
}

func (ch *ChatResponse) GetAnswer() string {
	if len(ch.Choices) > 0 {
		return ch.Choices[0].Message.Content
	}

	return ""
}

type ServiceFactory struct {
}

func NewServiceFactory() *ServiceFactory {
	return &ServiceFactory{}
}

func (factory *ServiceFactory) MakeOpenAiService() Service {
	apiKey := os.Getenv("OPENAI_KEY")
	if len(apiKey) == 0 {
		panic("Please set OPENAI_KEY environment variable")
	}

	return chatservice.NewChatService(apiKey)
}

func (factory *ServiceFactory) MakeMockService() Service {
	return mockservice.NewMockService("DUMMY_API_KEY")
}
