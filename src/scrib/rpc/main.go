//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/rpc


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/rpc/RPC/Proto1"	
)


var (
	barrier = new(sync.WaitGroup)
)


func main() {
	//log.Println("chan transport")
	barrier.Add(2)

	c := net.NewGoBinChan(make(chan net.T))
	P := *Proto1.NewProto1()

	epC := net.NewMPSTEndpoint(P, P.C)
	go RunC(P, c, epC)

	epS := net.NewMPSTEndpoint(P, P.S)
	go RunS(P, c, epS)

	barrier.Wait()
}

func RunC(P Proto1.Proto1, c *net.GoBinChan, epC *net.MPSTEndpoint) {
	log.Println("(C) start")
	defer barrier.Done()

	defer epC.Close()
	epC.Connect(P.S, c)
	c1 := Proto1.NewProto1_C_1(epC)

	var y string
	c1.Send_S_req("ABCD").Recv_S_resp(&y)

	log.Println("(C) received from S:", y)
}

func RunS(P Proto1.Proto1, c *net.GoBinChan, epS *net.MPSTEndpoint) {
	log.Println("(S) start")
	defer barrier.Done()

	defer epS.Close()
	epS.Accept(P.C, c)
	s1 := Proto1.NewProto1_S_1(epS)

	var x string
	s1.Recv_C_req(&x).Send_C_resp(x + x)

	log.Println("(S) received from C:", x)
}
