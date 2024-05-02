package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/wekeeroad/GoWeb/global"
	"github.com/wekeeroad/GoWeb/pkg/logger"
	"github.com/wekeeroad/GoWeb/pkg/model"
	"github.com/wekeeroad/GoWeb/pkg/routers"
	"github.com/wekeeroad/GoWeb/pkg/setting"
	"github.com/wekeeroad/GoWeb/pkg/tracer"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port    string
	runMode string
	config  string

	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag is err:%v", err)
	}

	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting is err:%v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title	blog system
// @version	1.0
// @description	Golang: Produce by Golang together
// @termsOfService	https://github.com/wekeeroad/GoWeb
// @basePath	/api/v1
func main() {
	if isVersion {
		fmt.Println("========The version detail========")
		fmt.Printf("The time of build is: %v\n", buildTime)
		fmt.Printf("The version of build is: %v\n", buildVersion)
		fmt.Printf("The git commit ID is: %v\n", gitCommitID)
		return
	}

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "start port")
	flag.StringVar(&runMode, "mode", "", "start mode")
	flag.StringVar(&config, "config", "configs/", "define the path of config file")
	flag.BoolVar(&isVersion, "version", false, "if show the detail of version")
	flag.Parse()

	return nil
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ", ")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	if port != "" {
		global.ServerSetting.HttpPort = port
	}

	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
