//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/rpc


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/rpc/RPC/Proto1"	
)


func main() {
	barrier := new(sync.WaitGroup)
	barrier.Add(2)

	c := net.NewGoBinChan(make(chan net.T))
	P := Proto1.NewProto1()

	go RunC(barrier, P, c)
	go RunS(barrier, P, c)

	barrier.Wait()
}

func RunC(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan) {
	log.Println("(C) start")
	defer barrier.Done()

	ep := Proto1.NewEndpointProto1_C(P)
	defer ep.C.Close()
	ep.C.Connect(P.S, c)	

	c1 := Proto1.NewProto1_C_1(ep)

	var y string
	c1.Send_S_req("ABCD").Recv_S_resp(&y)

	log.Println("(C) received from S:", y)
}

func RunS(barrier *sync.WaitGroup, P *Proto1.Proto1, c net.BinChan) {
	log.Println("(S) start")
	defer barrier.Done()

	ep := Proto1.NewEndpointProto1_S(P)
	defer ep.S.Close()
	ep.S.Accept(P.C, c)	

	s1 := Proto1.NewProto1_S_1(ep)

	var x string
	s1.Recv_C_req(&x).Send_C_resp(x + x)

	log.Println("(S) received from C:", x)
}
