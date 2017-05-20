package net;


type LinearResource struct {
	used bool
}


func (cs *LinearResource) Use() {
	if cs.used {
		panic("already done")
	}
	cs.used = true;
}


type Role interface {
	IsRole() bool
}


type T interface {}

type BinChan interface {
	Write(t T)	 // FIXME: error?
	Read() T
	Close() error
}


type P interface { }  // FIXME:


type MPSTEndpoint struct {
	Proto P
	Self Role
	Chans map[Role]BinChan
}

func NewMPSTEndpoint(proto P, self Role) *MPSTEndpoint {
	return &MPSTEndpoint { Proto: proto, Self: self, Chans: make(map[Role]BinChan) }
}

func (ep *MPSTEndpoint) Connect(role Role, c BinChan) {   // FIXME: proper client/server connect/accept operations
	ep.Chans[role] = c
}

func (ep *MPSTEndpoint) Accept(role Role, c BinChan) {
	ep.Chans[role] = c
}


type GoBinChan struct {
	c chan T	
}

func NewGoBinChan(c chan T) *GoBinChan {
	return &GoBinChan { c: c }
}

func (bc *GoBinChan) Write(t T) {
	bc.c <- t	
}

func (bc *GoBinChan) Read() T {
	return <-bc.c
}

func (bc *GoBinChan) Close() error {
	close(bc.c)
	return nil  // FIXME?
}
