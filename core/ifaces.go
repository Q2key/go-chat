package core

type IOService interface {
	Execute(fileName, fileContent string) error
}

type Service interface {
	Execute(model string, r *[]Message) (*ChatResponse, error)
}
