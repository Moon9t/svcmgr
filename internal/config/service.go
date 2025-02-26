package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Service represents a service configuration.
type Service struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category,omitempty"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Extra    string `json:"extra"` // Add the Extra field
}

func (s Service) Save() any {
	panic("unimplemented")
}

// LoadService loads a service configuration from a JSON file.
func LoadService(filePath string) (*Service, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var service Service
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&service); err != nil {
		return nil, fmt.Errorf("could not decode JSON: %v", err)
	}

	return &service, nil
}
