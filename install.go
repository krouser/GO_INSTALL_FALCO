package main

import (
    "fmt"
    "os/exec"
)

func main() {
    // Install flaco
    installCmd := exec.Command("apt-get", "install", "-y", "falco")
    _, err := installCmd.Output()
    if err != nil {
        fmt.Println("Error installing falco service:", err)
        return
    }
    fmt.Println("Falco service installed successfully")

    // Start falco
    startCmd := exec.Command("systemctl", "start", "falco")
    _, err = startCmd.Output()
    if err != nil {
        fmt.Println("Error starting Falco service:", err)
        return
    }
    fmt.Println("Falco service started successfully")

    // Enable falco to start on boot
    enableCmd := exec.Command("systemctl", "enable", "falco")
    _, err = enableCmd.Output()
    if err != nil {
        fmt.Println("Error enabling falco:", err)
        return
    }
    fmt.Println("falco service enabled successfully")
}
