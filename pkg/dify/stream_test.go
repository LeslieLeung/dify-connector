package dify

import (
	"github.com/google/uuid"
	"os"
)

func ExampleStreamReader_Wait() {
	client := New(os.Getenv("DIFY_URL"), os.Getenv("DIFY_API_KEY"))
	client.SetDebug()

	text := "Hello, how are you?"
	req := CompletionMessageRequest{
		Inputs: map[string]interface{}{
			"query":    text,
			"language": "English",
		},
		User: uuid.New().String(),
	}

	resp, err := client.CompletionMessageStreaming(req)
	if err != nil {
		panic(err)
	}
	out, err := resp.Wait()
	if err != nil {
		panic(err)
	}
	print(out)
	// Output:
}
