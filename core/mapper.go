package core

import (
	goutils "chat/pkg/go-utils"
	"encoding/json"
)

type ChatServiceMapper struct {
}

func (m *ChatServiceMapper) ToChatRequest(bytes []byte) (*ChatResponse, error) {
	var t ChatResponse
	err := json.Unmarshal(bytes, &t)
	goutils.Check(err)
	return &t, nil
}
