package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Header struct {
	Key   string
	Value string
}

type BasicAuth struct {
	Username string
	Password string
}

func NewRequest(method string, url string, data *bytes.Buffer, header []Header, auth BasicAuth) (map[string]interface{}, error) {
	var (
		request  *http.Request
		response *http.Response
		client   http.Client
		err      error
		body     map[string]interface{}
	)

	switch method {
	case "GET":
		request, err = http.NewRequest(method, url, nil)
	case "POST":
		request, err = http.NewRequest(method, url, data)
	case "PATCH":
		request, err = http.NewRequest(method, url, data)
	case "PUT":
		request, err = http.NewRequest(method, url, data)
	case "DELETE":
		request, err = http.NewRequest(method, url, nil)
	default:
		return body, errors.New("method not found")
	}

	if err != nil {
		return body, err
	}

	if (BasicAuth{}) != auth {
		request.SetBasicAuth(auth.Username, auth.Password)
	}

	for _, v := range header {
		request.Header.Add(v.Key, v.Value)
	}

	response, err = client.Do(request)

	if err != nil {
		Log(LogTypeError, fmt.Sprintf("[HELPER REST_API] Error client do: %v", err.Error()))
		return body, err
	}

	basicBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		Log(LogTypeError, "[HELPER REST_API] Error while parsing data")
	}

	json.Unmarshal(basicBody, &body)

	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 204 {
		Log(LogTypeError, fmt.Sprintf("[HELPER REST_API] %s Error response: %d | %v", url, response.StatusCode, string(basicBody)))
		return body, fmt.Errorf("%v", string(basicBody))
	}

	return body, nil

}
