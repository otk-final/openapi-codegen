# OpenApi （Swagger）接口代码生成器 

This is a tool that generates API call code in various programming languages based on the content of an OpenAPI document.

根据openapi (swagger) 文档生成不同编程语言客户端接口代码

| Language   | Demo                                                                                      |
|------------|-------------------------------------------------------------------------------------------|
| Java       | ![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/api_java.jpg)   |
| Typescript | ![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/api_ts.jpg)     |
| Go         | ![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/api_go.jpg)     |
| Swift      | ![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/api_swift.jpg)  |
| Python     | ![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/api_python.jpg) |
| Kotlin     | ![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/api_kotlin.jpg) |
| C#         | ![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/api_cs.jpg)     |
| Rust       | ![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/api_rust.jpg)   |

![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/home_api.png)

![image](https://github.com/otk-final/openapi-codegen/blob/master/_images/home_struct.png)

## Feature

- 支持`v2`/`v3`文档格式
- 支持语言：`ts`，`swift`，`kotlin`，`java`，`python`，`go`, `rust`, `c#`
- ***支持泛型***
- 支持自定义模版

## Install 

源码编译

```
go build -o openapi
```

## Download - 1.0.2

- [mac](https://github.com/otk-final/openapi-codegen/releases/download/v1.0.2/openapi_darwin.zip)
- [windows](https://github.com/otk-final/openapi-codegen/releases/download/v1.0.2/openapi_windows.zip)
- [linux](https://github.com/otk-final/openapi-codegen/releases/download/v1.0.2/openapi_linux.zip)

安装并添加到环境变量

## Tutorial

### start

```shell
openapi start -h

Quick start

Usage:
   start [flags]

Flags:
  -e, --endpoint string        example：https://{server}:{port}/v3/api-docs
  -h, --help                   help for start
  -l, --lang string            kotlin,python,ts,typescript,swift,java,go,golang,rust,cs(c#)
  -v, --version string         openapi version (default "v3")

```

例子

```shell
openapi -l ts -v v3 -e http://localhost:8080/v3/api_docs
```

```
openapi -l kotlin -v v2 -e http://localhost:8080/v2/api_docs
```



### init

```shell
openapi init
```

在当前目录下生成[openapi.json](https://github.com/otk-final/openapi-codegen/openapi.json)配置文件


### reload

```shell
#默认当前目录下 openapi.json
openapi reload

#自定义文件路径
openapi reload -f /app/openapi.json
```

根据`openapi.json`配置文件重新生成接口代码，默认全部

```shell
#指定name 或者 language
openapi reload -f /app/openapi.json -n ts
```



### Configuration File

```json
[{
    "name:"server_name",
    //openapi 文档地址
    "endpoint": "http://localhost:8083/v3/api-docs",
    //目标语言
    "lang": "ts",
    //openapi 版本
    "version": "v3",
    //输出文件 - 可自定义，系统仅内置（api,model,client）
    "output": {
      // api 模版
      "api": {
        //输出文件 文件头
        "header": [
          "package demo.api;",
          "",
          "import demo.model.*;",
          "import demo.ApiClient;"
        ],
        //自定义模版路径
        "template": "custom.tmpl",
        //输出文件路径
        "file": "src/main/java/demo/api/{api}.java",
        // 自定义变量 - 继承全局变量
        "variables": {
          "k1": "k2"
        }
      },
      // api 模版
      "client": {
        "file": "src/main/java/demo/ApiClient.java",
        // 自定义修改后，设置忽略，防止下次重新生成被覆盖
        "ignore":true
      },
      "model": {
        "file": "src/main/java/demo/model/{model}.java"
      }
    },
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
      "models": {
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
    //模版全局变量
    "variables": {
      "name": "value"
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
    },
    // 解决 openapi重复operation_id 导出带数字的方法名： batch_12 => batch
    "repeatable_operation_id": true
  }]
```

- openapi文档规范中不支持泛型，泛型表达式需在配置文件中预声明
- 当server端采用统一标准泛型结构返回时，如：`ApiResult<T>`时，可开启`generics.unfold`  参数，业务方法可以直接获取`T`
- 集合泛型声明：`属性+` ，采用 `+` 符号
- Map泛型声明：`属性~` ，采用 `~`  符号
- 内置模版文件：[模版文件](https://github.com/otk-final/openapi-codegen/tree/master/tmpl)

### Example

#### server端

> [openapi-server](https://github.com/otk-final/openapi-server)


#### client端

> [ts](https://github.com/otk-final/openapi-codegen/tree/master/_example/ts)
>
> [go](https://github.com/otk-final/openapi-codegen/tree/master/_example/golang)
>
> [java](https://github.com/otk-final/openapi-codegen/tree/master/_example/java)
>
> [kotlin](https://github.com/otk-final/openapi-codegen/tree/master/_example/kotlin)
>
> [swift](https://github.com/otk-final/openapi-codegen/tree/master/_example/swift)
>
> [python](https://github.com/otk-final/openapi-codegen/tree/master/_example/python)
> 
> [c#](https://github.com/otk-final/openapi-codegen/tree/master/_example/cs)
> 
> [rust](https://github.com/otk-final/openapi-codegen/tree/master/_example/rust)