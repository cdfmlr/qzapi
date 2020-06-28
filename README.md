# qzgo

强智教务系统 API 的 Go SDK 实现。

> 本项目基于 [TLingC/QZAPI](https://github.com/TLingC/QZAPI/)。

## Getting Started

First, use go get to install the latest version of the library.

```sh
go get -u github.com/cdfmlr/qzgo
```

Next, include qzapi in your application:

```go
import "github.com/cdfmlr/qzgo"
```

Then enjoy your "强(ruò)智教务系统":

```go
// 登录
client, err := qzgo.NewClientLogin("你的学校", '你本人的学号', '你本人的密码')
// 你的学校填 jwxt.xxxx.edu.cn 的那个 xxxx
if err != nil {
    panic(err)
}

// 查成绩
r, err := client.GetCjcx("目标学号", "")
if err != nil {
    panic(err)
}
for _, c := range r.Result {
    fmt.Println(c.Kcmc, c.Zcj)
}
```

完整使用文档正在施工。

## 已完成的API

- [x] authUser：登录
- [x] getCurrentTime：时间信息
- [x] getCjcx：成绩信息
- [x] getKbcxAzc：课程信息

## 贡献

1. Issues Welcome!
2. Forks & PRs Welcome!
3. 鄙人穷苦，感谢支持🙏
* BTC: `1DgTSywmxeYvwpSxtNaU1zJE3VwK345v9A`

## 开放源代码

Copyright 2020 CDFMLR

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


