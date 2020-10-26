package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apriendeau/basic-go-api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	testcases := []struct {
		description string
		code        int
		message     handlers.Message
	}{
		{
			description: "success",
			code:        200,
			message: handlers.Message{
				Message: "hello",
			},
		},
	}

	for i, test := range testcases {
		test := test
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Log(test.description)

			req, err := http.NewRequest("GET", "/", nil)
			assert.NoError(t, err)

			w := httptest.NewRecorder()

			h := handlers.New()

			h.Index(w, req)

			assert.Equal(t, test.code, w.Code)

			body := handlers.Message{}
			err = json.Unmarshal(w.Body.Bytes(), &body)

			assert.NoError(t, err)
			assert.Equal(t, test.message.Message, body.Message)
		})
	}
}
