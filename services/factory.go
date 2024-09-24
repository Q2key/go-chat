package services

import (
	"chat/core"
	"os"
)

type ServiceFactory struct {
}

func NewServiceFactory() *ServiceFactory {
	return &ServiceFactory{}
}

func (factory *ServiceFactory) MakeOpenAiService() core.Service {
	apiKey := os.Getenv("OPENAI_KEY")
	if len(apiKey) == 0 {
		panic("Please set OPENAI_KEY environment variable")
	}

	return NewChatService(apiKey)
}

func (factory *ServiceFactory) MakeMockService() core.Service {
	responseJson := os.Getenv("MOCK_RESPONSE_JSON")
	if len(responseJson) == 0 {
		panic("Please set MOCK_RESPONSE_JSON environment variable")
	}

	return NewMockService("DUMMY_API_KEY", responseJson)
}
