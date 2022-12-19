package config

import (
	"strings"
)

type LogtailToken 	string
type ProductionMode bool

func (l *LogtailToken) UnmarshalText(text []byte) error {
	*l = LogtailToken(text)
	return nil
}

func (p *ProductionMode) UnmarshalText(text []byte) error {
	if strings.ToLower(string(text)) == "production"  {
		*p = true
	} else {
		*p = false
	}

	return nil
}