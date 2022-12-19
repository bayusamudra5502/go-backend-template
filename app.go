package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/bayusamudra5502/go-backend-template/app"
	"github.com/bayusamudra5502/go-backend-template/config"
	log "github.com/bayusamudra5502/go-backend-template/lib/log"
	"github.com/bayusamudra5502/go-backend-template/lib/output"
)

//go:embed embed/art.txt
var art string

//go:embed embed/version.txt
var version string

func headerApp() {
	fmt.Println(output.ForeGreen, art, output.ForeCyan)
	fmt.Println("[System] Application Version: " + version, output.Reset)
	fmt.Println()
}

func runApp() {
	var cfg *config.Config
	var err error


	if strings.ToLower(os.Getenv("ENV")) == "production" {
		cfg, err = config.NewEnv()
	} else {
		cfg, err = config.NewDotEnv()
	}

	if err != nil {
		log.FatalErrorLog("error while loading config")
		log.FatalErrorLog(err.Error())
		os.Exit(1)
	}
	
	app, err := app.CreateServer(cfg)
	
	if err != nil {
		log.FatalErrorLog("error while creating server")
		log.FatalErrorLog(err.Error())
		os.Exit(1)
	}

	addr := fmt.Sprintf("%s:%d", cfg.ListenAddress, cfg.ListenPort)
	app.Start(addr)

}