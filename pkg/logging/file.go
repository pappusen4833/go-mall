package logging

import (
	"fmt"
	"go-mall/pkg/global"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", global.GOMALL_CONFIG.App.RuntimeRootPath, global.GOMALL_CONFIG.App.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		global.GOMALL_CONFIG.App.LogSaveName,
		time.Now().Format(global.GOMALL_CONFIG.App.TimeFormat),
		global.GOMALL_CONFIG.App.LogFileExt,
	)
}
