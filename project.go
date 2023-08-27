package gorinth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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

func (project Project) CreateVersion(version Version, auth string) error {
	if version.Status == "" {
		version.Status = "listed"
	}

	if version.ProjectId == "" {
		version.ProjectId = project.Id
	}

	files := map[string]io.Reader{}

	for _, file := range version.FileParts {
		fileReader, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		files[file] = fileReader
	}

	body, status := post("https://api.modrinth.com/v2/version", version, authHeader(auth), files)
	if status == 200 {
		return nil
	}

	if status == 400 {
		return fmt.Errorf("invalid request when attempting to create version %q: %s", version.Name, string(body))
	}

	if status == 401 {
		return fmt.Errorf("no authorisation to create version %q", version.Name)
	}

	return fmt.Errorf("unexpected status code %d", status)
}

func (project Project) Modify(modified Project, auth string) error {
	overriddenValues := removeNullValues(modified)

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

func (project *Project) Validate() {
	if project.Categories == nil {
		project.Categories = []string{}
	}

	if project.InitialVersions == nil {
		project.InitialVersions = []map[string]any{}
	}

	if project.Body == "" {
		project.Body = " "
	}
}
