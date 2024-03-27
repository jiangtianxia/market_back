package router

import (
	"fmt"
	"market_back/common/conf"
	"market_back/handle"
	mlog "market_back/logger"
	"market_back/middlewares"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(cnf *conf.ServerConf) {
	// http
	server := gin.Default()

	if cnf.IsReleaseMode {
		// 设置成发布模式
		gin.SetMode(gin.ReleaseMode)
	}

	// 跨域中间件
	if cnf.OpenCORS {
		server.Use(cors.New(cors.Config{
			AllowMethods: []string{"GET", "POST", "PUT", "HEAD", "PATCH", "OPTIONS", "DELETE"},
			AllowHeaders: []string{"Origin", "Content-Length", "Authorization", "Content-Type", "X-TOKEN",
				"Cookie", "Tus-Extension", "Tus-Resumable", "Tus-Version",
				"Upload-Length", "Upload-Metadata", "Upload-Offset",
				"Access-Control-Allow-Origin", "X-HTTP-Method-Override"},
			AllowCredentials: true,
			AllowAllOrigins:  false,
			AllowOriginFunc: func(origin string) bool {
				for _, once := range cnf.AllowOrigins {
					if once == origin || once == "" {
						return true
					}

					if strings.Contains(origin, once) {
						return true
					}
				}
				return false
			},
			MaxAge: 12 * time.Hour,
		}))
	}

	// 静态资源
	server.StaticFile("/market/doc/v1/swagger.json", "apidocs/swagger.json")
	server.StaticFS("/upload", http.Dir("upload"))

	// 路由配置
	v1 := server.Group("/market/api/v1", middlewares.CurrentLimit())
	{
		// hello
		v1.GET("/hello", handle.Hello)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		if err := server.Run(fmt.Sprintf(":%d", cnf.HTTPPort)); err != nil {
			mlog.Errorf("[Market] http Server start failed Port: %d, err: %v \n", cnf.HTTPPort, err)
			stopChan <- syscall.SIGTERM
		}
		mlog.Infof("[Market] http Server started Port: %d \n", cnf.HTTPPort)
	}()

	<-stopChan
}
