# GoFrame WebSocket Channel 通信，支付多客户端，可以通过http接口向对应客户端发送消息 

## webSocket 连接
``ws://127.0.0.1/ws``

## http接口发送消息

- Form-Data 参数:

|参数名|类型| 说明   |
|:---- |:----|:-----|
|uid |string| 组id  |
|to_uid |string| 子id  |
|msg |string| 消息内容 |


- 接口地址
``http://localhost:8000/send``
