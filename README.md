## Best Practices
```shell
# Use `export` instead of godotenv
# source .env
export `cat .env` # tested in mac and unix
export `cat .env | grep -v "#"` && go run cmd/main.go
```


## Request
```go
// 获取 Path Parameters
name := c.Param("name")

// 获取参数 GET
name := c.QueryParam("name")

// 获取参数 GET POST PUT PATCH
name := c.FormValue("name")

// 获取 body 任意请求类型
var r io.Reader = c.Request().Body
b, _ := io.ReadAll(r)
fmt.Println(string(b))
```


## Cookie
```go
// 读取 cookie
c.Cookies()

// 写入 cookie
cookie := &http.Cookie{Name: "userId", Value: "123456"}
http.SetCookie(w, cookie)
```


## Defined
```golang
// 分隔符
const (
    nul uint8 = 0x00 // 空字符
    lf  uint8 = 0x0A // 换行
    cr  uint8 = 0x0D // 回车键
    fs  uint8 = 0x1C // 文件分隔符
    gs  uint8 = 0x1D // 组分隔符
    rs  uint8 = 0x1E // 记录分隔符
    us  uint8 = 0x1F // 单元分隔符
)
```


## Uint32/16/8 to bytes
```golang
// uint64 to []byte
// e.g. 1
var num uint64 = 258
b := make([]byte, 8)
binary.BigEndian.PutUint64(b, num)

// e.g. 2
var num uint16 = 2
bf := bytes.NewBuffer(nil)
binary.Write(bf, binary.BigEndian, num)
b := bf.Bytes() // [0 2]
```


## Bytes to uint32/16/8
```golang
bs := []byte{0x00, 0x00, 0x01, 0x02}
num := binary.BigEndian.Uint32(bs)

bs := []byte{0x01, 0x02}
num := binary.BigEndian.Uint16(bs)

bs := []byte{0x01}
num := bs[0]
```


## Copy
```golang
// bytes
bs[0] = 0x1F
copy(bs, bs1)
bytes.Split(bs, []byte{0x1F})
```


```golang
// 守护进程
var w sync.WaitGroup
w.Add(2)
go func () {
    // do something
    w.Done()
}
go func () {
    // do something
    w.Done()
}
w.Wait()
```


```golang
// map是无序的, 数组是有序的
var foo = [3]int{1, 2, 3}
var bar = map[string]int64{
    "a": 1,
    "b": 2,
    "c": 3,
}

for index, item := range foo {
    fmt.Println(index, item)
}

for index, item := range bar {
    fmt.Println(index, item)
}

// 输出：
// 0 1
// 1 2
// 2 3

// c 3
// a 1
// b 2

// 解决方案 - 引入其他排序数组
import "sort"

var m map[int]string
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}
```


```go
// 切片特性 - 引用传递
func main() {
	alice := []int{0, 1, 2, 3, 4}
	fmt.Println(alice) // [0 1 2 3 4]

	bob := nothingDid(alice)
	bob[1] = 9

	fmt.Println(alice) // [0 9 2 3 4]
}

func nothingDid(data []int) []int {
	return data
}
```


```go
// 数组特性 - 值传递
func main() {

	alice := [5]int{0, 1, 2, 3, 4}
	fmt.Println(alice) // [0 1 2 3 4]

	bob := nothingDid(alice)
	bob[1] = 9

	fmt.Println(alice) // [0 1 2 3 4]
}

func nothingDid(data [5]int) [5]int {
	return data
}
```


## Libraries
Visit [here](https://github.com/avelino/awesome-go) for more about the libraries.  

| 包名 | 简介 |
| --- | --- |
| [github.com/gorilla/mux](https://github.com/gorilla/mux) | web 框架 |
| [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) | web 框架 |
| [github.com/gorilla/sessions](https://github.com/gorilla/sessions) | session |
| [github.com/joho/godotenv](https://github.com/joho/godotenv) | 配置, env 环境变量 |
| [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) | 配置, yaml |
| [github.com/spf13/viper](https://github.com/spf13/viper) | 配置, 支持多种格式 |
| [gorm.io/gorm](https://gorm.io) | MySQL ORM |
| [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) | 数据库驱动, MySQL |
| [github.com/lib/pq](https://github.com/lib/pq) | 数据库驱动, Postgres |
| [github.com/gocql/gocql](https://github.com/gocql/gocql) | 数据库驱动, Cassandra |
| [go.mongodb.org/mongo-driver](https://go.mongodb.org/mongo-driver) | 数据库驱动, mongoDB |
| [github.com/go-redis/redis](https://github.com/go-redis/redis) | 数据库驱动, redis |
| [github.com/go-resty/resty/](https://github.com/go-resty/resty/) | http 请求客户端 |
| [github.com/valyala/fastjson](https://github.com/valyala/fastjson) | json 解析 |
| [github.com/tidwall/gjson](https://github.com/tidwall/gjson) | json 解析, 注意: []强制转为字符串后为[] |
| [github.com/buger/jsonparser](https://github.com/buger/jsonparser) | json 解析 |
| [github.com/Sirupsen/logrus](https://github.com/Sirupsen/logrus) | 日志 |
| [github.com/uber-go/zap](https://github.com/uber-go/zap) | 日志 |
| [github.com/bwmarrin/snowflake](https://github.com/bwmarrin/snowflake) | ID生成 |
| [github.com/google/uuid](https://github.com/google/uuid) | UUID生成 |
| [github.com/silenceper/pool](https://github.com/silenceper/pool) | 线程池 |
| [github.com/apache/thrift](https://github.com/apache/thrift) | Thrift |
| [github.com/golang-jwt/jwt/v4](https://github.com/golang-jwt/jwt/v4) | Json web token |


## Technology
| 技术 | 简介 |
| --- | --- |
| [Zipkin](https://zipkin.io/) | 链路追踪 |
| [skywalking](https://skywalking.apache.org/) | 链路追踪 |
| [Jaeger](https://www.jaegertracing.io/)   | 链路追踪 |


## Tools
https://mholt.github.io/json-to-go  
https://www.processon.com  
https://www.figma.com  
https://swagger.io  


## Docs
https://bigbully.github.io/Dapper-translation  
https://www.zhihu.com/question/65502802  
https://github.com/avelino/awesome-go  
https://zhuanlan.zhihu.com/p/39326315  
https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1  
https://studygolang.com/articles/17467?fr=sidebar  
http://hbase.org.cn  
https://blog.boot.dev/golang/golang-project-structure  
