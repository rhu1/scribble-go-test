//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/twobuyer
//$ bin/twobuyer.exe


package main;


import (
	"log"
	"sync"

	"org/scribble/runtime/net"	

	"scrib/twobuyer/TwoBuyer/TwoBuyer"	
)


func main() {
	barrier := new(sync.WaitGroup)
	barrier.Add(3)

	c_AS := net.NewGoBinChan(make(chan net.T))
	c_BS := net.NewGoBinChan(make(chan net.T))
	c_AB := net.NewGoBinChan(make(chan net.T))
	P := TwoBuyer.NewTwoBuyer()

	go runTwoBuyerS(barrier, P, c_AS, c_BS)
	go runTwoBuyerB(barrier, P, c_BS, c_AB)
	go runTwoBuyerA(barrier, P, c_AS, c_AB)

	barrier.Wait()
}


func runTwoBuyerA(barrier *sync.WaitGroup, P *TwoBuyer.TwoBuyer, c_AS net.BinChan, c_AB net.BinChan) {
	log.Println("(A) start")
	defer barrier.Done()

	ep := TwoBuyer.NewEndpointTwoBuyer_A(P)
	defer ep.A.Close()
	ep.A.Connect(P.S, c_AS)	
	ep.A.Connect(P.B, c_AB)	

	a1 := TwoBuyer.NewTwoBuyer_A_1(ep)

	var quote int
	a1.Send_S_title("Homo Deus").Recv_S_quote(&quote).Send_B_quoteByTwo(quote/2)

	log.Println("(A) received quote from B:", quote)
}


func runTwoBuyerB(barrier *sync.WaitGroup, P *TwoBuyer.TwoBuyer, c_BS net.BinChan, c_AB net.BinChan) {
	log.Println("(B) start")
	defer barrier.Done()

	ep := TwoBuyer.NewEndpointTwoBuyer_B(P)
	defer ep.B.Close()
	ep.B.Connect(P.S, c_BS)	
	ep.B.Accept(P.A, c_AB)	

	b1 := TwoBuyer.NewTwoBuyer_B_1(ep)

	var quote, quoteByTwo int
	var date TwoBuyer.Date
	b3 := b1.Recv_S_quote(&quote).Recv_A_quoteByTwo(&quoteByTwo)
	log.Println("(B) received quote and quoteByTwo from A and S:", quote, quoteByTwo)
	if (quote - quoteByTwo) < 20 {
		b3.Send_S_Ok(TwoBuyer.Address{ "my address" }).Recv_S_(&date)	
		log.Println("(B) received Date from S:", date)
	} else {
		b3.Send_S_Quit()	
	}
}


func runTwoBuyerS(barrier *sync.WaitGroup, P *TwoBuyer.TwoBuyer, c_AS net.BinChan, c_BS net.BinChan) {
	log.Println("(S) start")
	defer barrier.Done()

	ep := TwoBuyer.NewEndpointTwoBuyer_S(P)
	defer ep.S.Close()
	ep.S.Accept(P.A, c_AS)	
	ep.S.Accept(P.B, c_BS)	

	s1 := TwoBuyer.NewTwoBuyer_S_1(ep)

	var title string
	quote := 30
	switch cases := s1.Recv_A_title(&title).Send_A_quote(quote).Send_B_quote(quote).Branch_B().(type) {
		case *TwoBuyer.Ok:
			var addr TwoBuyer.Address
			cases.Recv_B_Ok(&addr).Send_B_(TwoBuyer.Date{ "date" })
			log.Println("(S) received Ok from B: " + addr.Addr)
		case *TwoBuyer.Quit:
			cases.Recv_B_Quit()
			log.Println("(S) received Quit from B: ")
		default: panic ("Shouldn't get in here: ")
	}
}