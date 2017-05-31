/**
  * Created by jemy on 31/05/2017.
  */

object NamedParameter {
  def printName(first: String, last: String): Unit = {
    println(first + " " + last)
  }

  def main(args: Array[String]): Unit = {
    printName("John", "Smith")
    printName(first = "John", last = "Smith")
    printName(last = "Smith", first = "John")
    printName("John", last = "Smith")
    //无法编译
    //printName(last = "Smith", "John")
  }
}