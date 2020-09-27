package main

import (
	"io"
	"net/http"
)

// 错误响应
func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)                // 写入状态码
	_, _ = io.WriteString(w, errMsg) // 写入响应中
}
