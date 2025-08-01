package gitlab

import (
	"log"

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
func ParseTask(data []byte) Task {
	var task Task
	err := yaml.Unmarshal(data, &task)
	if err != nil {
		log.Printf("Ошибка при разборе YAML: %v", err)
	} else {
		log.Printf("DEBUG: разобранный task: %+v", task)
	}
	return task
}
