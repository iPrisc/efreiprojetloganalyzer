package cmd

import (
	"fmt"
	"sync"

	"loganalyzer/internal/analyzer"
	"loganalyzer/internal/config"
	"loganalyzer/internal/reporter"

	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse des fichiers de logs en parallèle",
	Long:  `La commande analyze traite une liste de logs depuis un fichier JSON. Chaque log est analysé en parallèle et les résultats sont exportés au format JSON.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			fmt.Println("Erreur: le chemin du fichier de configuration (--config) est obligatoire.")
			return
		}

		logs, err := config.LoadLogsFromFile(configPath)
		if err != nil {
			fmt.Printf("Erreur lors du chargement de la configuration: %v\n", err)
			return
		}

		if len(logs) == 0 {
			fmt.Println("Aucun log à analyser trouvé dans le fichier de configuration.")
			return
		}

		fmt.Printf("Analyse de %d fichiers de logs...\n", len(logs))

		var wg sync.WaitGroup
		resultsChan := make(chan analyzer.AnalysisResult, len(logs))

		wg.Add(len(logs))
		for _, logEntry := range logs {
			go func(log config.LogEntry) {
				defer wg.Done()
				result := analyzer.AnalyzeLog(log)
				resultsChan <- result
			}(logEntry)
		}

		wg.Wait()
		close(resultsChan)

		var results []analyzer.AnalysisResult
		for result := range resultsChan {
			results = append(results, result)

			if result.Status == "OK" {
				fmt.Printf("OK %s (%s) : %s - %s\n",
					result.LogID, result.FilePath, result.Status, result.Message)
			} else {
				fmt.Printf("KO %s (%s) : %s - %s\n",
					result.LogID, result.FilePath, result.Status, result.Message)
				if result.ErrorDetails != "" {
					fmt.Printf("Détails: %s\n", result.ErrorDetails)
				}
			}
		}

		// Export si demandé
		if outputPath != "" {
			err := reporter.ExportResults(outputPath, results)
			if err != nil {
				fmt.Printf("Erreur lors de l'exportation: %v\n", err)
			} else {
				fmt.Printf("Rapport exporté vers %s\n", outputPath)
			}
		}

		fmt.Printf("\nAnalyse terminée: %d fichiers traités\n", len(results))
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin vers le fichier JSON de configuration")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin vers le fichier JSON de sortie (optionnel)")

	analyzeCmd.MarkFlagRequired("config")
}
