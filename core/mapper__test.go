package core_test

import (
	"chat/core"
	goutils "chat/pkg/go-utils"
	"testing"
)

func TestParseRequest(t *testing.T) {
	bytes, err := goutils.ReadContent("./mock-data/response.json")
	if err != nil {
		t.Error(err)
	}

	m := new(core.ChatServiceMapper)

	cr, err := m.ToChatRequest([]byte(bytes))
	if err != nil {
		t.Error(err)
	}

	if cr.Id != "chatcmpl-AAbcjlI2o36zxemt9tlXLrP0544Jo" {
		t.Error("Wrong Id")
	}

	if len(cr.Choices) == 0 {
		t.Error("Wrong choices")
	}

	if len(cr.GetAnswer()) == 0 {
		t.Error("Wrong answer")
	}

}
