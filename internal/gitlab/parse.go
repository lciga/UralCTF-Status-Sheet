package gitlab

import (
	"gopkg.in/yaml.v3"
)

type Task struct {
	Name        string `yaml:"name"`
	Author      string `yaml:"author"`
	Category    string `yaml:"category"`
	Description string `yaml:"description"`
	Attribution string `yaml:"attribution"`
	Value       int    `yaml:"value"`
	Type        string `yaml:"type"`
	Extra       struct {
		Initial int `yaml:"initial"`
		Decay   int `yaml:"decay"`
		Minimum int `yaml:"minimum"`
	} `yaml:"extra"`
	Flags          []string `yaml:"flags"`
	Tags           []string `yaml:"tags"`
	ConnectionInfo string   `yaml:"connection_info"`
	Hints          []string `yaml:"hints"`
	State          string   `yaml:"state"`
	Version        string   `yaml:"version"`
}

// Парсинг YAML задачи в структуру Task
func ParseTask(data []byte) (Task, error) {
	var t Task

	err := yaml.Unmarshal(data, &t)
	if err != nil {
		return Task{}, err
	}

	return t, nil
}
