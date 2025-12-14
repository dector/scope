package internal

import (
	"bufio"
	"os"
	"runtime"
	"scope/pkg"
	"strings"
	"syscall"
	"time"
)

func BuildReport() (*pkg.Report, error) {
	r := pkg.Report{
		Info: pkg.ReportInfo{
			Version: "0.1.0",
			Tool: pkg.ToolInfo{
				Name:    "Scope",
				Version: "0.1.0",
			},
			CreatedAt: time.Now().String(),
		},
	}

	kernelInfo := getKernelInfo()

	r.Device = pkg.DeviceInfo{
		Os:     parseOs(),
		Kernel: kernelInfo,
	}

	return &r, nil
}

func parseOs() pkg.OsInfo {
	r := pkg.OsInfo{
		Kind: runtime.GOOS,
		Name: getOsName(),
	}
	return r
}

func getOsName() string {
	// Try to read from /etc/os-release
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return ""
	}
	defer file.Close()

	data := map[string]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		kv := strings.SplitN(line, "=", 2)

		if len(kv) != 2 {
			continue
		}

		k := kv[0]
		v := strings.Trim(kv[1], "\"")
		data[k] = v
	}

	if value, ok := data["PRETTY_NAME"]; ok {
		return value
	}

	return ""
}

func getKernelInfo() pkg.KernelInfo {
	// TODO avoid syscalls
	var utsname syscall.Utsname
	if err := syscall.Uname(&utsname); err != nil {
		return pkg.KernelInfo{
			Version: "",
			Release: "",
			Arch:    runtime.GOARCH,
		}
	}

	return pkg.KernelInfo{
		Version: int8ArrayToString(utsname.Version[:]),
		Release: int8ArrayToString(utsname.Release[:]),
		Arch:    int8ArrayToString(utsname.Machine[:]),
	}
}

func int8ArrayToString(arr []int8) string {
	b := make([]byte, 0, len(arr))
	for _, v := range arr {
		if v == 0 {
			break
		}
		b = append(b, byte(v))
	}
	return string(b)
}
