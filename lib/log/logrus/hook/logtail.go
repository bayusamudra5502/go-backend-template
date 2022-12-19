package hook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bayusamudra5502/go-backend-template/config"
	"github.com/bayusamudra5502/go-backend-template/lib/output"
	"github.com/sirupsen/logrus"
)

type LogtailHook struct	{
	token string
}

type requestPayload struct {
	Timestamp string 			 `json:"dt"`
	Level			logrus.Level `json:"level"`
	Message		string			 `json:"message"`
}

func NewLogtailHook(token config.LogtailToken) (*LogtailHook) {
	return &LogtailHook{
		token: string(token),
	}
}

func (LogtailHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel, 
		logrus.FatalLevel, 
		logrus.ErrorLevel, 
		logrus.WarnLevel, 
		logrus.InfoLevel,
	}
}

func (l LogtailHook) Fire(entry *logrus.Entry) error {
	payload := requestPayload{
		Level: entry.Level,
		Timestamp: entry.Time.Format(time.RFC3339),
		Message: entry.Message,
	}

	go func() {
		if err := sendToLogtail(payload, l.token); err != nil {
			output.FormattedOutput("Failed to send log to server: " + err.Error(), "LOG", "error", output.ForeRed)
		}
	}()

	return nil
}

var client = &http.Client{
	Transport: &http.Transport{
		IdleConnTimeout: 5 * time.Second,
	},
}

func sendToLogtail(payload requestPayload, token string) error {
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return err
	}
	
	reader := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "https://in.logtail.com", reader)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("response code is not 200 OK : got %d", res.StatusCode)
}