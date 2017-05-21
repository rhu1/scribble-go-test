//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/test


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/test/Test/Proto1"	
)


var (
	barrier = new(sync.WaitGroup)
)


func main() {
	log.Println("chan transport")
	barrier.Add(2)

	c := net.NewGoBinChan(make(chan net.T))
	P := *Proto1.NewProto1()  // FIXME: pointer mess

	epA := net.NewMPSTEndpoint(P, P.A)  // FIXME: generate role-specific EP types (no parameterised types)
	go RunA(P, c, epA)

	epB := net.NewMPSTEndpoint(P, P.B)
	go RunB(P, c, epB)

	barrier.Wait()
}


// FIXME: some internal value passing (e.g., roles?) needs to be changed to pointers -- pointer mess in general
// FIXME: label constants need to be separated


//*
func RunA(P Proto1.Proto1, c *net.GoBinChan, epA *net.MPSTEndpoint) {
	log.Println("A: start")
	defer barrier.Done()

	defer epA.Close()
	epA.Connect(P.B, c)
	a1 := Proto1.NewProto1_A_1(epA)

	var y int
	a1.Send_B_Ok(1234)//.Recv_B_Bye(&y)
	//a1.Send_B_Ok(1234)  // FIXME: panic seems non-deterministic...

	log.Println("A: received from B:", y)
}
/*	//defer endA.Close()  // FIXME

	//NewProto1_A_1(
}
//*/


//*/
func RunB(P Proto1.Proto1, c *net.GoBinChan, epB *net.MPSTEndpoint) {
	log.Println("B: start")
	defer barrier.Done()

	defer epB.Close()
	epB.Accept(P.A, c)
	b1 := Proto1.NewProto1_B_1(epB)

	var x int
	//b1.Recv_A_Ok(&x)//.Send_A_Bye(x * 2)
	
	switch cases := b1.Branch_A().(type) {
		case Proto1.Ok:	
			cases.Recv_A_Ok(&x)
			log.Println("B: received from A:", x)
		case Proto1.Bye:	
			cases.Recv_A_Bye(&x)
	}

	log.Println("B: received from A:", x)
}
/*
	b1, endB := Proto1.NewB(BA)
	defer endB.Close()

	var loop = true
	var x int

	//B.Recv_A_1(&x).Send_A_2(x * 2)
	for loop {
		switch cases := b1.Branch_A().(type) {
			case Proto1.Ok:
				b1 = cases.Recv_A_Ok(&x)
				log.Println("B: received from A:", x)
			case Proto1.Bye:
				cases.Recv_A_Bye(&x)
				loop = false
		}
	}
}
*/
