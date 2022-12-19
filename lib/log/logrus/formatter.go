package logrus

import (
	"fmt"

	"github.com/bayusamudra5502/go-backend-template/lib/output"
	"github.com/sirupsen/logrus"
)


type LogrusFormatter struct {}

var colorMap = map[logrus.Level]output.Color{
	logrus.TraceLevel: output.ForeWhite,
	logrus.DebugLevel: output.ForeWhite,
	logrus.InfoLevel: output.ForeGreen,
	logrus.WarnLevel: output.ForeYellow,
	logrus.ErrorLevel: output.ForeRed,
	logrus.PanicLevel: output.ForeRed,
}

func (f* LogrusFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s%s%s: [App] %s\n", 
		entry.Time.Format("2006-01-02 15:04:05 MST"),
		colorMap[entry.Level],
		entry.Level,
		output.Reset,
		entry.Message,
	)), nil
}

func NewFormatter() (*LogrusFormatter) {
	return &LogrusFormatter{}
}