//go:build windows

package main

import (
	"os/exec"
	"strconv"
	"syscall"
	"unsafe"
)

func shellCommand(command string) *exec.Cmd {
	return exec.Command("cmd", "/C", command)
}

func setDetached(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP}
}

func processRunning(pid int) bool {
	if pid <= 0 {
		return false
	}
	const da = syscall.STANDARD_RIGHTS_READ | syscall.PROCESS_QUERY_INFORMATION | syscall.SYNCHRONIZE
	h, err := syscall.OpenProcess(da, false, uint32(pid))
	if err != nil {
		return false
	}
	defer syscall.CloseHandle(h)

	s, err := syscall.WaitForSingleObject(h, 0)
	if err != nil {
		return false
	}
	return s == syscall.WAIT_TIMEOUT
}

func stopPid(pid int, force bool) {
	if pid <= 0 {
		return
	}
	args := []string{"/PID", strconv.Itoa(pid), "/T"}
	if force {
		args = append(args, "/F")
	}
	exec.Command("taskkill", args...).Run()
}

type processMemoryCounters struct {
	cb                         uint32
	pageFaultCount             uint32
	peakWorkingSetSize         uintptr
	workingSetSize             uintptr
	quotaPeakPagedPoolUsage    uintptr
	quotaPagedPoolUsage        uintptr
	quotaPeakNonPagedPoolUsage uintptr
	quotaNonPagedPoolUsage     uintptr
	pagefileUsage              uintptr
	peakPagefileUsage          uintptr
}

var (
	modpsapi                 = syscall.NewLazyDLL("psapi.dll")
	procGetProcessMemoryInfo = modpsapi.NewProc("GetProcessMemoryInfo")
)

const processVMRead = 0x0010

func readRSS(pid int) int64 {
	if pid <= 0 {
		return 0
	}
	const da = syscall.PROCESS_QUERY_INFORMATION | processVMRead
	h, err := syscall.OpenProcess(da, false, uint32(pid))
	if err != nil {
		return 0
	}
	defer syscall.CloseHandle(h)

	var pmc processMemoryCounters
	pmc.cb = uint32(unsafe.Sizeof(pmc))
	r1, _, _ := procGetProcessMemoryInfo.Call(
		uintptr(h),
		uintptr(unsafe.Pointer(&pmc)),
		uintptr(pmc.cb),
	)
	if r1 == 0 {
		return 0
	}
	return int64(pmc.workingSetSize / 1024)
}
