- 创建(注册)用户:
    - URL:/user 
    - Method:POST
    - SC:201(created 创建成功), 400(bad request 错误请求), 500(内部错误)

- 用户登陆: 
    - URL:/user/:username
    - Method: POST
    - SC: 200(ok), 400

- 获取用户基本信息
    - URL:/user/:username
    - Method: GET
    - SC: 200, 400, 401(没有验证), 403(没有权限), 500
   
- 用户注销
    - URL:/user/:username
    - Method: DELETE
    - SC: 204(no content 没有返回信息), 400, 401, 403, 500