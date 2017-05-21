package main;


import "fmt"

type A struct {}

func main() {
	fmt.Println(A{})
}


/*
type Role interface {
	isRole() bool
}

type A struct {
	
}	

func (A) isRole() bool {
	return true
}

type P struct {
	A A
}
	

func main() {
	var p P = P{ A: A{}}

	fmt.Println(p.A)	
}*/


/*type ChanState struct{
	done bool
	e error
}

func (cs *ChanState) use() {
	if cs.done {
		panic("already done")
	}
	cs.done = true;
}

type P_1 struct{
	state ChanState
}

func P_init() P_1 {
	return P_1{ state: ChanState{ false, nil } }
}

func (p1 *P_1) a() P_2 {
	p1.state.use()
	fmt.Println("a: ", p1.state.done)
	return P_2{ state: ChanState{} }
}

type P_2 struct{
	state ChanState
}

func (p2 *P_2) b() P_3 {
	p2.state.use()
	fmt.Println("b: ", p2.state.done)
	return P_3{ state: ChanState{} }
}

type P_3 struct{
	state ChanState
}


func main() {
	p1 := P_init()
	p2 := p1.a()
	//p1.a()
	
	p2.b()
	//p2.b()
}*/



/*
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

}*/
