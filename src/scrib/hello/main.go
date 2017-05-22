//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/test


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/hello/Hello/Proto1"	
)


func main() {
	barrier := new(sync.WaitGroup)
	barrier.Add(2)

	c := net.NewGoBinChan(make(chan net.T))
	P := Proto1.NewProto1()

	go runHelloB(barrier, P, c)
	go runHelloA(barrier, P, c)

	barrier.Wait()
}


func runHelloA(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan) {
	log.Println("(A) start")
	defer barrier.Done()

	ep := Proto1.NewEndpointProto1_A(P)
	defer ep.A.Close()
	ep.A.Connect(P.B, c)	

	a1 := Proto1.NewProto1_A_1(ep)

	var y int

	for y = 0; y < 5; y++ {
		a1 = a1.Send_B_Ok(y)
	}
	a1.Send_B_Bye(y)

	//log.Println("(A) received from B:", y)
}


func runHelloB(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan) {
	log.Println("(B) start")
	defer barrier.Done()

	ep := Proto1.NewEndpointProto1_B(P)
	defer ep.B.Close()
	ep.B.Accept(P.A, c)	

	b1 := Proto1.NewProto1_B_1(ep)

	var loop = true
	var x int

	for loop {
		switch cases := b1.Branch_A().(type) {
			case *Proto1.Ok:
				b1 = cases.Recv_A_Ok(&x)
				log.Println("(B) received Ok from A:", x)
			case *Proto1.Bye:
				cases.Recv_A_Bye(&x)
				loop = false
				log.Println("(B) received Bye from A:", x)
			default:
				panic("Shouldn't get in here: ")
		}
	}
}
