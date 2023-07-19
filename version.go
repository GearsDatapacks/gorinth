package gorinth

import (
	"encoding/json"
	"fmt"
	"log"
)

func GetVersion(versionId string) Version {
	url := fmt.Sprintf("https://api.modrinth.com/v2/version/%s", versionId)
	result, statusCode := get(url)

	if statusCode == 404 {
		log.Fatalf("Version %q wasn't found or no authorization to see this project", versionId)
	}

	version := Version{}

	json.Unmarshal(result, &version)

	return version
}
