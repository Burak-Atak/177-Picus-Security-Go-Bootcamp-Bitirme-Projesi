package main

import (
	"fmt"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/api"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/graceful"
	"net/http"
	"time"

	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	api.RegisterHandlers(r)

	err := r.Run(":8090")
	registerMiddlewares(r)
	if err != nil {
		fmt.Println(err)
	}

	srv := &http.Server{
		Addr:    ":8090",
		Handler: r,
	}

	graceful.ShutdownGin(srv, time.Second*15)
}

func registerMiddlewares(r *gin.Engine) {
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	r.Use(middleware.LatencyLogger())

}
