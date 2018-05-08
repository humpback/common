package models

import "github.com/docker/libcompose/project"

var (
	ProjectActions = map[string]interface{}{
		"start":   true,
		"stop":    true,
		"restart": true,
		"kill":    true,
		"pause":   true,
		"unpause": true,
	}
)

// CreateProject - define compose create project struct
type CreateProject struct {
	Name        string `json:"Name"`
	ComposeData string `json:"ComposeData"`
	PackageFile string `json:"PackageFile"`
}

// OperateProject - define compose operate project struct
type OperateProject struct {
	Name   string `json:"Name"`
	Action string `json:"Action"`
}

// ProjectConfig - define compose project info struct
type ProjectConfig struct {
	Name        string          `json:"Name"`
	HashCode    string          `json:"HashCode"`
	ComposeData string          `json:"ComposeData"`
	PackageFile string          `json:"PackageFile"`
	Timestamp   int64           `json:"Timestamp"`
	Containers  project.InfoSet `json:"Containers"`
}

// ProjectStatus - define compose project status struct
type ProjectStatus struct {
	Name   string      `json:"Name"`
	Status interface{} `json:"Status"`
}

// ProjectUploadStatus - define compose project package file upload struct
type ProjectUploadStatus struct {
	Name        string      `json:"Name"`
	PackageFile string      `json:"PackageFile"`
	Status      interface{} `json:"Status"`
}
