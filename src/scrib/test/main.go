//$ Raymond@HZHL3 ~/code/scribble-neon/github-rhu1-go/scribble-go-test
//$ go install scrib/test


package main;


import (
	"log"
	"sync"

	"scrib/test/Test/Proto1"	
)


var (
	barrier = new(sync.WaitGroup)
)


func main() {
	log.Println("chan transport")
	barrier.Add(2)

	AB := make(chan Proto1.T)	

	go RunA(AB)
	go RunB(AB)

	barrier.Wait()
}

func RunA(AB chan Proto1.T) {
	log.Println("A: start")
	defer barrier.Done()

	a1, endA := Proto1.NewA(AB)
	defer endA.Close()

	var y int
	//A.Send_B_1(1234).Recv_B_2(&y)

	//if true
	for y = 0; y < 5; y++ {
		a1 = a1.Send_B_Ok(y)
	} //else
	{
		a1.Send_B_Bye(y)
	}

	//log.Println("A: received from B:", y)
}


// FIXME: message op check (and use underscore version for internal)
// FIXME: linearity


func RunB(BA chan Proto1.T) {
	log.Println("B: start")
	defer barrier.Done()

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
