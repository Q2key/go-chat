package chat

import (
	"bytes"
	core2 "chat/core"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Service struct {
	apiKey   string
	endpoint string
	mapper   core2.ChatServiceMapper
}

func NewChatService(apiKey string) core2.Service {
	return &Service{
		apiKey:   apiKey,
		endpoint: core2.OpenAICompletionsApiEndpoint,
		mapper:   *new(core2.ChatServiceMapper),
	}
}

func (ch *Service) formRequest(model string, messages *[]core2.Message) *core2.ChatRequest {
	return &core2.ChatRequest{
		Model:    model,
		Messages: messages,
	}
}

func (ch *Service) getChatGptResponse(request *core2.ChatRequest) (*http.Response, error) {
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

func (ch *Service) Execute(model string, r *[]core2.Message) (*core2.ChatResponse, error) {
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
