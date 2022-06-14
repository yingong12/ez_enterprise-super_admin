package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"super_admin/http"
	"super_admin/library"
	"super_admin/library/env"
	"super_admin/logger"
	"super_admin/providers"
	"syscall"

	"github.com/joho/godotenv"
)

//bootstrap providers,以及routines
func bootStrap() (err error) {
	//加载环境变量
	filePath := ".env"
	serverPort := 8688
	flag.StringVar(&filePath, "c", ".env", "配置文件")
	flag.IntVar(&serverPort, "p", 8688, "http端口")
	flag.Parse()
	if err = godotenv.Load(filePath); err != nil {
		return
	}
	log.Println("env loadded from file ", filePath)

	err, shutdownLogger := logger.Start()
	if err != nil {
		return
	}
	log.Println("Logger Started ")
	//加载Redis连接池
	port := env.GetIntVal("REDIS_PORT_ACCOUNT")
	poolSize := env.GetIntVal("REDIS_POOL_SIZE")
	redisConf := library.RedisConfig{
		ConnectionName: os.Getenv("SERVICE_NAME"),
		Addr:           os.Getenv("REDIS_ADDR_ACCOUNT"),
		Port:           port,
		Password:       env.GetStringVal("REDIS_PSWD_ACCOUNT"),
		DB:             0,
		PoolSize:       poolSize,
	}
	providers.RedisClient, err = library.NewRedisClient(&redisConf)
	if err != nil {
		return
	}
	// DB GORM初始化
	GormConfigs := []*library.GormConfig{
		{
			Receiver:       &providers.DBAccount,
			ConnectionName: "gorm-core",
			DBName:         env.GetStringVal("DB_ACCOUNT_RW_NAME"),
			Host:           env.GetStringVal("DB_ACCOUNT_RW_HOST"),
			Port:           env.GetStringVal("DB_ACCOUNT_RW_PORT"),
			UserName:       env.GetStringVal("DB_ACCOUNT_RW_USERNAME"),
			Password:       env.GetStringVal("DB_ACCOUNT_RW_PASSWORD"),
			MaxLifeTime:    env.GetIntVal("DB_MAX_LIFE_TIME"),
			MaxOpenConn:    env.GetIntVal("DB_MAX_OPEN_CONN"),
			MaxIdleConn:    env.GetIntVal("DB_MAX_IDLE_CONN"),
		},
	}

	for _, cfg := range GormConfigs {
		if cfg.Receiver == nil {
			return fmt.Errorf("[%s] config receiver cannot be nil", cfg.ConnectionName)
		}
		if *cfg.Receiver, err = library.NewGormDB(cfg); err != nil {
			return err
		}
		_, e := (*cfg.Receiver).DB.DB()
		if e != nil {
			return e
		}
	}
	//http client
	httpClients := []*library.HttpClientConfig{
		{
			Name:     "static_server",
			BaseURL:  `http://` + env.GetStringVal("LB_COMPANY_SERVICE"),
			Receiver: &providers.HttpClientCompanyService,
		},
	}
	for _, cfg := range httpClients {
		if cfg.Receiver == nil {
			return fmt.Errorf("config receiver cannot be nil")
		}
		*cfg.Receiver = library.NewHttpClient(cfg)
		(**cfg.Receiver).BaseURL = cfg.BaseURL
	}
	//http server
	err, shutdownHttpServer := http.Start(serverPort)
	if err != nil {
		return
	}
	log.Println("Httpserver started ")

	//wait for sys signals
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case sig := <-exitChan:
		log.Println("Doing cleaning works before shutdown...")
		shutdownLogger()
		shutdownHttpServer()
		log.Println("You abandoned me, bye bye", sig)
	}
	return
}
