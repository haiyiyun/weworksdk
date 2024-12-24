package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	wework "github.com/haiyiyun/weworksdk"
	"github.com/haiyiyun/weworksdk/pkg/demo"
	"github.com/haiyiyun/weworksdk/pkg/svr"
	"github.com/haiyiyun/weworksdk/pkg/svr/logic"
	"github.com/haiyiyun/weworksdk/pkg/svr/middleware"
)

func main() {
	wwconfig := wework.WeWorkConfig{
		CorpId:              "",
		ProviderSecret:      "",
		SuiteId:             "",
		SuiteSecret:         "",
		SuiteToken:          "",
		SuiteEncodingAesKey: "",
		Dsn:                 "",
	}

	router := gin.Default()
	ww := wework.NewWeWork(wwconfig)
	logic.Migrate(wwconfig.Dsn)
	router.Use(middleware.InjectSdk(ww))

	svr.InjectRouter(router)
	demo.InjectRouter(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": ww.UserGet("", "haiyiyun")})
	})
	srv01 := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%v", "8888"),
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv01.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv01.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
}
