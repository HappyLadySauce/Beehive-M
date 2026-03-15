// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package gateway

import (
	"net/http"

	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/logic/gateway"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func WsEntryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := gateway.NewWsEntryLogic(r.Context(), svcCtx)
		err := l.WsEntry()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, struct{}{})
		}
	}
}
