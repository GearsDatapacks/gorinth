package gorinth

import "log"

func (dep Dependency) GetVersion() Version {
	if dep.VersionId == "" {
		project, err := GetProject(dep.ProjectId, "")
		
		if err != nil {
			log.Fatal(err)
		}
		
		return project.GetLatestVersion()
	}

	return GetVersion(dep.VersionId, "")
}
