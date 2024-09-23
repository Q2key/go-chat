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
	return NewMockService("DUMMY_API_KEY")
}
