package mock

import (
	core2 "chat/core"
	goutils "chat/pkg/go-utils"
)

type Service struct {
	apiKey   string
	endpoint string
	mapper   core2.ChatServiceMapper
}

func (s *Service) Execute(_ string, _ *[]core2.Message) (*core2.ChatResponse, error) {
	content, err := goutils.ReadContent("./services/mock/response.json")
	if err != nil {
		return nil, err
	}

	return s.mapper.ToChatRequest([]byte(content))
}

func NewMockService(apiKey string) core2.Service {
	return &Service{
		apiKey:   apiKey,
		endpoint: core2.TestApiEndpoint,
		mapper:   *new(core2.ChatServiceMapper),
	}
}
