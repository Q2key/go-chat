package main

import (
	"chat/core"
	u "chat/pkg/go-utils"
	"chat/services"
)

func main() {

	openAiService := services.NewServiceFactory().MakeOpenAiService()
	chatResponse, err := openAiService.Execute(core.GptModel40,
		&[]core.Message{
			{
				core.GptRoleUser,
				"write a haiku about ai",
			},
		})

	u.Check(err)
	io := services.NewTextService()
	u.Check(io.Execute(chatResponse.Id, chatResponse.GetAnswer()))
}
