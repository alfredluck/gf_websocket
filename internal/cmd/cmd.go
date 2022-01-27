package cmd

import (
	"context"
	"encoding/json"
	"gf_websocket/internal/controller"
	"gf_websocket/ws"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			//websocket
			go ws.Manager.Start()
			s.BindHandler("/ws", ws.WsHandler)
			//提供过接口发送消息到客户端
			s.BindHandler("/send", func(r *ghttp.Request) {
				uid := gconv.String(r.Get("uid"))
				touid := gconv.String(r.Get("to_uid"))
				msg := gconv.String(r.Get("msg"))
				clientId := uid + "_" + touid
				conn := ws.Manager.Clients[clientId]
				if conn != nil {
					jsonMessage, _ := json.Marshal(&ws.Message{Content: msg})
					conn.Socket.WriteMessage(websocket.TextMessage, jsonMessage)
					r.Response.Write("发送成功")
				} else {
					r.Response.Write("webSocket未找到")
				}
			})
			//websocket

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(controller.Hello)
			})
			s.Run()
			return nil
		},
	}
)
