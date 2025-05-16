package models

type Server struct {
    IP       string `json:"ip"`
    Login    string `json:"login"`
    Password string `json:"password"`
    Name     string `json:"name"`
}

type ScanResult struct {
    Name          string   `json:"name"`
    IP            string   `json:"ip"`
    Containers    []string `json:"containers"`
    Ports         []string `json:"ports"`
    GitlabRunners []string `json:"gitlab_runners"`
    Error         string   `json:"error,omitempty"`
}
