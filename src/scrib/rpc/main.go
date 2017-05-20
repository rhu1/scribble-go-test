//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/rpc


package main;


import (
	"log"
	"sync"

	"scrib/rpc/RPC/Proto1"	
)


var (
	barrier = new(sync.WaitGroup)
)


func main() {
	//log.Println("chan transport")
	barrier.Add(2)

	AB := make(chan Proto1.T)	

	go RunA(AB)
	go RunB(AB)

	barrier.Wait()
}

func RunA(AB chan Proto1.T) {
	log.Println("(C) start")
	defer barrier.Done()

	a1, endA := Proto1.NewC(AB)
	defer endA.Close()

	var y string
	a1.Send_S_req("ABCD").Recv_S_resp(&y)

	log.Println("(C) received from S:", y)
}


// FIXME: case constants for unary send
// FIXME: linearity
// FIXME: message op check (and use underscore version for internal)


func RunB(BA chan Proto1.T) {
	log.Println("(S) start")
	defer barrier.Done()

	b1, endB := Proto1.NewS(BA)
	defer endB.Close()

	var x string
	b1.Recv_C_req(&x).Send_C_resp(x + x)

	log.Println("(S) received from C:", x)
}
