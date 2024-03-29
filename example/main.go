// @Description  TODO
// @Author  	 jiangyang  
// @Created  	 2020/11/9 5:54 下午
package main

import (
	"encoding/json"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/comeonjy/util/elastic"
	"github.com/comeonjy/util/email"
	"github.com/comeonjy/util/log"

	"github.com/comeonjy/util/mysql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/comeonjy/util/config"
	"github.com/comeonjy/util/ctx"
	"github.com/comeonjy/util/errno"
	"github.com/comeonjy/util/jwt"
	"github.com/comeonjy/util/middlewares"
	"github.com/comeonjy/util/rbac"
	"github.com/comeonjy/util/server"
)

func init() {
	config.LoadConfig()
	email.Init(config.GetConfig().Email)
	elastic.Init(config.GetConfig().Elastic)
	log.Init(config.GetConfig().Log)
	mysql.Init(config.GetConfig().Mysql)
	rbac.Init(config.GetConfig().Rbac)
}

func main() {
	r := gin.New()
	r.Use(middlewares.Recovery())
	r.Use(middlewares.LoggerToLogrus())
	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	r.GET("token", ctx.Handle(token))

	auth := r.Group("")
	auth.Use(middlewares.JwtAuth())
	auth.GET("/ping", ctx.Handle(ping))
	auth.Use(middlewares.Rbac(rbac.Check)).GET("/auth", ctx.Handle(ping))

	rbac.Register(r, "/rbac")

	server.Server(r, viper.GetInt("http_port"))

}

func token(ctx *ctx.Context) {
	bus := jwt.Business{
		UID:  1,
		Role: 2,
	}
	tokenResp, err := jwt.CreateToken(bus, 24*time.Hour)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(tokenResp)
}

func ping(ctx *ctx.Context) {
	bus, exists := ctx.Get("business")
	if !exists {
		ctx.Fail(errno.BusNotFound)
		return
	}
	b := jwt.Business{}
	marshal, err := json.Marshal(bus)
	if err != nil {
		ctx.Fail(err)
		return
	}
	if err := json.Unmarshal(marshal, &b); err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(b)
}
