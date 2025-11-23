package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "unknown"
	}

	version := os.Getenv("VERSION")
	if version == "" {
		version = "unknown"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Hello World - %s</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 600px; margin: 0 auto; }
        .info { background: #f0f0f0; padding: 20px; border-radius: 5px; }
        h1 { color: #333; }
        .env { color: #0066cc; font-weight: bold; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Hello World Go!</h1>
        <div class="info">
            <p><strong>Environment:</strong> <span class="env">%s</span></p>
            <p><strong>Version:</strong> %s</p>
            <p><strong>Timestamp:</strong> %s</p>
            <p><strong>Hostname:</strong> %s</p>
        </div>
    </div>
</body>
</html>
`, environment, environment, version, time.Now().Format(time.RFC3339), getHostname())
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, response)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","environment":"%s","version":"%s"}`, environment, version)
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ready","environment":"%s","version":"%s"}`, environment, version)
	})

	log.Printf("Starting Hello World server on port %s", port)
	log.Printf("Environment: %s", environment)
	log.Printf("Version: %s", version)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}
