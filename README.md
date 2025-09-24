# Projet Efrei Loganalyzer

LogAnalyzer est un outil en ligne de commande (CLI) en Go (1.25.1) qui permet d'analyser des fichiers de logs provenant de différents endroits (serveurs, applications) en parallèle.

## Fonctionnalités

- Analyse concurrentielle : une goroutine par log.
- Gestion d’erreurs : erreurs personnalisées (FileNotFoundError, ParseError).
- CLI avec Cobra : commandes avec flags.
- JSON : import de la configuration, export du rapport.

## Utilisation

```bash
go run main.go analyze --config config.json [--output report.json]
```
ou
```bash
go build -o loganalyzer.exe
./loganalyzer analyze --config config.json [--output report.json]
```

Options :<br>
--config, -c : chemin vers le fichier de configuration JSON (obligatoire)<br>
--output, -o : chemin vers le rapport JSON (optionnel)

## Exemple de configuration

```json
[
  { "id": "web-server-1", "path": "./test_logs/access.log", "type": "nginx-access" },
  { "id": "app-backend-2", "path": "./test_logs/errors.log", "type": "custom-app" }
]
```

## Exemple de rapport

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "./test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès."
  },
  {
    "log_id": "fichier-inexistant",
    "file_path": "./inexistant.log",
    "status": "FAILED",
    "message": "Fichier introuvable."
  }
]
```



