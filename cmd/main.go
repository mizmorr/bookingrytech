package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mizmorr/ingrytech/docs"
	"github.com/mizmorr/ingrytech/internal/app"
)

// @title           Swagger API
// @version         1.0
// @description     This is a book aggregation service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	if err := execute(); err != nil {
		panic(err)
	}
}

func execute() error {
	ctx := context.Background()
	app := app.New()

	if err := app.Start(ctx); err != nil {
		return err
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-stopCh

	return app.Stop(ctx)
}
