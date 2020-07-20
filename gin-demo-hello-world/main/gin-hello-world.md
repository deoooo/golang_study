## gin hello world
gin官网：https://github.com/gin-gonic/gin

**环境**
- go 1.13
- macos 10.15.5 (19F101)

#### 启动第一个gin项目

1. 下载安装gin，执行命令行

```
go get -u github.com/gin-gonic/gin
```

2.在gopath的src目录下创建目录gin-demo-hello-world作为项目根目录.若不知道gopath目录可以运行下面命令查看，gopath就是执行结果里双引号的路径
```
go env | grep "GOPATH"
```


3.在项目根目录下创建/main/mian.go文件,输入下面代码
```
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	r.Run()
}

```

4.打开命令行,命令
```
go run 项目根目录/main/main.go
```

5.请求服务执行命令`curl 127.0.0.1:8080`看到下面的结果则为正确
```
{"message":"hello world"}
```

---
#### 写在最后
- 代码地址:https://github.com/deoooo/golang_study/tree/master/gin-demo-hello-world
- 欢迎start、watch、fork
- 任何疑问或者建议可以直接PR或者Issues
