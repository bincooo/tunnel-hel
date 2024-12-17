### tunnel-hel - GO Proxy Collector

### A simple proxy pool written in Golang

#### [中文文档](README.md)

Features
------

* Support Redis or MySQL as storage
* Automatically grab free proxy IP on the Internet
* Regularly check IP availability and recheck failed IP
* Get the list of available proxy IP through API
* Tunnel proxy service: it supports opening multiple ports for monitoring, multi protocol, multi-level forwarding, user name and password verification, and user flow restriction

## Usage：

#### install

    git clone https://github.com/tongsq/tunnel-hel.git
    cd tunnel-hel

#### start service
1、Start script to collect free proxy IP

    go run ./cmd -S=job -C=conf.yaml
2、Start API service

    go run ./cmd -S=api -C=conf.yaml
2、Start Tunnel service

    go run ./cmd -S=tunnel -C=conf.yaml
#### Get available IP through the API

    curl 127.0.0.1:8090/all?city=上海&duration=100

#### Start multiple services

    go run ./cmd -S=job -S=api -C=conf.yaml
#### One touch start
    go run ./cmd -S=all -C=conf.yaml

## yaml configure
### Set storage media
A、 Set redis as storage media

* Modify the configuration file conf.yaml as follows
```
    dao: redis
    redis:
      MaxIdle: 10
      MaxActive: 20
      Network: tcp
      Address: 127.0.0.1:6379
      Password: your password
```
B、Set mysql as storage media

* Modify the configuration file conf.yaml as follows
```
    dao: database
    database:
      Dialect: mysql
      Args: user:password@(127.0.0.1:3306)/dbname?charset=utf8&loc=Local
```
* Create table
```
    CREATE TABLE `proxy` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `host` varchar(255) NOT NULL DEFAULT '',
      `port` int(11) NOT NULL DEFAULT '0',
      `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '0:无效，1：有效',
      `create_time` int(11) NOT NULL DEFAULT '0',
      `update_time` int(11) NOT NULL DEFAULT '0',
      `active_time` int(11) NOT NULL DEFAULT '0',
      `country` varchar(100) NOT NULL DEFAULT '',
      `region` varchar(100) NOT NULL DEFAULT '',
      `city` varchar(100) NOT NULL DEFAULT '',
      `isp` varchar(255) NOT NULL DEFAULT '',
      `check_count` int(11) NOT NULL DEFAULT '10',
      `source` varchar(50) NOT NULL DEFAULT '',
      `proto` varchar(20) NOT NULL DEFAULT 'http',
      `user` varchar(50) NOT NULL DEFAULT '',
      `password` varchar(50) NOT NULL DEFAULT '',
      `ping` int(11) NOT NULL DEFAULT '0',
      PRIMARY KEY (`id`) USING BTREE,
      UNIQUE KEY `IDX_HOST_PORT_PROTO` (`host`,`port`, `proto`) USING BTREE,
      KEY `IDX_STATUS` (`status`) USING BTREE,
      KEY `IDX_ACTIVE_TIME` (`active_time`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT
```
# TODO list
- [x] Supports proxy collection and comparison of other protocols such as socket5
- [x] Support configuring log classification
- [x] Support tunnel agent services
