package main

import (
	"chat/core"
	u "chat/pkg/go-utils"
	"chat/services"
)

func main() {

	intro, _ := u.ReadContent("./templates/intro.txt")
	outro, _ := u.ReadContent("./templates/outro.txt")
	payload, _ := u.ReadContent("./templates/payload.txt")
	prompt, _ := u.ReadContent("./templates/prompt.txt")

	openAiService := services.NewServiceFactory().MakeOpenAiService()
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
	io := services.NewTextService()
	u.Check(io.Execute(chatResponse.Id, chatResponse.GetAnswer()))
}
