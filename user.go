package gorinth

import (
	"bytes"
	"encoding/json"
	"io"
)

func GetUserFromAuth(auth string) (*User, error) {
	body, status, err := get("https://api.modrinth.com/v2/user", authHeader(auth))
	if err != nil {
		return nil, err
	}

	if status == 401 {
		return nil, makeError("Invalid authorisation token given")
	}

	if status == 200 {
		user := User{}

		err := json.Unmarshal(body, &user)
		if err != nil {
			return nil, makeError(err.Error())
		}
		user.auth = auth

		return &user, nil
	}

	return nil, makeError("Unexpected response status %d", status)
}

func (user User) CreateProject(project Project) error {
	project.Validate()
	overriddenValues, err := removeNullValues(project)
	if err != nil {
		return err
	}

	overriddenValues["license_id"] = project.License.Id
	overriddenValues["is_draft"] = true

	parts := map[string]io.Reader{}

	if project.Icon != nil {
		parts["icon"] = bytes.NewBuffer(project.Icon)
	}

	body, status, err := post(
		"https://api.modrinth.com/v2/project",
		overriddenValues,
		authHeader(user.auth),
		parts,
	)
	if err != nil {
		return err
	}

	if status == 200 {
		return nil
	}

	if status == 400 {
		return makeError("invalid request when attempting to create project %q: %s", project.Title, string(body))
	}

	if status == 401 {
		return makeError("no authorisation to create project %q", project.Title)
	}

	return makeError("unexpected status code %d", status)
}
