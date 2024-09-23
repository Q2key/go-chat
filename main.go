package main

import (
	"chat/core"
	u "chat/pkg/go-utils"
	ioservice "chat/services/io-service"
	mockservice "chat/services/mock-service"
	"os"
)

func main() {
	intro, _ := u.ReadContent("./templates/intro.txt")
	outro, _ := u.ReadContent("./templates/outro.txt")
	payload, _ := u.ReadContent("./templates/payload.txt")
	prompt, _ := u.ReadContent("./templates/prompt.txt")

	apiKey := os.Getenv("OPENAI_KEY")
	if len(apiKey) == 0 {
		panic("Please set OPENAI_KEY environment variable")
	}

	service := mockservice.NewMockService(apiKey)
	chatResponse, err := service.Execute(core.GptModel40,
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
	err = ioService.Execute(chatResponse.Id, chatResponse.GetAnswer())

	u.Check(err)
}
