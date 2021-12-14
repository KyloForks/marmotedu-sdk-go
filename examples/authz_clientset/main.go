// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/ory/ladon"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/component-base/pkg/util/homedir"

	"github.com/marmotedu/marmotedu-sdk-go/marmotedu"
	"github.com/marmotedu/marmotedu-sdk-go/tools/clientcmd"
)

// ywh: 调用关系：包含项目、应用、服务三层接口。
// 1. API 接口调用格式规范，层次清晰，可以使 API 接口调用更加清晰易记。
// 2. 可以根据需要，自行选择客户端类型，调用灵活。
//      在服务中同时用到 iam-apiserver 和 iam-authz-server 提供的接口，
//      可以创建应用级别客户端，通过 iamclient.APIV1() 和 iamclient.AuthzV1() 调用不同接口。
//
// +------Project-----+      +--marmotedu-sdk-go---+
// | +-ApplicationA-+ |      | +------iam--------+ |
// | | +-ServiceA-+ | |      | | +--apiserver--+ | |
// | | +----------+ | |      | | +-------------+ | |
// | | +-ServiceB-+ | |      | | +-authzserver-+ | |
// | | +----------+ | |      | | +-------------+ | |
// | +--------------+ | ===> | +-----------------+ |
// | +-ApplicationB-+ |      | +------tms--------+ |
// | +--------------+ |      | +-----------------+ |
// +------------------+      +---------------------+
//
// 分层结构：
// API 层：Client
//           ↓
// 基础层：RESTClient -> Request -> gorequest
//                    构建请求信息  完成 HTTP 请求
//
// 提供两种客户端：
// RESTClient：Raw 类型的客户端，可以通过指定 HTTP 的请求方法、请求路径、请求参数等信息，直接发送 HTTP 请求，
//      比如 client.Get().AbsPath("/version").Do().Into() 。
// 基于 RESTClient 封装的客户端：例如 AuthzV1Client、APIV1Client 等，执行特定 REST 资源、特定 API 接口的请求，方便开发者调用。
//
func main() {
	var iamconfig *string
	if home := homedir.HomeDir(); home != "" {
		iamconfig = flag.String(
			"iamconfig",
			filepath.Join(home, ".iam", "config"),
			"(optional) absolute path to the iamconfig file",
		)
	} else {
		iamconfig = flag.String("iamconfig", "", "absolute path to the iamconfig file")
	}
	flag.Parse()

	// use the current context in iamconfig
	// 创建 SDK 的配置实例 config，加载包含了服务的地址和认证信息的配置文件。
	config, err := clientcmd.BuildConfigFromFlags("", *iamconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	// 创建项目级别客户端。
	clientset, err := marmotedu.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	request := &ladon.Request{
		Resource: "resources:articles:ladon-introduction",
		Action:   "delete",
		Subject:  "users:peter",
		Context: ladon.Context{
			"remoteIP": "192.168.0.5",
		},
	}
	// Authorize the request
	fmt.Println("Authorize request...")

	// 请求 /v1/authz 接口执行资源授权请求：
	ret, err := clientset.Iam().AuthzV1().Authz().Authorize(context.TODO(), request, metav1.AuthorizeOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Authorize response: %s.\n", ret.ToString())
}
