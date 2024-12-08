package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

var ModelLLAMA32 = "llama3.2"

type Client struct {
	baseURL string
	client  *http.Client
}

type request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type response struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
}

func NewClient(baseURL string) *Client {
	return &Client{client: &http.Client{Timeout: time.Second * 10}, baseURL: baseURL}
}
func (c Client) Prompt(model, prompt string) (string, error) {
	data, err := json.Marshal(request{model, prompt})
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(data))
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	reader := bufio.NewReader(resp.Body)
	result := ""
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return "", err
		}
		if err == io.EOF {
			break
		}
		tmp := response{}
		err = json.Unmarshal(line, &tmp)
		if err != nil {
			return "", err
		}
		result += tmp.Response
	}
	return result, nil
}
