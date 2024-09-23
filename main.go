package main

import (
	u "chat/pkg/go-utils"
	"chat/services/core"
	"chat/services/factory"
	"fmt"
)

func main() {

	intro, _ := u.ReadContent("./templates/intro.txt")
	outro, _ := u.ReadContent("./templates/outro.txt")
	payload, _ := u.ReadContent("./templates/payload.txt")
	prompt, _ := u.ReadContent("./templates/prompt.txt")

	openAiService := factory.NewServiceFactory().MakeMockService()
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

	fmt.Println(chatResponse)

	u.Check(err)
}
