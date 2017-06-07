//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/mini
//$ bin/mini.exe


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/mini/Mini/Proto1"	
)


func main() {
	barrier := new(sync.WaitGroup)
	barrier.Add(2)

	c := net.NewGoBinChan(make(chan net.T))
	P := Proto1.NewProto1()

	go runMiniA(barrier, P, c)
	go runMiniB(barrier, P, c)

	barrier.Wait()
}

func runMiniA(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan) {
	log.Println("(A) start")
	defer barrier.Done()

	ep := Proto1.NewEndpointProto1_A(P)
	defer ep.A.Close()
	ep.A.Connect(P.B, c)	

	a1 := Proto1.NewProto1_A_1(ep)

	var y string
	a1.Send_B_a("ABCD").Recv_B_b(&y)
	//a1.Send_B_a("ABCD")

	log.Println("(A) done: ", y)  // FIXME: where is extra space coming from?
}

func runMiniB(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan) {
	log.Println("(B) start")
	defer barrier.Done()

	ep := Proto1.NewEndpointProto1_B(P)
	defer ep.B.Close()
	ep.B.Accept(P.A, c)	

	b1 := Proto1.NewProto1_B_1(ep)

	var x string
	b1.Recv_A_a(&x).Send_A_b(x + x)

	log.Println("(B) received from A:", x)
}
