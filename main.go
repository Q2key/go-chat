package main

import (
	"chat/core"
	u "chat/pkg/go-utils"
	chatservice "chat/services/chat-service"
	ioservice "chat/services/io-service"
	mockservice "chat/services/mock-service"
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

	return chatservice.NewChatService(apiKey)
}

func (factory *ServiceFactory) MakeMockService() core.Service {
	return mockservice.NewMockService("DUMMY_API_KEY")
}

func main() {

	intro, _ := u.ReadContent("./templates/intro.txt")
	outro, _ := u.ReadContent("./templates/outro.txt")
	payload, _ := u.ReadContent("./templates/payload.txt")
	prompt, _ := u.ReadContent("./templates/prompt.txt")

	openAiService := NewServiceFactory().MakeOpenAiService()
	chatResponse, err := openAiService.Execute(core.GptModel40,
		&[]core.Message{
			{
				core.GptRoleUser,
				intro,
			},
			{
				core.GptRoleUser,
				outro,
			},
			{
				core.GptRoleUser,
				payload,
			},
			{
				core.GptRoleUser,
				prompt,
			},
		})

	u.Check(err)

	ioService := ioservice.NewTextService()

	coverLetterContent := intro
	coverLetterContent += "---"
	coverLetterContent += chatResponse.GetAnswer()
	coverLetterContent += "---"
	coverLetterContent += outro

	err = ioService.Execute(chatResponse.Id, coverLetterContent)

	u.Check(err)
}
