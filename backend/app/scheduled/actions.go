package scheduled

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/app/handlers"
	"github.com/guionardo/go-gstools/scheduler"
)

var (
	Schedule     *scheduler.Scheduler
	actionsCtx   context.Context
	actionCancel context.CancelFunc
)

func Setup(ctx *handlers.MappaUserContextData) {
	logger := log.New(os.Stdout, "Scheduled Actions: ", log.LstdFlags)
	Schedule = scheduler.NewScheduler().SetLogger(logger)
	Schedule.AddEvent(get_especialidades, scheduler.RunEvery(time.Hour*24*7), scheduler.Id("get_especialidades"))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	actionsCtx, actionCancel = context.WithCancel(context.WithValue(context.Background(), "actions", ctx))

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		actionCancel()
	}()
}

func Run() {
	go Schedule.Run(actionsCtx)
}

func getActionContext() (*handlers.MappaUserContextData, error) {
	if ctx := actionsCtx.Value("actions"); ctx != nil {
		switch f := ctx.(type) {
		case *handlers.MappaUserContextData:
			return f, nil
		}
	}
	return nil, fmt.Errorf("Action context not found")
}
