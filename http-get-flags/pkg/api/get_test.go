package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	GetResponseOutput  *http.Response
	PostResponseOutput *http.Response
}

func (m MockClient) Get(url string) (resp *http.Response, err error) {
	return m.GetResponseOutput, nil
}

func (m MockClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return m.PostResponseOutput, nil
}

func TestDoGetRequest(t *testing.T) {
	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}

	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Errorf("marshal error: %s", err)
	}

	apiInstance := api{ //initiate the api Struct and then pass option to client
		Options: Options{},
		//you need to create a fake request
		//in order to check the logic of  our program
		//without the need of the server
		Client: MockClient{
			GetResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}

	response, err := apiInstance.DoGetRequest("http://localhost/words")
	if err != nil {
		t.Errorf("DoGetRequest error: %s", err)
	}
	if response == nil {
		t.Fatal("response is empty")
	}
	if response.GetResponse() != `Words: a, b` {
		t.Errorf("Got wrong output: %s", response.GetResponse())
	}

}
