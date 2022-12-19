package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/bayusamudra5502/go-backend-template/lib/log"
	"github.com/bayusamudra5502/go-backend-template/lib/output"
)

type App struct {
	Log log.Log
	handler http.Handler
}

func New(logger log.Log, handler http.Handler) (*App) {
	return &App{logger, handler}
} 

func (a *App) Start(addr string) {
	server := &http.Server{
		Addr: addr,
		Handler: a.handler,
	}
	
	sig := make(chan os.Signal, 3)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serverCtx, cancelServer := context.WithCancel(context.Background())

	go func(){
		defer cancelServer()
		defer a.Log.Info(context.Background(), "Server is down.")

		<- sig

		forceQuit, cancelForceQuit := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancelForceQuit()

		group := sync.WaitGroup{}
		group.Add(1)

		go func() {
			defer group.Done()

			a.Log.Info(context.Background(), "Gracefully shutdown....")
			<- forceQuit.Done()

			if forceQuit.Err() == context.DeadlineExceeded {
				a.Log.Error(context.Background(), "Waiting timeout, force shutdown...")
			}
		}()

		err := server.Shutdown(forceQuit)
		if err != nil {
			a.Log.Error(context.Background(), err.Error())
		}

		cancelForceQuit()
		group.Wait()
	}()

	
	a.Log.Info(context.Background(), fmt.Sprintf("Server Listen at %shttp://%s%s", output.ForeGreen,addr,output.Reset))
	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		a.Log.Error(context.Background(), "Failed to start server")
		a.Log.Error(context.Background(), err.Error())
		os.Exit(1)
	}

	<-serverCtx.Done()
}