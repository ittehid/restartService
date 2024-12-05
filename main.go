package main

import (
	"fmt"
	"os/exec"
	"time"
)

var serviceName = "GATE_LNK"

func stopService(name string) {
	stopCmd := exec.Command("taskkill", "/F", "/FI", fmt.Sprintf("SERVICES eq %s", name))
	stopCmd.CombinedOutput()
	fmt.Println("[INFO] Служба связи остановлена.")
}

func startService(name string) {
	startCmd := exec.Command("net", "start", name)
	startCmd.CombinedOutput()
	fmt.Println("[INFO] Служба связи запущена.")
}

func main() {
	stopService(serviceName)
	fmt.Println("Запуск службы через 5 секунд...")
	time.Sleep(5 * time.Second)
	startService(serviceName)
}
