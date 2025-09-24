package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogEntry struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadLogsFromFile(filePath string) ([]LogEntry, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("impossible de lire le fichier %s: %w", filePath, err)
	}

	var logs []LogEntry
	if err := json.Unmarshal(data, &logs); err != nil {
		return nil, fmt.Errorf("impossible de parser le JSON dans %s: %w", filePath, err)
	}

	return logs, nil
}
