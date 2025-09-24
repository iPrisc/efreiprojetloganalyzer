package analyzer

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"loganalyzer/internal/config"
)

type AnalysisResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details,omitempty"`
}

func AnalyzeLog(log config.LogEntry) AnalysisResult {
	delay := time.Duration(50+rand.Intn(150)) * time.Millisecond
	time.Sleep(delay)

	result := AnalysisResult{
		LogID:    log.ID,
		FilePath: log.Path,
	}

	if _, err := os.Stat(log.Path); err != nil {
		if os.IsNotExist(err) {
			fileErr := &FileNotFoundError{
				FilePath: log.Path,
				Err:      err,
			}
			result.Status = "FAILED"
			result.Message = "Fichier introuvable."
			result.ErrorDetails = fileErr.Error()
			return result
		}
		result.Status = "FAILED"
		result.Message = "Erreur d'accès au fichier."
		result.ErrorDetails = err.Error()
		return result
	}

	file, err := os.Open(log.Path)
	if err != nil {
		result.Status = "FAILED"
		result.Message = "Fichier inaccessible en lecture."
		result.ErrorDetails = err.Error()
		return result
	}
	file.Close()

	result.Status = "OK"
	result.Message = "Analyse terminée avec succès."
	return result
}

func IsFileNotFoundError(err error) bool {
	var fileErr *FileNotFoundError
	return errors.As(err, &fileErr)
}

func IsParseError(err error) bool {
	var parseErr *ParseError
	return errors.As(err, &parseErr)
}
