package main

import "fmt"
import "math"

const s string = "constant"

func main() {
   fmt.Println("Hello, World!")
   fmt.Println("go"+"lang")
   fmt.Println("1+1=", 1+1)
   fmt.Println("7.0/3.0", 7.0/3.0)

   fmt.Println(true && true)
   fmt.Println(true || false)
   fmt.Println(!true)
   fmt.Println("----------------------------------------------")

var a = "initial"
fmt.Println(a)

var b, c int = 1, 2
fmt.Println(b, c)

var d = true
fmt.Println(d)

var e int
fmt.Println(e)

f:= "Short"
fmt.Println(f)
//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.

fmt.Println("----------------------------------------------")
fmt.Println(s)

const i = 500000000

const g = 3e20 / i
fmt.Println(g)

fmt.Println(int64(g))

fmt.Println(math.Sin(i))
fmt.Println("----------------------------------------------")

h:=1
for h<=3 {
  fmt.Println(h)
  h = h+1
}

for j:=7; j <=9; j++{
  fmt.Println(j)
}

for{
  fmt.Println("loop")
  break
}

for n:=0 ; n<=5; n++{
  if n%2 == 0{
    //fmt.Println(n)
    continue
  }
  fmt.Println(n)
}

}
