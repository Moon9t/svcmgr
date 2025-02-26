package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

type SvcService struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category,omitempty"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}

const configDirName = "svcmgr"

func getConfigPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to find cosmic config directory")
	}
	return filepath.Join(configDir, configDirName, "services.enc")
}

func (s *SvcService) Save() error {
	services, err := LoadServices()
	if err != nil {
		return err
	}

	// Update existing or add new
	found := false
	for i, existing := range services {
		if existing.Name == s.Name {
			services[i] = *s
			found = true
			break
		}
	}
	if !found {
		services = append(services, *s)
	}

	return saveServices(services)
}

func LoadServices() ([]SvcService, error) {
	configPath := getConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return []SvcService{}, nil
	}

	encryptedData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	decryptedData, err := DecryptConfig(encryptedData)
	if err != nil {
		return nil, err
	}

	var services []SvcService
	if err := json.Unmarshal(decryptedData, &services); err != nil {
		return nil, err
	}

	return services, nil
}

func saveServices(services []SvcService) error {
	data, err := json.MarshalIndent(services, "", "  ")
	if err != nil {
		return err
	}

	encryptedData, err := EncryptConfig(data)
	if err != nil {
		return err
	}

	configPath := getConfigPath()
	if err := os.MkdirAll(filepath.Dir(configPath), 0700); err != nil {
		return err
	}

	return os.WriteFile(configPath, encryptedData, 0600)
}
