package dify

import (
	"errors"
	"github.com/google/uuid"
	"io"
)

func ExampleApp_CompletionMessageStreaming() {
	client := New("", "")
	client.SetDebug()

	text := "Hello, how are you?"

	req := CompletionMessageRequest{
		Inputs: map[string]interface{}{
			"query": text,
		},
		User: uuid.New().String(),
	}

	resp, err := client.CompletionMessageStreaming(req)
	if err != nil {
		panic(err)
	}
	for {
		r, err := resp.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}
		if r.Answer != "" {
			print(r.Answer)
		}
	}
	// Output:
}
