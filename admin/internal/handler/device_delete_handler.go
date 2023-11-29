package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"iot-platform/admin/internal/logic"
	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"
)

func DeviceDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeviceDeleteLogic(r.Context(), svcCtx)
		resp, err := l.DeviceDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
