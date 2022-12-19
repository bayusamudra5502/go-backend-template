package output

import (
	"fmt"
	"time"
)

func FormattedOutput(text string, process string , logType string, color Color) {
	fmt.Printf("%s %s%s%s: [%s] %s\n", 
	time.Now().Format("2006-01-02 15:04:05 MST"),
	color,
	logType,
	Reset,
	process,
	text,
)
}