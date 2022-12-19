package logrus

import (
	"context"
)

func (l *LogrusLog) Debug(ctx context.Context, text string) {
	l.log.Debug(text)
}

func (l *LogrusLog) Info(ctx context.Context, text string) {
	l.log.Info(text)
}

func (l *LogrusLog) Warning(ctx context.Context, text string) {
	l.log.Warn(text)
}

func (l *LogrusLog) Error(ctx context.Context, text string) {
	l.log.Error(text)
}
