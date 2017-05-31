# 参数默认值

Scala提供了一个指定参数默认值的功能，可以允许调用方忽略这些拥有默认值的参数。

```
object DefaultParameterValue {
  def log(message: String, level: String = "INFO") = println(s"$level: $message")

  def main(args: Array[String]): Unit = {
    log("System starting") // 输出 INFO: System starting
    log("User not found", "WARNING") //输出 WARNING: User not found
  }
}
```

参数 `level` 有一个默认的值，所以它是可选的。在上面代码的最后一行，参数 `"WARNING"` 覆盖了默认的值 `"INFO"`。在Java中你或许需要使用方法重载来实现这种功能，而在Scala中，你可以使用可选参数来实现相同的效果。然而，当调用方忽略一个参数的时候，剩余的参数必须采用命名参数的方式。

```
class Point(val x: Double = 0, val y: Double = 0)

val point1 = new Point(y = 1)
```

这里，我们这样来使用，`y = 1`。

注意，当从Java调用Scala方法时，Scala方法中的默认参数不是可选的。

```
// Point.scala
class Point(val x: Double = 0, val y: Double = 0)

// Main.java
public class Main {
    public static void main(String[] args) {
        Point point = new Point(1);  // 无法编译
    }
}
```