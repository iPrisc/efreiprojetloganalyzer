package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Un outil CLI pour analyser des fichiers de logs",
	Long:  `LogAnalyzer est un outil CLI en Go pour analyser des fichiers de logs (serveurs, applications) en parallèle, d'extraire les informations clés, et de gérer les erreurs.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		os.Exit(1)
	}
}

func init() {
}
