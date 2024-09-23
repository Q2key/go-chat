package services

import (
	"bytes"
	"chat/core"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ChatService struct {
	apiKey   string
	endpoint string
	mapper   core.ChatServiceMapper
}

func NewChatService(apiKey string) core.Service {
	return &ChatService{
		apiKey:   apiKey,
		endpoint: core.OpenAICompletionsApiEndpoint,
		mapper:   *new(core.ChatServiceMapper),
	}
}

func (ch *ChatService) formRequest(model string, messages *[]core.Message) *core.ChatRequest {
	return &core.ChatRequest{
		Model:    model,
		Messages: messages,
	}
}

func (ch *ChatService) getChatGptResponse(request *core.ChatRequest) (*http.Response, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(b)

	r, err := http.NewRequest(http.MethodPost, ch.endpoint, buff)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ch.apiKey))

	response, err := http.DefaultClient.Do(r)

	log.Printf("\r\nStatus Code: %d", response.StatusCode)

	return response, err
}

func (ch *ChatService) Execute(model string, r *[]core.Message) (*core.ChatResponse, error) {
	request := ch.formRequest(model, r)

	response, err := ch.getChatGptResponse(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return ch.mapper.ToChatRequest(body)
}
