<p align="center">
   <img alt="GgShark logo" src="https://s1.ax1x.com/2018/10/17/idhZvj.png" />
   <h3 align="center">GShark</h3>
   <p align="center">Scan for sensitive information easily and effectively.</p>
</p>

# GShark [![Go Report Card](https://goreportcard.com/badge/github.com/madneal/gshark)](https://goreportcard.com/report/github.com/madneal/gshark)   

The project is based on go with vue to build a management system for sensitive information detection. This is the total fresh version, you can refer the [old version](https://github.com/madneal/gshark/blob/gin/OLD_README.md) here.


## Features

* Support multi platform, including Gitlab, Github, Searchcode
* Flexible menu and API permission setting
* Flexible rules and filter rules
* Utilize gobuster to brute force subdomain
* Easily used management system

## Quick start

### Deployment

For the deployment of frontend, it's suggested to install nginx. Place the gshark folder under `html`, modify the `nginx.conf` to reverse proxy the backend service.

```
location /api/ {
proxy_set_header Host $http_host;
proxy_set_header  X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
proxy_set_header X-Forwarded-Proto $scheme;
rewrite ^/api/(.*)$ /$1 break;
proxy_pass http://127.0.0.1:8888; # 设置代理服务器的协议和地址
}
```
### Web service

```
./ghsark web
```

### Scan service

```
./gshark scan
```

### Development




## Config

The configuration can be set according to `app-template.ini`. You should rename it to `app.ini` to config the project.

```
HTTP_HOST = 127.0.0.1
HTTP_PORT = 8000
MAX_INDEXERS = 2
DEBUG_MODE = true
REPO_PATH = repos
MAX_Concurrency_REPOS = 5

; server酱配置口令
SCKEY =
; gobuster file path
gobuster_path =
; gobuster subdomain wordlist file path
subdomain_wordlist_file =

[database]
;support sqlite3, mysql, postgres
DB_TYPE = sqlite
HOST = 127.0.0.1
PORT = 3306
NAME = misec
USER = root
PASSWD = 
SSL_MODE = disable
;the path to store the database file of sqlite3
PATH = 
```

## Before Running

* Make sure you have installed dependencies, suggest to use go mod
* Make sure the `app.ini` in config folder, you can rename `app-template.ini` to `app.ini`
* Make sure that you have config and set database correctly, make sure create the corresponding database when using mysqp or postgresql
* Make sure that you have config corresponding tokens for Github or Gitlab

## Run

You should build the `main.go` file firstly with the command `go build main.go`.

```
USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
     web      Startup a web Service
     scan     Start to scan github leak info
     help, h  Show a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d             Debug Mode
   --host value, -H value  web listen address (default: "0.0.0.0")
   --port value, -p value  web listen port (default: 8000)
   --time value, -t value  scan interval(second) (default: 900)
   --help, -h              show help
   --version, -v           print the version
```

### Add Token

To execute `main scan`, you need to add a Github token for crawl information in github. You can generate a token in [tokens](https://github.com/settings/tokens). Most access scopes are enough. For Gitlab search, remember to add token too.

[![iR2TMt.md.png](https://s1.ax1x.com/2018/10/31/iR2TMt.md.png)](https://imgchr.com/i/iR2TMt)

## Docker support(not suggested)

Make sure rename `app-docker.ini` to `app.ini`.

### Build 

```
 docker build -t gshark-docker .      
```

### Run web

`sqlite_database_folder` is the folder for the sqlite database folder, make sure create `gshark.db` file inside the folder.
```
docker run -e OPTION=web -p 8000:8000 -v sqlite_database_folder:/data/gshark gshark-docker
```

### Run Scan 

```
docker run -e OPTION=scan -v sqlite_database_folder:/data/gshark gshark-docker
```

### Add notification

Now support notification by `server 酱`. Set the config of `SCKEY` in `app.ini` file.

## FAQ

1. Access web service 403 forbidden

Access to http://127.0.0.1/admin/login

2. Default username and password

gshark/gshark

3. `# github.com/mattn/go-sqlite3
exec: "gcc": executable file not found in %PATH%`

https://github.com/mattn/go-sqlite3/issues/435#issuecomment-314247676

4. `go get ./... connection error`

It's suggested to enable goproxy(refer this [article](https://madneal.com/post/gproxy/) for golang upgrade):

```
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
```

## Reference

* [x-patrol](https://github.com/MiSecurity/x-patrol)
* [authz](https://github.com/go-macaron/authz)
* [macaron](https://github.com/go-macaron/macaron)

## Wechat

If you would like to join wechat group, you can add my wechat `mmadneal` with the message `gshark`.

## License

[Apache License 2.0](https://github.com/madneal/gshark/blob/master/LICENSE)

## 404StarLink 2.0 - Galaxy

![](https://github.com/knownsec/404StarLink-Project/raw/master/logo.png)

GShark 是 404Team [星链计划2.0](https://github.com/knownsec/404StarLink2.0-Galaxy)中的一环，如果对 GShark 有任何疑问又或是想要找小伙伴交流，可以参考星链计划的加群方式。

- [https://github.com/knownsec/404StarLink2.0-Galaxy#community](https://github.com/knownsec/404StarLink2.0-Galaxy#community)
