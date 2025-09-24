package analyzer

import "fmt"

type FileNotFoundError struct {
	FilePath string
	Err      error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s (%v)", e.FilePath, e.Err)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

type ParseError struct {
	FilePath string
	Line     int
	Err      error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("erreur de parsing dans %s (ligne %d): %v", e.FilePath, e.Line, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}
