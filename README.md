## 1. 技术栈

- GO 1.6
- gin-gonic/gin
- viper
- golang-migrate
- xorm
- postgresql
- pkg/error

## 2. 初次运行

### 安装migration CLI
```
brew install golang-migrate
```

### Start postgres
```$xslt
cd docker
docker build -t postgres-test:11.2 . 
docker-compose up -d
```

### Create database
```$xslt
创建数据库的脚本位于application/src/main/resources/db/init.sql

psql -h localhost -p 15432 -U postgres -W -c "create database test;" 
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
migrate -source file://application/main/resources/db/migrations -database postgres://localhost:15432/test?sslmode=disable up
```

## 需要解决
- exception[x]
- restfulapi[x]
- snowflaks[x]
- validator[x]
- authentication[x]
- authorization
- test
- 在k8s中运行
- 外部请求
- 统一错误处理[x]
- 外部服务调用接口化
- 分布式事务
- goroutine
- 垃圾回收
- model变为私有的，没发捕捉名字映射错误
