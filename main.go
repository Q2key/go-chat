package main

import (
	"chat/core"
	u "chat/pkg/go-utils"
	"chat/services"
)

func main() {

	m1, _ := u.ReadContent("./templates/intro.txt")
	m2, _ := u.ReadContent("./templates/outro.txt")
	m3, _ := u.ReadContent("./templates/payload.txt")
	m4, _ := u.ReadContent("./templates/prompt.txt")

	openAiService := services.NewServiceFactory().MakeOpenAiService()
	chatResponse, err := openAiService.Execute(core.GptModel40,
		&[]core.Message{
			{
				core.GptRoleUser,
				m1,
			},
			{
				core.GptRoleUser,
				m2,
			},
			{
				core.GptRoleUser,
				m3,
			},
			{
				core.GptRoleUser,
				m4,
			},
		})

	u.Check(err)
	io := services.NewTextService("./build")
	u.Check(io.Execute(chatResponse.Id, chatResponse.GetAnswer()))
}
