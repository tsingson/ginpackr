# ginpackr

bind [packr2](https://github.com/gobuffalo/packr/tree/master/v2) to [gin-gonic/gin](https://github.com/gin-gonic/gin), to embed static file inside golang webserv

## usage

### 0. before 
assume code struct in /User/tsingson/go/src/github.com/tsingson/ginpackr-example/
```
./main.go
./public
./public/css
./public/js
./public/index.html
....

```


### 1. get packr2 
```
cd /User/tsingson/go/src/github.com/tsingson/ginpackr-example/
go mod init
go get -u github.com/gobuffalo/packr/v2/...
go get -u github.com/gobuffalo/packr/v2/packr2

```

### 2. code struct
```
./main.go
./public
./public/css
./public/js
./public/index.html
....

```

### 3. use ginpackr in gin 
in main.go  ( in /User/tsingson/go/src/github.com/tsingson/ginpackr-example/)
```
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/tsingson/ginpackr"
	...
)

func main() {

	box := packr.NewBox( "./public") // static HTML file here 
	r := gin.Default()
	r.Use(ginpackr.PackrServe("/", box)) // use packr2 box ...

	r.Static("/css", "public/css")
	r.Static("/js", "public/js")
	r.Run()
}

```

### 4. run it in develop moe
```
cd /User/tsingson/go/src/github.com/tsingson/ginpackr-example/
go run ./main.go

```

### 5. build it 
```
cd /User/tsingson/go/src/github.com/tsingson/ginpackr-example/
pack2 install .
```


##  inspired by gin-static

## MIT license
see LICENSE

