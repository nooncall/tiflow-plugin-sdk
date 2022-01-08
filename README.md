# tiflow-plugin-sdk

tiflow-plugin-sdk是针对TiFlow的Go语言插件SDK, 可面向 TiCDC / DM 的WebAssembly插件引擎针对特定功能编写插件.

本项目是 2021 TiDB Hackathon 参赛项目, RFC请参考: https://github.com/nooncall/letetlrock

目前我们对DM组件的 `Column Mapping` 和 `Table Router` 功能实现了WebAssembly插件.

## Quick Start

在examples下新建一个目录, add_suffix, 并新建一个main.go.

```go
package main

import (
	"github.com/nooncall/tiflow-plugin-sdk/proxywasm"
	"github.com/nooncall/tiflow-plugin-sdk/sdk"
)

func main() {
	// 1. 提供一个ColumnMappingHandler
	// 该handler做修改列值, 对该列的值添加-suffix后缀
	handler := func(s string) string {
		return s + "-suffix"
	}

	// 2. 创建wasm VM Context
	vmContext := sdk.NewColumnMappingPlugin(handler)

	// 3. 注册VM Context
	proxywasm.SetVMContext(vmContext)
}
```

执行命令:

```shell
make build.example name=add_suffix
```

在examples/add_suffix/目录下会生成main.go.wasm文件, 将该文件放到tiflow的指定目录下,

并使用修改后的 [tidb-tools](https://github.com/nooncall/tidb-tools) 即可工作.  

## 插件
### Column Mapping

与DM中的Column Mapping规则搭配使用, 可实现自定义列值修改.

### Table Router

与DM中的Table Router规则搭配使用, 可实现扩展列值计算.
