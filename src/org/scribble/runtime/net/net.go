package net;


import "strings"


type LinearResource struct {
	used bool
}


func (cs *LinearResource) Use() {
	if cs.used {
		panic("Linear resource already used.")  // FIXME: panic seems non-deterministic?
	}
	cs.used = true;
}


type Role interface {
	GetRoleName() string
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
	Done bool
}

func (ep *MPSTEndpoint) Close() error {  // FIXME: should be pointer receiver?
	for r, c := range ep.Chans {
		if strings.Compare(ep.Self.GetRoleName(), r.GetRoleName()) < 1 {  // errors?  // FIXME: this hack should only be GoBinChan
			c.Close()
		}
	}
	if !ep.Done {
		panic("MPSTEndpoint incomplete")  // FIXME: integrate better with LinearResource -- MPSTEndpoint should be a "LinResManager", that tracks LinRes's within its scope  // FIXME:  EndSocket special case (not LinRes)
	}
	return nil  // FIXME: ?
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
