# OpenApi 接口代码生成器

This is a tool that generates API call code in various programming languages based on the content of an OpenAPI document.

根据swagger (openapi) 文档生成客户端接口代码

## Feature

- 支持`v2`/`v3`文档格式
- 支持语言：`ts`，`swift`，`kotlin`，`java`，`python`，`go`
- 支持泛型
- 支持自定义模版

## Install 

源码编译

```
go build -o openapi
```



## Download 

- windows
- mac
- linux

## Tutorial

### start

```shell
./openapi start -h

Quick start

Usage:
   start [flags]

Flags:
  -c, --client_output string   client output file
  -e, --endpoint string        example：https://{server}:{port}/v3/api-docs
  -h, --help                   help for start
  -l, --lang string            kotlin,python,ts,typescript,swift,java,go,golang
  -o, --output string          api output file
  -s, --style string           customize template file
  -v, --version string         openapi version (default "v3")

```

例子

```shell
openapi -l ts -o src/api.ts -c src/client.ts  -v v3 -e http://localhost:8081/v3/api_docs
```

```
openapi -l kotlin -o src/api.kt -c src/client.kt  -v v2 -e http://localhost:8081/v2/api_docs
```



### init

```shell
./openapi init
```

在当前目录下生成`openapi.json`配置文件

### reload

```shell
默认当前路径
./openapi reload

自定义文件路径
./openapi reload -f /app/openapi.json
```

根据`openapi.json`配置文件重新生成接口代码

### Configuration File

```json
{
    //openapi 文档地址
    "endpoint": "http://localhost:8083/v3/api-docs",
    //api文件路径
    "output": "src/api.ts",
    //client文件路径
    "client_output": "src/client.ts",
    //目标语言
    "lang": "ts",
    //自定义模版路径
    "style": "custom.tmpl",
    //openapi 版本
    "version": "v3",
    //忽略路径
    "ignore": ["/error","/v3/"],
    //匹配路径
    "filter": ["/user/"],
    //别名 - 当目标语言遇到语法关键词冲突时使用
    "alias": {
      //属性别名
      "properties": {
        "JsonNode":"any"
      },
      //结构体别名
      "modes": {
        "User":"People"
      },
      //类型别名
      "types": {
        //自定义数据类型
        "string": "CustomString",
        //format 自定义数据类型
       	"interge+int64": "CustomLong" 
      },
      //参数别名
      "parameters": {
        "export": "output"
      }
    },
    //模版文件变量
    "variables": {
      "apiPackage": "com.demo.api",
      "structPackage": "com.demo.dto",
      "clientPackage": "com.demo"
    },
    //泛型
    "generics": {
      //是否开启
      "enable": true,
      //是否展开Root包装类型
      "unfold": false,
      //泛型表达式
      "expressions": {
        //单一类型 - ApiResult<T>{... T data ...}
        "ApiResult": [
          "data"
        ],
        //集合类型 - PageData<T>{... List<T> entities ...}
        "PageData": [
          "entities+"
        ],
        //Map类型 - Map<String,T>{... T data ...}
        "MapResult": [
          "data~"
        ],
        //复合泛型 - KVPair<T,V> {... T key,V value ...}
        "KVPair": [
          "key",
          "value"
        ]
      }
    }
  }
```

### Example

#### server

参考

#### client

##### 