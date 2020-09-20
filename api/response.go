package main

import (
	"api/defs"
	"encoding/json"
	"io"
	"net/http"
)

// 错误响应
func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC) // 写入状态码

	resStr, _ := json.Marshal(&errResp.Error) // 序列化
	_, _ = io.WriteString(w, string(resStr))  // 写入响应中
}

// 正常响应
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	_, _ = io.WriteString(w, resp)
}
