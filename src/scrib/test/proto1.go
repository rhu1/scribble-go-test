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

	go runProto1B(barrier, P, c)
	go runProto1A(barrier, P, c)

	barrier.Wait()
}


//*
func runProto1A(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan) {
	log.Println("(A) start")
	defer barrier.Done()

	ep := Proto1.NewEndpointProto1_A(P)
	defer ep.A.Close()
	ep.A.Connect(P.B, c)	

	a1 := Proto1.NewProto1_A_1(ep)

	var y int
	a1.Send_B_Ok(1234).Recv_B_PPP()
	//a1.Send_B_Ok(1234)  // FIXME: panic seems non-deterministic...

	log.Println("(A) received from B:", y)
}
//*/


//*/
func runProto1B(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan) {
	log.Println("(B) start")
	defer barrier.Done()

	ep := Proto1.NewEndpointProto1_B(P)
	defer ep.B.Close()
	ep.B.Accept(P.A, c)	

	b1 := Proto1.NewProto1_B_1(ep)

	var x int
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
