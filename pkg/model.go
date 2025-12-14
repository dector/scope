package pkg

type Report struct {
	Info   ReportInfo
	Device DeviceInfo
}

type ReportInfo struct {
	Version   string
	CreatedAt string
	Tool      ToolInfo
}

type ToolInfo struct {
	Name    string
	Version string
}

type DeviceInfo struct {
	Os     OsInfo
	Kernel KernelInfo
}

type OsInfo struct {
	Name string
	Kind string
}

type KernelInfo struct {
	Version string
	Release string
	Arch    string
}
