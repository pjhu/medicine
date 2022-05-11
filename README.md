## 1. 技术栈

- GO 1.8.1
- gin-gonic/gin 【web框架】
- viper 【配置文件处理】
- golang-migrate 【数据库migrate】
- gorm 【ORM框架】
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

### 安装依赖包
```
go mod tidy
go mod download
```

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
migrate -source "file://scripts/db/migrations" -database "mysql://root:123@tcp(localhost:3306)/usercenter" up
```

### Gopls报错配置
```
https://commandnotfound.cn/go/5/577/VSCode-%E5%AE%89%E8%A3%85-Go-%E6%8F%92%E4%BB%B6%E5%A4%B1%E8%B4%A5%E8%A7%A3%E5%86%B3%E6%96%B9%E6%A1%88
```

## test

### create mock file
```
go install github.com/golang/mock/mockgen@v1.6.0

/Users/jinhupeng/.asdf/installs/golang/1.18.1/packages/bin/mockgen -source=domain/repository.go -destination=mock/orderrepository_mock.go -package=mock

/Users/jinhupeng/.asdf/installs/golang/1.18.1/packages/bin/mockgen -source=application/applicationservice.go -destination=mock/application_mock.go -package=mock
```

### run all test
```
go test ./... -v
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
- goroutine
- 垃圾回收
- 分布式事务
- 性能调优
