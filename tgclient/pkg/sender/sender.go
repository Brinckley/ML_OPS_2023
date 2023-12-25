package sender

import (
	"bytes"
	"net/http"
)

func SendPostOperation(value string) {
	_, err := http.NewRequest("POST", "http://localhost:8089/operation"+value, nil)
	if err != nil {
		return
	}
}

func SendPostType(value string) {
	_, err := http.NewRequest("POST", "http://localhost:8089/type"+value, nil)
	if err != nil {
		return
	}
}

func SendPostURL(url string) {
	var jsonStr = []byte(`{"title": "` + url + `"}`)
	_, err := http.NewRequest("POST", "http://localhost:8089/url", bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
}
