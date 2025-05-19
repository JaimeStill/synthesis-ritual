package utils

import (
	"fmt"
	"runtime"
	"time"
)

func GetCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

func GetSystemInfo() string {
	return fmt.Sprintf("OS: %s, Architecture: %s, CPUs: %d",
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
	)
}
