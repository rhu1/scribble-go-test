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
	P := *Proto1.NewProto1()

	epA := net.NewMPSTEndpoint(P, P.A)  // FIXME: generate role-specific EP types (no parameterised types)
	epA.Connect(P.B, c)
	a1 := *Proto1.NewProto1_A_1(epA)
	go RunA(a1)

	epB := net.NewMPSTEndpoint(P, P.B)
	epB.Accept(P.A, c)
	b1 := *Proto1.NewProto1_B_1(epB)
	go RunB(b1)

	barrier.Wait()
}


// FIXME: some internal value passing (e.g., roles?) needs to be changed to pointers


//*
func RunA(a1 Proto1.Proto1_A_1) {
	log.Println("A: start")
	defer barrier.Done()

	a1.Send_B_Ok(1234)
}
/*	//defer endA.Close()  // FIXME



	//NewProto1_A_1(

	//log.Println("A: received from B:", y)
}
//*/


//*/
func RunB(b1 Proto1.Proto1_B_1) {
	log.Println("B: start")
	defer barrier.Done()

	var x int
	b1.Recv_A_Ok(&x)

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

	log.Println("B: received from A:", x)
}
*/
