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


type T interface {}  // Message types

type BinChan interface {
	Write(t T)	 // FIXME: error?
	Read() T
	Close() error
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


type P interface { }  // Generated Protocol types // FIXME?

type Role interface {
	GetRoleName() string
}


type MPSTEndpoint struct {
	Proto P
	Self Role
	Chans map[Role]BinChan
	init bool
	done bool
}

func NewMPSTEndpoint(proto P, self Role) *MPSTEndpoint {
	return &MPSTEndpoint { Proto: proto, Self: self, Chans: make(map[Role]BinChan) }
}

func (ep *MPSTEndpoint) SetInit() {
	ep.init = true
}

func (ep *MPSTEndpoint) SetDone() {
	ep.done = true
}

func (ep *MPSTEndpoint) Close() error {  // FIXME: should be pointer receiver?
	for r, c := range ep.Chans {
		if strings.Compare(ep.Self.GetRoleName(), r.GetRoleName()) < 1 {  // errors?  // FIXME: this hack should only be GoBinChan
			c.Close()
		}
	}
	if !ep.done {
		panic("MPSTEndpoint incomplete")  // FIXME: integrate better with LinearResource -- MPSTEndpoint should be a "LinResManager", that tracks LinRes's within its scope  // FIXME:  EndSocket special case (not LinRes)
	}
	return nil  // FIXME: ?
}

func (ep *MPSTEndpoint) GetChan(role Role) BinChan {
	return ep.Chans[role]
}

func (ep *MPSTEndpoint) Connect(role Role, c BinChan) {   // FIXME: proper client/server connect/accept operations
	ep.checkConnectionAction(role)
	ep.Chans[role] = c  // FIXME: interface types will auto deref the pointer values?
}

func (ep *MPSTEndpoint) Accept(role Role, c BinChan) {
	ep.checkConnectionAction(role)
	ep.Chans[role] = c
}

func (ep *MPSTEndpoint) checkConnectionAction(role Role) {
	if ep.init {
		panic("Illegal accept after initial state channel has been created.")	
	}
	if ep.Chans[role] != nil {
		panic("Illegal duplicate connection with: " + role.GetRoleName())
	}
	if role.GetRoleName() == ep.Self.GetRoleName() {
		panic("Illegal self-connection: " + ep.Self.GetRoleName())
	}
}
