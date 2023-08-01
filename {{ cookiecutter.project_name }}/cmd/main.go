package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
	"{{ cookiecutter.project_name }}/setting"
	"{{ cookiecutter.project_name }}/task"
)

type app struct {
	ctx     context.Context
	cancel  func()
	httpSrv *http.Server
	log     *logrus.Logger
	task    *task.Task
}

func newApp(httpSrv *http.Server, log *logrus.Logger, task *task.Task) *app {
	ctx, cancel := context.WithCancel(context.Background())
	return &app{
		ctx:     ctx,
		cancel:  cancel,
		httpSrv: httpSrv,
		log:     log,
		task:    task,
	}
}

func (a *app) run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}

	eg.Go(func() error {
		<-ctx.Done() // wait for stop signal
		stopCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		a.log.Info("shutdown server")
		return a.httpSrv.Shutdown(stopCtx)
	})

	eg.Go(func() error {
		a.log.Info("server listen server")
		err := a.httpSrv.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	})

	eg.Go(func() error {
		a.log.Info("register tasks")
		return a.task.Run(ctx)
	})

	wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case <-c:
			a.cancel()
			return nil
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil

}

//func newRouter() *gin.Engine {
//	router := gin.Default()
//	router.GET("/", func(c *gin.Context) {
//		//time.Sleep(5 * time.Second)
//		c.String(http.StatusOK, "Welcome Gin Server")
//	})
//	return router
//}

func newHttpServer(router *gin.Engine, log *logrus.Logger) *http.Server {
	log.Infof("binding port: %v", setting.ServerSetting.HttpPort)
	log.Infof("running mod: %v", setting.ServerSetting.RunMode)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler: router,
	}
}

func main() {
	err := setting.Load()
	if err != nil {
		panic(err)
	}

	var log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if setting.ServerSetting.RunMode == "debug" {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	log.SetReportCaller(true)

	app, cleanup, err := wireApp(log)

	if err != nil {
		panic(err)
	}

	defer cleanup()

	if err := app.run(); err != nil {
		panic(err)
	}

	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//	//time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "Welcome Gin Server")
	//})
	//
	//srv := &http.Server{
	//	Addr:    ":8080",
	//	Handler: router,
	//}
	//
	//go func() {
	//	// 服务连接
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//
	//// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//log.Println("Shutdown Server ...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	////ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//log.Println("Server exiting")
}
