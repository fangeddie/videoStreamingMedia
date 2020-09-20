package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)       // 注册路由和处理函数映射
	router.POST("/user/:user_name", Login) // : 是代表变量名称

	return router
}

func main() {
	r := RegisterHandlers() // 路由集
	mh := NewMiddleWareHandler(r)

	_ = http.ListenAndServe(":8000", mh)

}
