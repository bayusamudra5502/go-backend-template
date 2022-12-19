package logrus

import (
	"github.com/bayusamudra5502/go-backend-template/config"
	"github.com/bayusamudra5502/go-backend-template/lib/log/logrus/hook"
	"github.com/sirupsen/logrus"
)

type LogrusLog struct {
	log *logrus.Logger
}

type LogrusLogHook struct {
	Hook logrus.Hook
	IsProductionOnly bool
}

var logSingle *LogrusLog = nil

func New(
		isProduction config.ProductionMode,
		formatter logrus.Formatter,
		hooks []LogrusLogHook,
	) (*LogrusLog) {
	if logSingle != nil {
		return logSingle
	}
		
	log := logrus.New()

	if formatter != nil {
		log.SetFormatter(formatter)
	}

	for _, hook := range hooks {
		if hook.IsProductionOnly && bool(isProduction) || 
			 !hook.IsProductionOnly {
			log.AddHook(hook.Hook)
		}
	}

	if isProduction {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}

	logSingle = &LogrusLog{
		log: log,
	}

	return logSingle
}

func NewLogtailHooks(logtailToken config.LogtailToken) ([]LogrusLogHook) {
	return []LogrusLogHook{
		{
			Hook: hook.NewLogtailHook(logtailToken),
			IsProductionOnly: true,
		},
	}
}
