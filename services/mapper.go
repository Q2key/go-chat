package services

import (
	"chat/core"
	goutils "chat/pkg/go-utils"
	"encoding/json"
)

type ChatServiceMapper struct {
}

func (m *ChatServiceMapper) ToChatRequest(bytes []byte) (*core.ChatResponse, error) {
	var t core.ChatResponse
	err := json.Unmarshal(bytes, &t)
	goutils.Check(err)
	return &t, nil
}
