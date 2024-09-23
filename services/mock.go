package services

import (
	"chat/core"
	goutils "chat/pkg/go-utils"
)

type MockService struct {
	apiKey   string
	endpoint string
	mapper   core.ChatServiceMapper
}

func (s *MockService) Execute(_ string, _ *[]core.Message) (*core.ChatResponse, error) {
	content, err := goutils.ReadContent("./services/response.json")
	if err != nil {
		return nil, err
	}

	return s.mapper.ToChatRequest([]byte(content))
}

func NewMockService(apiKey string) core.Service {
	return &MockService{
		apiKey:   apiKey,
		endpoint: core.TestApiEndpoint,
		mapper:   *new(core.ChatServiceMapper),
	}
}
