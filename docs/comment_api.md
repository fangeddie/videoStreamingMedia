- 展示视频所有评论:
    - URL:/user:vid-id/comments
    - Method:GET
    - SC:200, 400(bad request 错误请求), 500(内部错误)

- 发布一个视频评论: 
    - URL:/user:vid-id/comments
    - Method: POST
    - SC: 201, 400, 500

- 删除一个视频评论
    - URL:/user:vid-id/comment/:comment-id
    - Method: DELETE
    - SC: 204(no content 没有返回信息), 400, 401, 403, 500