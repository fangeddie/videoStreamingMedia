package main

import (
	"api/dbops"
	"api/defs"
	"api/session"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body) // 从 body 取出 post 参数
	uBody := &defs.UserCredential{}  // 初始化一个结构体用来反序列化 post 参数

	if err := json.Unmarshal(res, uBody); err != nil { // 反序列化
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed) // 如果出现错误，api 直接返回
		return
	}
	// 使用用户名密码创建一个用户
	if err := dbops.AddUserCredential(uBody.Username, uBody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(uBody.Username) // 创建一个 session 认证
	su := &defs.SignedUp{Success: true, SessionId: id} // 初始化一个结构体用来序列化返回参数

	if resp, err := json.Marshal(su); err != nil { // 序列化返回参数
		sendErrorResponse(w, defs.ErrorInternalFaults) // 如果出现问题，api 直接返回
		return
	} else {
		sendNormalResponse(w, string(resp), 201) // 返回 返回参数，状态码
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	_, _ = io.WriteString(w, uname)
}
