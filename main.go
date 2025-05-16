package main

import (
    "encoding/json"
    "fmt"
    "go-server-scanner/models"
    "go-server-scanner/scanner"
    "io/ioutil"
    "net/http"
)

func main() {
    http.HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
        data, err := ioutil.ReadFile("servers.json")
        if err != nil {
            http.Error(w, "Can't read servers.json", 500)
            return
        }

        var servers []models.Server
        err = json.Unmarshal(data, &servers)
        if err != nil {
            http.Error(w, "Invalid JSON", 500)
            return
        }

        var results []models.ScanResult
        for _, srv := range servers {
            results = append(results, scanner.ScanServer(srv))
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(results)
    })

    fmt.Println("Server listening on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
