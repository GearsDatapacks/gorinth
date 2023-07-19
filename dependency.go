package gorinth

func (dep Dependency) GetVersion() Version {
	if dep.VersionId == "" {
		return GetProject(dep.ProjectId, "").GetLatestVersion()
	}

	return GetVersion(dep.VersionId)
}
