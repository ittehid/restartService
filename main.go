package main

import (
	"fmt"
	"os/exec"
	"time"
)

var serviceName = "GATE_LNK"

// stopService завершает службу по имени
func stopService(name string) {
	stopCmd := exec.Command("taskkill", "/F", "/FI", fmt.Sprintf("SERVICES eq %s", name))
	stopCmd.CombinedOutput()
	fmt.Printf("[INFO] Служба %s остановлена.\n", name)
}

// startService запускаем службу
func startService(name string) {
	startCmd := exec.Command("cmd", "net", "start", name)
	startCmd.CombinedOutput()
	fmt.Printf("[INFO] Служба %s запущена.", name)
}

func main() {
	stopService(serviceName)
	time.Sleep(5 * time.Second)
	startService(serviceName)
}
