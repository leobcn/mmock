package translate

import (
	"bytes"
	"net/http"
	"testing"
)

func TestBuildRequestDefinitionFromHTTPMethod(t *testing.T) {
	body := bytes.NewBufferString("name=jmartin")
	req, err := http.NewRequest("POST", "http://test.test", body)
	if err != nil {
		t.Fatal("Impossible create request")
	}

	translator := HTTPTranslator{}
	definition := translator.BuildRequestDefinitionFromHTTP(req)

	if definition.Method != "POST" {
		t.Error("Invalid method")
	}
}
