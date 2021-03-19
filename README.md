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
- kong/ratelimit 【网关限流】
- tracking
- test

## 2. 本地运行

### 安装数据库migration CLI
```
brew install golang-migrate
```

### build postgres
```$xslt
cd devops/postgresql
docker build -t postgres-test:11.2 . 
```

### 启动postgresql, redis
```
cd devops
docker-compose -f docker-compose-local.yml up -d
```

### Create database
```$xslt
创建数据库的脚本位于devops/init.sql

psql -U postgres 
```

### DB Migretion
https://github.com/golang-migrate/migrate

```
create table 

cd #{project path}
migrate create -ext sql -dir application/main/resources/db/migrations  create_order_table
```

```
migrate

cd #{project path}
migrate -source file://application/main/resources/db/migrations -database postgres://localhost:15432/order?sslmode=disable up
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

```
创建数据库的脚本位于devops/init.sql
psql -U postgres 
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
