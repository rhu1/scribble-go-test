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

	var x int
	var y string
	a1.Send_B_M1(1234).Recv_B_M1(&x).
		Send_B_M1(1234).Recv_B_M1(&x).
		Send_B_M1(1234).Recv_B_M1(&x).
		Send_B_M1(1234).Recv_B_M1(&x).
		Send_B_M2("bye").Recv_B_M2(&y)
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
	var y string
	for loop := true; loop; {
		switch cases := b1.Branch_A().(type) {
			case *Proto1.M1:	
				b1 = cases.Recv_A_M1(&x).Send_A_M1(x + x)
				log.Println("(B) received Ok from A:", x)
			case *Proto1.M2:	
				cases.Recv_A_M2(&y).Send_A_M2("bye")
				log.Println("(B) received Bye from A:", y)
				loop = false
			default:
				panic("Shouldn't get in here: ")
		}
	}
}
//*/
