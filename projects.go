package gorinth

import (
	"encoding/json"
	"fmt"
	"log"
)

// GetProject returns a Project Model of the project with a matching ID or slug with the one provided.
func GetProject(id_or_slug string, auth string) Project {
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s", id_or_slug)
	result, statusCode := get(url)

	if statusCode == 404 {
		log.Fatalf("Project %q wasn't found or no authorization to see this project", id_or_slug)
	}

	project := Project{}

	json.Unmarshal(result, &project)

	return project
}

// Gets the most recently created version of a project
func (project Project) GetLatestVersion() Version {
	versions := project.GetVersions()

	fmt.Println(versions)

	if len(versions) == 0 {
		log.Fatalf("Project %q has no versions.", project.Title)
	}

	return versions[0]
}

// Gets all versions of a project
func (project Project) GetVersions() []Version {
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s/version", project.Slug)
	result, statusCode := get(url)

	if statusCode == 404 {
		log.Fatalf("Project %q wasn't found or no authorization to see this project", project.Slug)
	}

	response := []Version{}
	json.Unmarshal(result, &response)

	return response
}
