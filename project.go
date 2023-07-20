package gorinth

import (
	"encoding/json"
	"fmt"
	"log"
)

// GetProject returns a Project Model of the project with a matching ID or slug with the one provided.
func GetProject(id_or_slug string, auth string) (Project, error) {
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s", id_or_slug)
	result, statusCode := get(url, authHeader(auth))

	if statusCode == 404 {
		return Project{}, fmt.Errorf("Project %q wasn't found or no authorization to see this project", id_or_slug)
	}

	project := Project{}

	json.Unmarshal(result, &project)

	project.auth = auth

	return project, nil
}

// Gets all versions of a project
func (project Project) GetVersions() []Version {
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s/version", project.Slug)
	result, statusCode := get(url, authHeader(project.auth))

	if statusCode == 404 {
		log.Fatalf("Project %q wasn't found or no authorization to see this project", project.Slug)
	}

	response := []Version{}
	json.Unmarshal(result, &response)

	return response
}

// Gets the most recently created version of a project
func (project Project) GetLatestVersion() Version {
	versions := project.GetVersions()

	if len(versions) == 0 {
		log.Fatalf("Project %q has no versions.", project.Title)
	}

	return versions[0]
}

// Get the version of the given project whose semver string matches the given string
func (project Project) GetSpecificVersion(versionNumber string) Version {
	versions := project.GetVersions()

	for _, version := range versions {
		if version.VersionNumber == versionNumber {
			return version
		}
	}

	log.Fatalf("Cannot find version %s of project %q", versionNumber, project.Title)
	return Version{}
}

func (project Project) Modify(modified Project, auth string) error {
	overriddenValues := removeZeroValues(modified)

	url := "https://api.modrinth.com/v2/project/" + project.Id
	body, status := patch(url, overriddenValues, authHeader(auth))

	if status == 204 {
		return nil
	}

	if status == 404 {
		return fmt.Errorf("Project %q wasn't found or no authorization to see this project", project.Slug)
	}

	if status == 401 {
		responseSchema := struct {
			Error       string
			Description string
		}{}

		json.Unmarshal(body, &responseSchema)

		return fmt.Errorf("%s: %s", responseSchema.Error, responseSchema.Description)
	}

	return fmt.Errorf("unexpected response, status code %d", status)
}
