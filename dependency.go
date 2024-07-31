package gorinth

// Gets the version associated with a dependency
func (dep *Dependency) GetVersion() (*Version, error) {
	if dep.VersionId == nil {
		project, err := GetProject(*dep.ProjectId, "")

		if err != nil {
			return nil, err
		}

		return project.GetLatestVersion()
	}

	return GetVersion(*dep.VersionId, "")
}
