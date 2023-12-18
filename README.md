# library

# Overview：
图书管理系统简易功能实现：

1.用户注册

2.用户登录

3.对图书的增删改查


# Structure

cfg.ini:配置信息

model:结构设计

pkg/config:配置信息初始化

pkg/database:数据库操作

pkg/logs:日志信息

pkg/routers:路由操作

pkg/service:具体功能服务实现

# Installation


# Getting started

## Required

- Mysql
- golang

## Ready

create a library database and import SQL

## Conf

modify conf/cfg.ini


## Run



```
go mod tidy
go run main.go
 ```

## Feature

- RESTful API
- Gorm
- logging
- Gin
- App configurable