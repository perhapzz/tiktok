// Code generated by hertz generator.

package main

import (
	"github.com/41197-yhkt/tiktok/gateway/internal/chat"
	"github.com/41197-yhkt/tiktok/gateway/internal/middleware"
	routers "github.com/41197-yhkt/tiktok/gateway/internal/routers"
	"github.com/41197-yhkt/tiktok/gateway/internal/rpc"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	chat.Init()
	rpc.InitRPC()
	middleware.InitJwt()
	h := server.Default()
	routers.Register(h)
	h.Spin()
}
