package gorinth

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Returns a Project Model of the project with a matching ID or slug with the one provided.
func GetProject(id_or_slug string, auth string) (*Project, error) {
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s", id_or_slug)
	result, statusCode, err := get(url, authHeader(auth))
	if err != nil {
		return nil, err
	}

	if statusCode == 404 {
		return nil, makeError("Project %q wasn't found or no authorization to see this project", id_or_slug)
	}

	project := Project{}

	json.Unmarshal(result, &project)

	project.auth = auth

	return &project, nil
}

// Returns all versions of the project
func (project *Project) GetVersions() ([]Version, error) {
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s/version", project.Slug)
	result, statusCode, err := get(url, authHeader(project.auth))
	if err != nil {
		return nil, err
	}

	if statusCode == 404 {
		return nil, makeError("Project %q wasn't found or no authorization to see this project", project.Slug)
	}

	response := []Version{}
	json.Unmarshal(result, &response)

	return response, nil
}

// Returns the most recently created version of the project
func (project *Project) GetLatestVersion() (*Version, error) {
	versions, err := project.GetVersions()
	if err != nil {
		return nil, err
	}

	if len(versions) == 0 {
		return nil, makeError("Project %q has no versions", project.Title)
	}

	return &versions[0], nil
}

// Returns the version of the project whose semver string matches the given string
func (project *Project) GetSpecificVersion(versionNumber string) (*Version, error) {
	versions, err := project.GetVersions()
	if err != nil {
		return nil, err
	}

	for _, version := range versions {
		if version.VersionNumber == versionNumber {
			return &version, nil
		}
	}

	return nil, makeError("Cannot find version %s of project %q", versionNumber, project.Title)
}

// Creates a version associated with the project
func (project *Project) CreateVersion(version Version, auth string) error {
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
			return err
		}

		files[file] = fileReader
	}

	body, status, err := post("https://api.modrinth.com/v2/version", version, authHeader(auth), files)
	if err != nil {
		return err
	}

	if status == 200 {
		return nil
	}

	if status == 400 {
		return makeError("invalid request when attempting to create version %q: %s", version.Name, string(body))
	}

	if status == 401 {
		return makeError("no authorisation to create version %q", version.Name)
	}

	return makeError("unexpected status code %d", status)
}

// Changes the icon of the project
func (project *Project) ChangeIcon(icon []byte, auth string) error {
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s/icon?ext=%s", project.Id, "png")
	body, status, err := patch(url, icon, authHeader(auth))
	if err != nil {
		return err
	}

	if status == 204 {
		return nil
	}

	if status == 400 {
		return makeError("invalid request when attempting to modify project icon: %s", string(body))
	}

	return makeError("unexpected response, status code %d, error %s", status, string(body))
}

// Modifies the project on Modrinth, updating all non-zero fields
func (project *Project) Modify(modified Project, auth string) error {
	overriddenValues, err := removeZeroValues(modified)
	if err != nil {
		return err
	}

	url := "https://api.modrinth.com/v2/project/" + project.Id
	body, status, err := patch(url, overriddenValues, authHeader(auth))
	if err != nil {
		return err
	}

	if status == 204 {
		if modified.Icon == nil {
			return nil
		}

		return project.ChangeIcon(modified.Icon, auth)
	}

	if status == 404 {
		return makeError("Project %q wasn't found or no authorization to see this project", project.Slug)
	}

	if status == 401 {
		responseSchema := struct {
			Error       string
			Description string
		}{}

		json.Unmarshal(body, &responseSchema)

		return makeError("%s: %s", responseSchema.Error, responseSchema.Description)
	}

	if status == 400 {
		return makeError("invalid request when attempting to modify project icon: %s", string(body))
	}

	return makeError("unexpected response, status code %d, error %s", status, string(body))
}

func validSlug(slug string) bool {
	match, err := regexp.Match("^[\\w!@$()`.+,\"\\-']{3,64}$", []byte(slug))

	return match && err == nil
}

func toTitle(slug string) string {
	spaced := strings.Replace(slug, "-", " ", -1)
	title := ""

	for i, char := range spaced {
		if i == 0 {
			title += strings.ToTitle(string(char))
		} else if spaced[i-1] == ' ' {
			title += strings.ToTitle(string(char))
		} else {
			title += string(char)
		}
	}

	return title
}

// Validates a project, fixing any straightforward problems,
// and returning an error for any problems which can't be resolved
func (project *Project) Validate() error {
	if !validSlug(project.Slug) {
		return makeError("Invalid project slug %q", project.Slug)
	}

	if len(project.Body) > 3 {
		return makeError("Invalid project body with fewer than 3 characters")
	}

	if project.Title == "" {
		title := toTitle(project.Slug)
		logWarning("Invalid project title %q, automatically generated title %q", project.Title, title)
		project.Title = title
	}

	if project.ClientSide != "required" && project.ClientSide != "optional" && project.ClientSide != "unsupported" {
		project.ClientSide = "optional"
	}

	if project.ServerSide != "required" && project.ServerSide != "optional" && project.ServerSide != "unsupported" {
		project.ServerSide = "optional"
	}

	if project.Status == "" {
		project.Status = "unknown"
	}

	if project.ProjectType == "" {
		project.ProjectType = "mod"
	}

	if project.Categories == nil {
		project.Categories = []string{}
	}

	return nil
}
