package main;


import "fmt"
import "scrib/foo/mygopack"
	

type TT interface{}

type Foo struct{}
type Bar struct{}


func main() {
	mygopack.MyFunc()

	f := Foo{}
	b := Bar{}

	var t1 TT = f
	var t2 TT = b

	//f = b        // Error
	t1 = t2
	//f = t1.(Foo)   // Error
	
	var m string = "hello world !"

    fmt.Println(m, f, b, t1, t2)

}
