package ollama

import (
	"testing"
)

func TestClient_Prompt(t *testing.T) {
	client := NewClient("http://localhost:11434/api/generate", nil)

	prompt, err := client.Prompt(ModelLLAMA32, "test prompt")
	if err != nil {
		t.Fatal(err)
	}
	println(prompt)
}
