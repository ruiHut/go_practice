package main

import (
	"os/exec"
)

func main() {
	cmdo := exec.Command("echo", "-n", "my first command come from golong,")
	cmdo.Start()
}
