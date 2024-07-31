package gorinth

import (
	"encoding/json"
	"fmt"
)

// Gets the version associated with the given ID
func GetVersion(versionId string, auth string) (*Version, error) {
	url := fmt.Sprintf("https://api.modrinth.com/v2/version/%s", versionId)
	result, statusCode, err := get(url, authHeader(auth))
	if err != nil {
		return nil, err
	}

	if statusCode == 404 {
		return nil, makeError("Version %q wasn't found or no authorization to see this project", versionId)
	}

	version := Version{}

	json.Unmarshal(result, &version)

	return &version, nil
}
