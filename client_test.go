package ollama

import (
	"testing"
)

func TestClient_Prompt(t *testing.T) {
	client := NewClient("http://localhost:11434/api/generate")

	prompt, err := client.Prompt(ModelLLAMA32, "based on url value categorise the meal as part of the 6 course meal and suggest country of origin https://www.allrecipes.com/recipe/282837/meaty-baked-ziti/"+
		"return response as single line in csv format url,category,country")
	if err != nil {
		t.Fatal(err)
	}
	println(prompt)
}
