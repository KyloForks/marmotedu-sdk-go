# marmotedu-sdk-go 

一种 client-go风格的go sdk 实现。

Group = marmotedu服务(iam|tms)

## 层级关系
```shell
├── examples                        # 存放 SDK 的使用示例
├── Makefile                        # 管理 SDK 源码，静态代码检查、代码格式化、测试、添加版权信息等
├── marmotedu
│   ├── clientset.go                # clientset 实现，clientset 中包含多个应用，多个服务的 API 接口
│   ├── fake                        # clientset 的 fake 实现，主要用于单元测试
│   └── service                     # 按应用进行分类，存放应用中各服务 API 接口的具体实现
│       ├── iam                     # iam 应用的 API 接口实现，包含多个服务
│       │   ├── apiserver           # iam 应用中，apiserver 服务的 API 接口，包含多个版本
│       │   │   └── v1              # apiserver v1 版本 API 接口
│       │   ├── authz               # iam 应用中，authz 服务的 API 接口
│       │   │   └── v1              # authz 服务 v1 版本接口
│       │   └── iam_client.go       # iam 应用的客户端，包含了 apiserver 和 authz 2个服务的客户端
│       └── tms                     # tms 应用的 API 接口实现
├── pkg                             # 存放一些共享包，可对外暴露
├── rest                            # HTTP 请求的底层实现
├── third_party                     # 存放修改过的第三方包，例如：gorequest
└── tools
    └── clientcmd                   # 一些函数用来帮助创建rest.Config配置
```

## 参考项目

- [medu-sdk-go](https://github.com/marmotedu/medu-sdk-go)
- [client-go](https://github.com/kubernetes/client-go)
