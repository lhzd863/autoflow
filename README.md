# autoflow

## 简介

autoflow分布调度系统服务端,支持批量调度之间无干扰，元数据相互对立，使用元数据能够快速恢复一个相同运行环境。系统采用服务端和客户端相分离方式，只有同时使用服务端和客户端才能完成整套调度系统搭建。服务端元数据存储采用bbolt无需安装，配置相关路径可以直接使用，客户端采用vue-element-admin 为原型基础上开发完成。

## 依赖
```
-github.com/satori/go.uuid
-google.golang.org/grpc
-github.com/emicklei/go-restful
-go.etcd.io/bbolt
-jwt
-workpool

```
## 名词
```
-镜像: 带有存储作业配置，依赖，参数，触发等相关配置的bbolt存储文件，建立目录创建ID和打上标记形成对象
-实例: 镜像文件被具体到某个批量，存储文件从镜像文件copy，不会影响镜像配置文件
-作业: 实现例具体执行操作对象

```
## 功能
```
-API
 -api文档采用swagger
 -实例作业的添加，删除，修改
 -镜像文件库添加，删除，修改
 -实例创建，启动，停止
 
-Mgr
 -检查Pending状态作业
 -检查Go状态作业
 
-Worker
 -执行作业
 -作业执行LOG信息返回
 -参数替换和环境变量设置顺序:系统->实例->作业
 
```

## Online Demo

[在线 Demo](https://122.51.161.53:12300)

## 项目声明：
本软件只供学习交流使用，勿作为商业用途,如有任何问题联系lhzd863(lhzd863@126.com)
