package mock

import (
	goutils "chat/pkg/go-utils"
	"chat/services/core"
	"chat/services/mapper"
)

type Service struct {
	apiKey   string
	endpoint string
	mapper   mapper.ChatServiceMapper
}

func (s *Service) Execute(_ string, _ *[]core.Message) (*core.ChatResponse, error) {
	content, err := goutils.ReadContent("./services/mock/response.json")
	if err != nil {
		return nil, err
	}

	return s.mapper.ToChatRequest([]byte(content))
}

func NewMockService(apiKey string) core.Service {
	return &Service{
		apiKey:   apiKey,
		endpoint: core.TestApiEndpoint,
		mapper:   *new(mapper.ChatServiceMapper),
	}
}
