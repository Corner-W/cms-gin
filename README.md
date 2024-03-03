# cms-gin

## 目录结构

- api gin框架的接口入口
- config config.yaml的配置读取
- core 启动gin服务
- docs swagger文档自动生成
- initialize 各种初始化系统
- middleware gin框架的中间件
- model 接口json结构
- router 路由注册
- utils 系统的基础接口
- deploy k8s部署yaml文件

## 环境要求：使用容器化环境，优先 k8s 编排管理

## 功能
- 1、开发后端 API，能够实现文件上传、文件查看、文件删除等基础文件管理功
能
- 2、实现统计 PV（Page View）数量功能

## 接口文档swagger地址

- http://127.0.0.1:8080/swagger/index.html