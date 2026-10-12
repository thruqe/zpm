//go:build !windows

package main

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func shellCommand(command string) *exec.Cmd {
	return exec.Command("sh", "-c", command)
}

func setDetached(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
}

func processRunning(pid int) bool {
	if pid <= 0 {
		return false
	}
	return syscall.Kill(pid, 0) == nil
}

func stopPid(pid int, force bool) {
	if pid <= 0 {
		return
	}
	signal := syscall.SIGTERM
	if force {
		signal = syscall.SIGKILL
	}
	syscall.Kill(-pid, signal)
	syscall.Kill(pid, signal)
}

func readRSS(pid int) int64 {
	file, err := os.Open(filepath.Join("/proc", strconv.Itoa(pid), "status"))
	if err != nil {
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "VmRSS:") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				value, _ := strconv.ParseInt(fields[1], 10, 64)
				return value
			}
		}
	}
	return 0
}
