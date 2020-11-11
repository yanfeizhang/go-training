
# 关键字
关键字是指被编程语言保留页不让编程人员作为标识符使用的字符序列。因此，关键字也称为保留字

Go 语言中所有的关键只有25个：

- 程序声明：import、package
- 程序实体声明和定义：chan、const、func、interface、map、struct、type、var
- 程序流程控制：go、select、break、case、continue、default、defer、else、fallthrough、for、goto、if、range、return、switch

# 预定义标识符
在Go语言代码中，每一个标识符可以代表一个变更或者一个类型（即标识符可以被看作是变量或者类型的代号或者名称），标识符是由若干字母、下划线（_）和数字组成的字符序列，第一个字符必须是字母。

在Go语言中还存在一类特殊的标识符，叫作预定义标识符，这些字符序列同样也是不能让开发者使用的。

- 所有基本数据类型的名称，如int、uint、string等
- 接口类型 error
- 常量 true、false 以及 iota
- 所有内部函数的名称，即 append、cap、close、complex、copy、delete、imag、len、make、new、panic、print、println、real和 recover
