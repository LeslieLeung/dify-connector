package dify

import "github.com/go-resty/resty/v2"

type App struct {
	Type int

	client *resty.Client
}

func New(baseURL, apiKey string) *App {
	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetHeader("Authorization", "Bearer "+apiKey)
	client.SetDebug(true)
	return &App{client: client}
}

func (a *App) CompletionMessage(req CompletionMessageRequest) (*ChatCompletionMessageResponse, error) {
	resp := &ChatCompletionMessageResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		Post(apiCompletionMessages)
	return resp, err
}

func (a *App) ChatMessage(req ChatMessageRequest) (*ChatMessageResponse, error) {
	resp := &ChatMessageResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		Post(apiChatMessages)
	return resp, err
}

func (a *App) Messages(req MessagesRequest) (*MessagesResponse, error) {
	resp := &MessagesResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		Get(apiMessages)
	return resp, err
}

func (a *App) Conversations(req ConversationsRequest) (*ConversationsResponse, error) {
	resp := &ConversationsResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		Get(apiConversations)
	return resp, err
}

func (a *App) Parameters(req ParametersRequest) (*ParametersResponse, error) {
	resp := &ParametersResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		Get(apiParameters)
	return resp, err
}
