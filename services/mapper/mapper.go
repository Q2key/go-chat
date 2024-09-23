package mapper

import (
	goutils "chat/pkg/go-utils"
	"chat/services/core"
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
