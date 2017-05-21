//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/test


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/hello/Hello/Proto1"	
)


var (
	barrier = new(sync.WaitGroup)
)


func main() {
	log.Println("chan transport")
	barrier.Add(2)

	c := net.NewGoBinChan(make(chan net.T))
	P := *Proto1.NewProto1()

	epA := net.NewMPSTEndpoint(P, P.A)
	go RunA(P, c, epA)

	epB := net.NewMPSTEndpoint(P, P.B)
	go RunB(P, c, epB)

	barrier.Wait()
}


func RunA(P Proto1.Proto1, c *net.GoBinChan, epA *net.MPSTEndpoint) {
	log.Println("A: start")
	defer barrier.Done()

	defer epA.Close()
	epA.Connect(P.B, c)
	a1 := Proto1.NewProto1_A_1(epA)

	var y int

	for y = 0; y < 5; y++ {
		a1 = a1.Send_B_Ok(y)
	}
	a1.Send_B_Bye(y)

	//log.Println("A: received from B:", y)
}


func RunB(P Proto1.Proto1, c *net.GoBinChan, epB *net.MPSTEndpoint) {
	log.Println("B: start")
	defer barrier.Done()

	defer epB.Close()
	epB.Accept(P.A, c)
	b1 := Proto1.NewProto1_B_1(epB)

	var loop = true
	var x int

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
