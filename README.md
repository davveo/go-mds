# go-mds

A reliable message service system developed with golang

# 消息发送流程
1. 主动方发送预消息(消息状态为"待确认")
2. 主动方根据(1)返回的结果执行响应的业务操作
3. 发送业务处理结果
4. 消息中间件根据业务结果, 更新("待发送")/删除对应的消息
5. 消息中间件监听并接收待发送状态的消息


# 消息服务系统
1. 存储预发送消息
2. 确认并发送消息
3. 查询状态确认超时的消息
4. 确认消息已被成功消息
5. 查询消费确认超时的消息
6. 删除消息
7. 消息状态确认子系统(crontab)
8. 消息恢复子系统

# Api
1. 消息发送
    POST http://127.0.0.1:3000/api/v1/message/create
 
    ```json
    {
        "messageBody": "{'a': 'this is a test'}",
        "messageQueue": "test",
        "messageId": "43296482634873",
        "extra": ""
    }
    ```