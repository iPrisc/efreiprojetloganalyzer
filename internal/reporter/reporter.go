package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"loganalyzer/internal/analyzer"
)

func ExportResults(filePath string, results []analyzer.AnalysisResult) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible d'encoder les résultats en JSON: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire le rapport dans %s: %w", filePath, err)
	}

	return nil
}
