### 使用
`go get github.com/wuyan94zl/redigo`

```go
package main

import (
	"github.com/wuyan94zl/redigo""wuyan94zl/gowebpkg/redis"
)
func init() {
	r := redis.Config{Host: "127.0.0.1:6379", Password: "123456", MaxActive: 200, MaxIdle: 50, IdleTimeout: 300}
	redis.ConRedis(r)
}
func main(){
    // redis 函数一致 如get,del,set,setnx,setex,mset,sadd,lpush等
	redis.Get("key")
}
```