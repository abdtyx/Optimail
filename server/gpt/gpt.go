package gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/abdtyx/Optimail/server/config"
)

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

type GPTCore struct {
	cfg *config.ConfigGPT
}

func New(cfg *config.ConfigGPT) *GPTCore {
	return &GPTCore{cfg: cfg}
}

func (g *GPTCore) Chat(content string) (gptResponse string, err error) {
	requestData := ChatRequest{
		Model: g.cfg.Model,
		Messages: []ChatMessage{
			{Role: "user", Content: content},
		},
	}

	gptResponse = ""
	requestJson, err := json.Marshal(requestData)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", g.cfg.Apiurl, bytes.NewBuffer(requestJson))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.cfg.Apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("gpt responded:" + strconv.FormatInt(int64(resp.StatusCode), 10))
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return
	}

	if len(chatResponse.Choices) > 0 {
		return chatResponse.Choices[0].Message.Content, nil
	} else {
		err = errors.New("no response from ChatGPT")
	}
	return
}
