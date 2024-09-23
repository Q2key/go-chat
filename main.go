package main

import (
	"chat/core"
	u "chat/pkg/go-utils"
	ioservice "chat/services/io-service"
)

func main() {

	intro, _ := u.ReadContent("./templates/intro.txt")
	outro, _ := u.ReadContent("./templates/outro.txt")
	payload, _ := u.ReadContent("./templates/payload.txt")
	prompt, _ := u.ReadContent("./templates/prompt.txt")

	openAiService := core.NewServiceFactory().MakeOpenAiService()
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
