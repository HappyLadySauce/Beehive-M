// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package gateway

import (
	"net/http"

	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/svc"
	"github.com/gorilla/websocket"
)

func WsEntryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 约定从 query 参数拿 userID/deviceID，后续可替换为鉴权结果
		userID := r.URL.Query().Get("userID")
		deviceID := r.URL.Query().Get("deviceID")
		if userID == "" || deviceID == "" {
			http.Error(w, "missing userID/deviceID", http.StatusBadRequest)
			return
		}

		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		c := svcCtx.Hub.Register(wsConn, userID, deviceID)
	// if err := svcCtx.Hub.RegisterSession(r.Context(), userID, deviceID, c.ConnID); err != nil {
	// 	c.Close()
	// 	return
	// }
		c.Close()

		// TODO: 这里应该进入读写循环/消息分发；目前仅完成注册与分布式互踢
	}
}
