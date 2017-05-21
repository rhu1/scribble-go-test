//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/test


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/test/Test/Proto1"	
)


func RunProto1() {
	barrier := new(sync.WaitGroup)
	barrier.Add(2)

	P := Proto1.NewProto1()
	c := net.NewGoBinChan(make(chan net.T))

	epA := net.NewMPSTEndpoint(P, P.A)  // FIXME: generate role-specific EP types (no parameterised types)
	go RunA(barrier, P, c, epA)

	epB := net.NewMPSTEndpoint(P, P.B)
	go RunB(barrier, P, c, epB)

	barrier.Wait()
}


//*
func RunA(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan, epA *net.MPSTEndpoint) {
	log.Println("(A) start")
	defer barrier.Done()

	defer epA.Close()
	epA.Connect(P.B, c)
	a1 := Proto1.NewProto1_A_1(epA)

	/*b := net.Bar{}
	f := b.Bar1()
	log.Println(f)*/

	var y int
	a1.Send_B_Ok(1234).Recv_B_PPP()
	//a1.Send_B_Ok(1234)  // FIXME: panic seems non-deterministic...

	log.Println("(A) received from B:", y)
}


//*/
func RunB(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan, epB *net.MPSTEndpoint) {
	log.Println("(B) start")
	defer barrier.Done()

	defer epB.Close()
	epB.Accept(P.A, c)
	b1 := Proto1.NewProto1_B_1(epB)

	var x int
	//b1.Recv_A_Ok(&x)//.Send_A_Bye(x * 2)
	
	switch cases := b1.Branch_A().(type) {
		case *Proto1.Ok:	
			cases.Recv_A_Ok(&x).Send_A_PPP()
			log.Println("(B) received Ok from A:", x)
		case *Proto1.Bye:	
			cases.Recv_A_Bye(&x)
			log.Println("(B) received Bye from A:", x)
		default:
			panic("Shouldn't get in here: ")
	}
}
//*/
