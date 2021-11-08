## 1. 技术栈

- GO 1.6
- gin-gonic/gin 【web框架】
- viper 【配置文件处理】
- golang-migrate 【数据库migrate】
- xorm 【ORM框架】
- postgresql 【数据库驱动】
- snowflaks【Id generator】
- pkg/error 【错误处理】
- logrus 【日志】
- oxy 【bff转发】
- resty 【访问外部服务】
- hystrix-go 【服务熔断降级】
- cosmtrek/air【热加载】
- tracking
- test
- kong/ratelimit 【网关限流】

## 2. 本地运行

### 安装数据库migration CLI
```
brew install golang-migrate
```

### 启动mysql, redis
```
cd devops
docker-compose up -d
```

### DB Migretion
https://github.com/golang-migrate/migrate

创建table文件，添加sql内容
```
cd #{project path}
migrate create -ext sql -dir db/migrations  create_order_table
```

执行迁移命令
```
cd #{project path}
migrate -source "file://db/migrations" -database "mysql://root:123@tcp(localhost:3306)/ordercenter" up
```

### Gopls报错配置
open settings.json add, you can search Gopls in setting, and then edit, it will auto add flowing config
```
"gopls": {
        "build.experimentalWorkspaceModule": true,
},
```

## 3. docker启动
### 使用docker compose 启动
```
cd ordercenter
docker build -t ordercenter:1.0 .
```

```
cd usercenter
docker build -t usercenter:1.0 .
```

```
cd tocbff
docker build -t tocbff:1.0 .
```

```
docker-compose -f devops/docker-compose.yml up -d
```

## 需要解决
- validator[x]
- 统一错误处理[x]
- authorization
- 并发
- goroutine
- 垃圾回收
- 分布式事务
- 性能调优
- model变为私有的，没发捕捉名字映射错误
