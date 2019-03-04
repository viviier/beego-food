## beego-food

一个简单的美食应用。[> beego-food](http://lab.github.com)

确保本地已安装 `go` 的使用环境
你还需要安装 `Beego` 和 `Bee` 的开发工具:

```
$ go get -u github.com/astaxie/beego
$ go get -u github.com/beego/bee
```

进入 `./models` 的各个文件下`init()`开启自动建表
```
# ./models/user.go
func init() {
  ...
  orm.RunSyncdb("default", true, true)  // 自动建表
}
```

运行 `bee run` 启动服务端

然后进入 `./static`
运行 `npm i && npm run serve` 启动客户端

#### 简介
使用 <b>beego</b>、<b>vue</b>、<b>element-ui</b> 和 <b>AdminLTE</b>构建的在线美食应用。前端使用Vue相关，后端使用了Beego

#### 技术栈
 - 前端
    * Vue
    * Vuex
    * Element-ui

 - 后端
    * Beego
    * Mysql

 - 服务器
    * Nginx

#### 实现的功能
* 用户购物车购买下单收藏功能
* 用户的登录注册功能
* 商户的管理商品,更新商品数据功能
