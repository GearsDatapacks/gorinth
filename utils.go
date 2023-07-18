package gorinth

import (
	"io"
	"log"
	"net/http"
)

func get(url string) (body []byte, status int) {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body, response.StatusCode
}
