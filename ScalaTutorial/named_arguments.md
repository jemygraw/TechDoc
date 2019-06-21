# 命名参数

当调用方法的时候，你可以使用变量名称来标记参数，就像这样：

```
object NamedParameter {
  def printName(first: String, last: String): Unit = {
    println(first + " " + last)
  }

  def main(args: Array[String]): Unit = {
    printName("John", "Smith")
    printName(first = "John", last = "Smith")
    printName(last = "Smith", first = "John")
    printName("John", last = "Smith")
  }
}
```

注意命名参数的顺序是可以重组的。不过，如果其中有一些参数是采用了命名的方式，另外一些没有，那么没有命名的参数必须出现在命名参数的前面，而且必须遵循它们在方法声明时的顺序。否则像下面的代码是无法编译的：

```
printName(last = "Smith", "John)
```