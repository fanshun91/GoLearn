# 基础语法
- **package**
   - package是最基础的 *分发单位* 和 *工程管理中依赖关系* 的体现
   - 每个 **GO** 语言源代码文件开头都拥有一个 *package* 声明，表示源码文件所属代码包
   - 要生成 **GO** 语言可执行程序，必须要有main的package包，且必须在该包下有main函数 
   - 同一个路径下只能存在一个package，一个package可以拆分成多个源文件组成
- **import**
   - import语句可以导入源码文件所依赖的package包
   - 如果一个main导入其他包，包将会被按顺序导入
   - 如果导入的包中依赖其他包，首先会导入该包，然后初始化该包中的常量和变量，最后如果该包中有init，会自动执行init()
   - 所有包导入完成后才会对main中常量和变量进行初始化，然后执行main中的init函数（如果存在），最后执行main函数
   - 如果该包被导入多次，则其实际导入只有一次
   - *import 的别名*
   - 别名的含义：将导入的包命名为另一个容易记忆和实用的名称
   - 点（.）操作的含义：点（.）标识的包导入后，调用该包中函数可以省略前缀包名
   - 下划线（_）操作的含义：导入该包，但不导入整个包，而是执行该包中的init函数，因为无法通过包名来调用包中的其他函数。使用下划线操作往往是为了注册包里的引擎，让外部可以方便地调用。
- **变量的分组声明**
   - ```
      var (
        a int
        b float32
        name string
      )
     ```
   - 同一行声明多个变量和赋值 `var a, b, c int = 1, 2, 3`或者`a, b := 1, 2`
   - Go语言的类型转换必须是显式的
   - 类型转换只能发生在两种兼容类型之间
   - 类型转换格式：*变量名称 [:]= 目标类型(变量名称)*
   - 大写字母开头的变量可以被导出，也可以被其他包读取，是公用变量
   - 小写字母开头的不可被导出，是私有变量
- **常量**
   - 常量范围目前只支持布尔类型、数值类型、字符串类型
- **iota常见使用法**
   - 1) 跳值使用法
   - 2) 插队使用法
   - 3) 表达式隐式使用法
   - 4) 单行使用法