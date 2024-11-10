package main

import (
	"breed/config"
	constant "breed/constant"
	"breed/database"
	"breed/middlewares"
	"breed/routers"
	utils "breed/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func InitLoadEnv() {
	// Load config from env
	appConfig, errLoadConfig := config.LoadConfig("./config/env")
	if errLoadConfig != nil {
		log.Fatal("cannot load config:", errLoadConfig)
	}
	config.AppConfig = appConfig

	fmt.Println("----------------------------------")
	fmt.Println("MODE:", appConfig.Mode)
	fmt.Println("PORT:", appConfig.Port)
	fmt.Println("----------------------------------")
}

func main() {
	app := gin.Default()

	if err := app.SetTrustedProxies(nil); err != nil {
		utils.Logger.Error(err.Error())
	}

	app.Use(middlewares.CORS())

	InitLoadEnv()

	if err := database.ConnectDB(); err != nil {
		panic(err)
	} else {
		fmt.Println("[INFO] Initialize Connect database successfully...")
	}

	utils.InitializeLogger()

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorMessage(constant.ERR_ENDPOINT_NOT_FOUND, nil))
	})
	app.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, utils.ErrorMessage(constant.ERR_ENDPOINT_NOT_FOUND, nil))
	})

	router := app.Group("")

	routers.InitRouter(router)

	errRunApp := app.Run(":8000")
	if errRunApp != nil {
		panic(errRunApp)

	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:         ":" + config.AppConfig.Port,
		Handler:      app,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if errListen := srv.ListenAndServe(); errListen != nil && errListen != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", errListen)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down ....")

	timeOutctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if errShutdown := srv.Shutdown(timeOutctx); errShutdown != nil {
		utils.Logger.Error(errShutdown.Error())
	}

}
