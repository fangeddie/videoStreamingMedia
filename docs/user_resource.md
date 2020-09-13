- 视频列表:
    - URL:/user:username/videos
    - Method:GET
    - SC:200, 400(bad request 错误请求), 500(内部错误)

- 获取一个视频: 
    - URL:/user/:username/videos/:vid-id
    - Method: GET
    - SC: 200(ok), 400, 500

- 删除一个视频
    - URL:/user/:username/videos/:vid-id
    - Method: DELETE
    - SC: 204(no content 没有返回信息), 400, 401, 403, 500