package logging

import (
	"fmt"
	"time"

	"github.com/IainMcl/go-behind-the-scenes/internal/settings"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", settings.AppSetting.RuntimeRootPath, settings.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		settings.AppSetting.LogSaveName,
		time.Now().Format(settings.AppSetting.TimeFormat),
		settings.AppSetting.LogFileExt,
	)
}
