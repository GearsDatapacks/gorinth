package gorinth

func (dep Dependency) GetVersion() (*Version, error) {
	if dep.VersionId == "" {
		project, err := GetProject(dep.ProjectId, "")

		if err != nil {
			return nil, err
		}

		return project.GetLatestVersion()
	}

	return GetVersion(dep.VersionId, "")
}
