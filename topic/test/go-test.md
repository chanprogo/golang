
- 单元测试

单元测试文件名 命名规范为 example_test.go   测试用例 的函数名称 必须以 Test 开头，例如：TestExample	


​单元测试，覆盖 所有分支
​




go test 命令 是一个 按照一定约定 和组织 的测试代码的驱动程序。 

在包目录内，所有以 `_test.go` 为后缀名 的源代码文件 都是 `go test` 测试 的一部分，不会被 `go build` 编译 到最终的可执行文件中。





在 `*_test.go` 文件中 有三种类型的函数，单元测试 函数、基准测试 函数 和 示例函数。

| 类型     | 格式                  | 作用                           |
| -------- | --------------------- | ------------------------------ |
| 测试函数 | 函数名前缀为 Test      | 测试程序的一些逻辑行为 是否正确 |
| 基准函数 | 函数名前缀为 Benchmark | 测试函数的性能                 |
| 示例函数 | 函数名前缀为 Example   | 为文档 提供示例文档             |

`go test` 命令会遍历所有的 `*_test.go` 文件中 符合上述命名规则的函数，然后生成一个临时的 main 包 
用于 调用相应的测试函数，然后 构建并运行、报告 测试结果，最后 清理测试中 生成的临时文件。   