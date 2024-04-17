package dify

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"strings"
)

var (
	headerData  = []byte("data: ")
	errorPrefix = []byte(`data: {"error":`)
)

type Unmarshaler interface {
	Unmarshal(data []byte, v any) error
}

type streamable interface {
	ChatCompletionMessageResponse | ChatMessageStreamResponse | WorkflowRunStreamResponse
	GetString() string
}

type streamReader[T streamable] struct {
	isFinished bool

	reader         *bufio.Reader
	response       *http.Response
	errAccumulator ErrorAccumulator
	unmarshaler    Unmarshaler

	http.Header
}

func (stream *streamReader[T]) Recv() (response T, err error) {
	if stream.isFinished {
		err = io.EOF
		return
	}

	response, err = stream.processLines()
	return
}

//nolint:gocognit
func (stream *streamReader[T]) processLines() (T, error) {
	var (
		hasErrorPrefix bool
	)

	for {
		rawLine, readErr := stream.reader.ReadBytes('\n')
		if readErr != nil || hasErrorPrefix {
			respErr := stream.unmarshalError()
			if respErr != nil {
				return *new(T), fmt.Errorf("error, %s", respErr.Error())
			}
			return *new(T), readErr
		}

		noSpaceLine := bytes.TrimSpace(rawLine)
		if bytes.HasPrefix(noSpaceLine, errorPrefix) {
			hasErrorPrefix = true
		}
		if !bytes.HasPrefix(noSpaceLine, headerData) || hasErrorPrefix {
			if hasErrorPrefix {
				noSpaceLine = bytes.TrimPrefix(noSpaceLine, headerData)
			}
			writeErr := stream.errAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return *new(T), writeErr
			}

			continue
		}

		noPrefixLine := bytes.TrimPrefix(noSpaceLine, headerData)
		if string(noPrefixLine) == "[DONE]" {
			stream.isFinished = true
			return *new(T), io.EOF
		}

		var response T
		unmarshalErr := stream.unmarshaler.Unmarshal(noPrefixLine, &response)
		if unmarshalErr != nil {
			return *new(T), unmarshalErr
		}

		return response, nil
	}
}

func (stream *streamReader[T]) unmarshalError() (errResp *ErrorResponse) {
	errBytes := stream.errAccumulator.Bytes()
	if len(errBytes) == 0 {
		return
	}

	err := stream.unmarshaler.Unmarshal(errBytes, &errResp)
	if err != nil {
		errResp = nil
	}

	return
}

func (stream *streamReader[T]) Close() error {
	return stream.response.Body.Close()
}

func (stream *streamReader[T]) Wait() (string, error) {
	sb := strings.Builder{}
	for {
		resp, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", err
		}
		sb.WriteString(resp.GetString())
	}
	return sb.String(), nil
}

func sendRequestStream[T streamable](req *resty.Request, method, url string) (*streamReader[T], error) {
	req.SetDoNotParseResponse(true)
	req.SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "text/event-stream",
		"Cache-Control": "no-cache",
		"Connection":    "keep-alive",
	})
	resp, err := req.Execute(method, url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		// read the response body
		defer resp.RawBody().Close()
		body, readErr := io.ReadAll(resp.RawBody())
		if readErr != nil {
			return nil, fmt.Errorf("error, %w", readErr)
		}
		// try unmarshal the message to ErrorMessage
		errResp := &ErrorResponse{}
		if unmarshalErr := json.Unmarshal(body, errResp); unmarshalErr == nil {
			return nil, errResp
		}
		return nil, fmt.Errorf("error, %s", body)
	}
	return &streamReader[T]{
		reader:         bufio.NewReader(resp.RawBody()),
		response:       resp.RawResponse,
		errAccumulator: NewErrorAccumulator(),
		unmarshaler:    &JSONUnmarshaler{},
		Header:         resp.Header(),
	}, nil
}

type ErrorAccumulator interface {
	Write(p []byte) error
	Bytes() []byte
}

type errorBuffer interface {
	io.Writer
	Len() int
	Bytes() []byte
}

type DefaultErrorAccumulator struct {
	Buffer errorBuffer
}

func NewErrorAccumulator() ErrorAccumulator {
	return &DefaultErrorAccumulator{
		Buffer: &bytes.Buffer{},
	}
}

func (e *DefaultErrorAccumulator) Write(p []byte) error {
	_, err := e.Buffer.Write(p)
	if err != nil {
		return fmt.Errorf("error accumulator write error, %w", err)
	}
	return nil
}

func (e *DefaultErrorAccumulator) Bytes() (errBytes []byte) {
	if e.Buffer.Len() == 0 {
		return
	}
	errBytes = e.Buffer.Bytes()
	return
}

type JSONUnmarshaler struct{}

func (u *JSONUnmarshaler) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
