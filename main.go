package main

import (
    "fmt"
    "flag"
    "os/exec"
    "strings"

    "github.com/AlecAivazis/survey/v2"
)

var (
    Version    string
)

func main() {
    version := flag.Bool("version", false, "print version information")
    flag.Parse()

    if *version {
        fmt.Printf("Version: %s\n", Version)
        fmt.Printf("Maintainer: https://github.com/achelak\n")
        return
    }
    contextsBytes, err := exec.Command("kubectl", "config", "get-contexts", "-o", "name").Output()
    if err != nil {
        fmt.Println("Failed to get Kubernetes contexts:", err)
        return
    }
    contexts := strings.Split(strings.TrimSpace(string(contextsBytes)), "\n")

    var selectedContext string
    prompt := &survey.Select{
        Message: "Choose Kubernetes context:",
        Options: contexts,
    }
    if err := survey.AskOne(prompt, &selectedContext); err != nil {
        fmt.Println("Failed to get selected context:", err)
        return
    }

    if _, err := exec.Command("kubectl", "config", "use-context", selectedContext).Output(); err != nil {
        fmt.Println("Failed to switch to selected context:", err)
        return
    }
    fmt.Println("Switched to context:", selectedContext)
}
