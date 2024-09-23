package factory

import (
	"chat/services/chat"
	"chat/services/core"
	"chat/services/mock"
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

	return chat.NewChatService(apiKey)
}

func (factory *ServiceFactory) MakeMockService() core.Service {
	return mock.NewMockService("DUMMY_API_KEY")
}
