package scanner

import (
    "bytes"
    "fmt"
    "golang.org/x/crypto/ssh"
    "go-server-scanner/models"
    "strings"
    "time"
)

func connectAndRun(server models.Server, cmd string) (string, error) {
    config := &ssh.ClientConfig{
        User:            server.Login,
        Auth:            []ssh.AuthMethod{ssh.Password(server.Password)},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Timeout:         5 * time.Second,
    }
    conn, err := ssh.Dial("tcp", server.IP+":22", config)
    if err != nil {
        return "", err
    }
    defer conn.Close()

    session, err := conn.NewSession()
    if err != nil {
        return "", err
    }
    defer session.Close()

    var out bytes.Buffer
    session.Stdout = &out
    err = session.Run(cmd)
    if err != nil {
        return "", err
    }
    return out.String(), nil
}

func ScanServer(server models.Server) models.ScanResult {
    result := models.ScanResult{
        Name: server.Name,
        IP:   server.IP,
    }

    dockerOut, err := connectAndRun(server, "docker ps --format '{{.Names}}: {{.Ports}}'")
    if err != nil {
        result.Error = fmt.Sprintf("Docker error: %s", err)
        return result
    }
    result.Containers = strings.Split(strings.TrimSpace(dockerOut), "\n")

    portOut, _ := connectAndRun(server, "ss -tuln | grep LISTEN")
    result.Ports = strings.Split(strings.TrimSpace(portOut), "\n")

    runnerOut, _ := connectAndRun(server, "gitlab-runner list || echo 'not installed'")
    result.GitlabRunners = strings.Split(strings.TrimSpace(runnerOut), "\n")

    return result
}