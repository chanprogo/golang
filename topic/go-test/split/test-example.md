



#### 测试函数示例

就像 细胞 是构成我们身体的基本单位，一个软件程序 也是由很多单元组件 构成的。
单元组件 可以是 函数、结构体、方法 和 最终用户 可能依赖的 任意东西。
总之我们需要 确保这些组件 是能够正常运行的。  
单元测试 是一些 利用各种方法 测试 单元组件 的程序，它会将 结果与 预期输出 进行比较。



我们可以为 `go test` 命令 添加 `-v` 参数，查看 测试函数 名称 和 运行时间    

还可以在 `go test` 命令后 添加 `-run` 参数，它对应一个 正则表达式，只有 函数名 匹配上 的测试函数 才会被 `go test` 命令执行。  
`go test -v -run="More"`    