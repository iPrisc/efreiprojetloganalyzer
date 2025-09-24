# Projet Efrei Loganalyzer

Mini-CRM en Go (1.25.1) pour gérer une liste de contacts.

## Fonctionnalités

- Analyse concurrentielle : une goroutine par log.
- Gestion d’erreurs : erreurs personnalisées (FileNotFoundError, ParseError).
- CLI avec Cobra : commandes avec flags.
- JSON : import de la configuration, export du rapport.

## Utilisation

```bash
go run main.go analyze --config config.json [--output report.json]
```

Options :
--config, -c : chemin vers le fichier de configuration JSON (obligatoire)
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



