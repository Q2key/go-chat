package chat

import (
	"bytes"
	"chat/services/core"
	"chat/services/mapper"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Service struct {
	apiKey   string
	endpoint string
	mapper   mapper.ChatServiceMapper
}

func NewChatService(apiKey string) core.Service {
	return &Service{
		apiKey:   apiKey,
		endpoint: core.OpenAICompletionsApiEndpoint,
		mapper:   *new(mapper.ChatServiceMapper),
	}
}

func (ch *Service) formRequest(model string, messages *[]core.Message) *core.ChatRequest {
	return &core.ChatRequest{
		Model:    model,
		Messages: messages,
	}
}

func (ch *Service) getChatGptResponse(request *core.ChatRequest) (*http.Response, error) {
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

func (ch *Service) Execute(model string, r *[]core.Message) (*core.ChatResponse, error) {
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
