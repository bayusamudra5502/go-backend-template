package log

import (
	"context"

	"github.com/bayusamudra5502/go-backend-template/lib/output"
)

type Log interface {
	Debug(ctx context.Context, text string)
	Info(ctx context.Context, text string)
	Warning(ctx context.Context, text string)
	Error(ctx context.Context, text string)
}

func FatalErrorLog(text string) {
	output.FormattedOutput(text, "System", "FATAL",output.ForeRed)
}
