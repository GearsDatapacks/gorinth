package gorinth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func authHeader(auth string) map[string]string {
	return map[string]string{"Authorization": auth}
}

func get(url string, headers map[string]string) (body []byte, status int) {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseBody, response.StatusCode
}

func patch(url string, payload any, headers map[string]string) (body []byte, status int) {
	requestSchema, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(requestSchema))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseBody, response.StatusCode
}

func toMap[T any](object T) map[string]any {
	str, err := json.Marshal(object)

	if err != nil {
		log.Fatal(err)
	}

	result := map[string]any{}

	err = json.Unmarshal(str, &result)

	if err != nil {
		log.Fatal(err)
	}
	return result
}

func removeZeroValues[T Project](object T) map[string]any {
	zeroValueStruct := T{}

	zeroValues := toMap(zeroValueStruct)
	values := toMap(object)

	for key, value := range values {
		if value == nil {
			delete(values, key)
			continue
		}

		switch value.(type) {
		case map[string]any:
			continue
		}

		if value == zeroValues[key] {
			delete(values, key)
		}
	}

	return values
}
