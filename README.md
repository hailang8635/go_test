## 设置go get代理
```
### 七牛
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
### 阿里云
go env -w GO111MODULE=on
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
```

## go build

```
### GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
### GOARCH：目标平台的体系架构（386、amd64、arm）
### 当CGO_ENABLED=1， 进行编译时， 会将文件中引用libc的库（比如常用的net包），以动态链接的方式生成目标文件。
### 当CGO_ENABLED=0， 进行编译时， 则会把在目标文件中未定义的符号（外部函数）一起链接到可执行文件中。
  
SET GOOS=linux
SET GOARCH=arm
SET CGO_ENABLED=0
```

## go get

```
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/sessions
go get -u github.com/go-sql-driver/mysql
go get -u github.com/ccding/go-stun
go get -u golang.org/x/tour/pic

go get -u github.com/bitly/go-simplejson
go get -u github.com/alibabacloud-go/alidns-20150109/v2/client
go get -u github.com/alibabacloud-go/darabonba-openapi/client
go get -u github.com/alibabacloud-go/tea/tea
```