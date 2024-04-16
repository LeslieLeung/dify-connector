package dify

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

type App struct {
	Type int

	client *resty.Client
}

func New(baseURL, apiKey string) *App {
	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetHeader("Authorization", "Bearer "+apiKey)
	return &App{client: client}
}

// SetDebug sets the client to debug mode.
func (a *App) SetDebug() {
	a.client.SetDebug(true)
}

// CompletionMessage sends a completion message and wait for completion in blocking mode.
func (a *App) CompletionMessage(req CompletionMessageRequest) (*ChatCompletionMessageResponse, error) {
	req.ResponseMode = ResponseModeBlocking
	resp := &ChatCompletionMessageResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		Post(apiCompletionMessages)
	return resp, err
}

type CompletionMessageStream struct {
	*streamReader[ChatCompletionMessageResponse]
}

// CompletionMessageStreaming sends a completion message and wait for completion in streaming mode.
func (a *App) CompletionMessageStreaming(req CompletionMessageRequest) (*CompletionMessageStream, error) {
	req.ResponseMode = ResponseModeStreaming
	c := a.client.R().SetBody(req)
	resp, err := sendRequestStream[ChatCompletionMessageResponse](c, http.MethodPost, apiCompletionMessages)
	if err != nil {
		return nil, err
	}
	return &CompletionMessageStream{streamReader: resp}, nil
}

// StopCompletionMessage stops the completion.
func (a *App) StopCompletionMessage(taskID string) error {
	_, err := a.client.R().
		SetPathParam("task_id", taskID).
		Post(apiCompletionMessagesStop)
	return err
}

// ChatMessage sends a chat message in blocking mode.
func (a *App) ChatMessage(req ChatMessageRequest) (*ChatMessageResponse, error) {
	req.ResponseMode = ResponseModeBlocking
	resp := &ChatMessageResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		Post(apiChatMessages)
	return resp, err
}

type ChatMessageStream struct {
	*streamReader[ChatMessageStreamResponse]
}

func (a *App) ChatMessageStream(req ChatMessageRequest) (*ChatMessageStream, error) {
	req.ResponseMode = ResponseModeStreaming
	c := a.client.R().SetBody(req)
	resp, err := sendRequestStream[ChatMessageStreamResponse](c, http.MethodPost, apiChatMessages)
	if err != nil {
		return nil, err
	}
	return &ChatMessageStream{streamReader: resp}, nil
}

func (a *App) StopChatMessage(taskID string) error {
	_, err := a.client.R().
		SetPathParam("task_id", taskID).
		Post(apiChatMessagesStop)
	return err
}

func (a *App) WorkflowRun(req WorkflowRunRequest) (*WorkflowRunResponse, error) {
	resp := &WorkflowRunResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		Post(apiWorkflowRun)
	return resp, err
}

type WorkflowRunStream struct {
	*streamReader[WorkflowRunStreamResponse]
}

func (a *App) WorkflowRunStream(req WorkflowRunRequest) (*WorkflowRunStream, error) {
	req.ResponseMode = ResponseModeStreaming
	c := a.client.R().SetBody(req)
	resp, err := sendRequestStream[WorkflowRunStreamResponse](c, http.MethodPost, apiWorkflowRun)
	if err != nil {
		return nil, err
	}
	return &WorkflowRunStream{streamReader: resp}, nil
}

func (a *App) StopWorkflowRun(taskID string) error {
	_, err := a.client.R().
		SetPathParam("task_id", taskID).
		Post(apiWorkflowRunStop)
	return err
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

func (a *App) Feedback(req MessagesFeedbacksRequest) (*MessagesFeedbacksResponse, error) {
	resp := &MessagesFeedbacksResponse{}
	_, err := a.client.R().
		SetBody(req).
		SetResult(resp).
		SetPathParam("message_id", req.MessageID).
		Post(apiMessagesFeedbacks)
	return resp, err
}

func (a *App) Suggested(messageID, user string) (*SuggestedResponse, error) {
	resp := &SuggestedResponse{}
	_, err := a.client.R().
		SetResult(resp).
		SetPathParam("message_id", messageID).
		SetQueryParam("user", user).
		Get(apiSuggested)
	return resp, err
}
