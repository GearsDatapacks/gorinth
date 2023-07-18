package gorinth

import (
	"encoding/json"
	"fmt"
	"log"
)

func GetProject(id string) Project {
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s", id)
	result, statusCode := get(url)

	if statusCode == 404 {
		log.Fatal("Project " + id + " was not found")
	}

	project := Project{}

	json.Unmarshal(result, &project)

	return project
}
