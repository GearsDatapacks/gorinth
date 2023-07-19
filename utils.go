package gorinth

import (
	"io"
	"log"
	"net/http"
)

func getFromAuth(url string, auth string) (body []byte, status int) {
	return get(url, map[string]string{"Authorization": auth})
}

func get(url string, args map[string]string) (body []byte, status int) {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range args {
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
