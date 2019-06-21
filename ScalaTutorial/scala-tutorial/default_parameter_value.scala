/**
  * Created by jemy on 31/05/2017.
  */

class Point(val x: Double = 0, val y: Double = 0)

val point1 = new Point(y = 1)

object DefaultParameterValue {
  def log(message: String, level: String = "INFO") = println(s"$level: $message")

  def main(args: Array[String]): Unit = {
    log("System starting") // 输出 INFO: System starting
    log("User not found", "WARNING") //输出 WARNING: User not found
  }
}