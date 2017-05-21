//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/test


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/test/Test/Proto2"	
)


func RunProto2() {
	barrier := new(sync.WaitGroup)
	barrier.Add(2)

	P := Proto2.NewProto2()
	c := net.NewGoBinChan(make(chan net.T))

	epA := net.NewMPSTEndpoint(P, P.A)
	go runProto2A(barrier, P, c, epA)

	epB := net.NewMPSTEndpoint(P, P.B)
	go runProto2B(barrier, P, c, epB)

	barrier.Wait()
}


func runProto2A(barrier *sync.WaitGroup, P *Proto2.Proto2, c net.BinChan, epA *net.MPSTEndpoint) {
	log.Println("(A) start")
	defer barrier.Done()

	defer epA.Close()
	epA.Connect(P.B, c)
	a2 := Proto2.NewProto2_A_1(epA)

	a2.Send_B_Ok().Send_B_Bye()

	log.Println("(A) done")
}


func runProto2B(barrier *sync.WaitGroup, P *Proto2.Proto2, c net.BinChan, epB *net.MPSTEndpoint) {
	log.Println("(B) start")
	defer barrier.Done()

	defer epB.Close()
	epB.Accept(P.A, c)
	b1 := Proto2.NewProto2_B_1(epB)

	var b2 *Proto2.Proto2_B_2
	switch cases := b1.Branch_A().(type) {
		case *Proto2.Ok:	
			log.Println("(B) received Ok")
			b2 = cases.Recv_A_Ok()
		case *Proto2.Bye:	
			log.Println("(B) received Bye")
			b2 = cases.Recv_A_Bye()
		default:
			panic("Shouldn't get in here: ")
	}
	switch cases2 := b2.Branch_A().(type) {
		case *Proto2.Ok_2:
			log.Println("(B) received Ok")
			cases2.Recv_A_Ok()
		case *Proto2.Bye_2:
			log.Println("(B) received Bye")
			cases2.Recv_A_Bye()
	}
	
	/*switch cases := b1.Branch_A().(type) {
		case *Proto2.Ok:	
			log.Println("(B) received Ok")
			switch cases2 := cases.Recv_A_Ok().Branch_A().(type) {
				case *Proto2.Ok_2:
					log.Println("(B) received Ok")
					cases2.Recv_A_Ok()
				case *Proto2.Bye_2:
					log.Println("(B) received Bye")
					cases2.Recv_A_Bye()
			}
		case *Proto2.Bye:	
			log.Println("(B) received Bye")
			switch cases2 := cases.Recv_A_Bye().Branch_A().(type) {
				case *Proto2.Ok_2:
					log.Println("(B) received Ok")
					cases2.Recv_A_Ok()
				case *Proto2.Bye_2:
					log.Println("(B) received Bye")
					cases2.Recv_A_Bye()
			}
		default:
			panic("Shouldn't get in here: ")
	}*/

	log.Println("(B) done")
}
