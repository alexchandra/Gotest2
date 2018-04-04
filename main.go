package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// const s string = "constant"

func hello() (string, error) {
	return "Hello ƛ!", nil
}

type MyEvent struct {
	// Name string `json:"name"`
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

// func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
// 	return fmt.Sprintf("Hello %s!", name.Name), nil
// }

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
	// lambda.Start(HandleRequest)
	lambda.Start(HandleLambdaEvent)

	// fmt.Println("Hello, World!")
	// fmt.Println("go" + "lang")
	// fmt.Println("1+1=", 1+1)
	// fmt.Println("7.0/3.0", 7.0/3.0)
	//
	// fmt.Println(true && true)
	// fmt.Println(true || false)
	// fmt.Println(!true)
	// fmt.Println("----------------------------------------------")
	//
	// var a = "initial"
	// fmt.Println(a)
	//
	// var b, c int = 1, 2
	// fmt.Println(b, c)
	//
	// var d = true
	// fmt.Println(d)
	//
	// var e int
	// fmt.Println(e)
	//
	// f := "Short"
	// fmt.Println(f)
	// //The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.
	//
	// fmt.Println("----------------------------------------------")
	// fmt.Println(s)
	//
	// const i = 500000000
	//
	// const g = 3e20 / i
	// fmt.Println(g)
	//
	// fmt.Println(int64(g))
	//
	// fmt.Println(math.Sin(i))
	// fmt.Println("----------------------------------------------")
	//
	// h := 1
	// for h <= 3 {
	// 	fmt.Println(h)
	// 	h = h + 1
	// }
	//
	// for j := 7; j <= 9; j++ {
	// 	fmt.Println(j)
	// }
	//
	// for {
	// 	fmt.Println("loop")
	// 	break
	// }
	//
	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		//fmt.Println(n)
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }
	// fmt.Println("----------------------------------------------")
	//
	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

}
