package main

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	ID      string
	Faction *Faction
	Claimed bool
	Mutex   sync.Mutex
}

type Faction struct {
	Name string
}

var (
	objects  = make(map[string]*Object)
	factions = make(map[string]*Faction)
)

func NewObject(id string) *Object {
	return &Object{ID: id}
}

func NewFaction(name string) *Faction {
	return &Faction{Name: name}
}

func (o *Object) Claim(f *Faction) {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()
	if !o.Claimed {
		o.Faction = f
		o.Claimed = true
		fmt.Printf("Object %s claimed by faction %s\n", o.ID, f.Name)
	}
}

func (o *Object) Release() {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()
	if o.Claimed {
		fmt.Printf("Object %s released from faction %s\n", o.ID, o.Faction.Name)
		o.Faction = nil
		o.Claimed = false
	}
}

func (o *Object) ConfirmFaction(f *Faction) {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()
	if o.Faction == f {
		fmt.Printf("Object %s confirmed in faction %s\n", o.ID, f.Name)
	} else {
		fmt.Printf("Object %s failed to confirm in faction %s\n", o.ID, f.Name)
	}
}

func (o *Object) LeaveFaction() {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()
	if o.Claimed {
		fmt.Printf("Object %s leaving faction %s\n", o.ID, o.Faction.Name)
		o.Faction = nil
		o.Claimed = false
	}
}

func (f *Faction) KillObject(o *Object) {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()
	if o.Faction == f {
		fmt.Printf("Object %s killed by faction %s\n", o.ID, f.Name)
		delete(objects, o.ID)
	}
}

func main() {
	// Example usage
	f1 := NewFaction("Faction1")
	f2 := NewFaction("Faction2")

	o1 := NewObject("Object1")
	objects[o1.ID] = o1

	o1.Claim(f1)
	time.Sleep(2 * time.Second)
	o1.ConfirmFaction(f1)
	time.Sleep(2 * time.Second)
	o1.LeaveFaction()
	time.Sleep(2 * time.Second)
	f1.KillObject(o1)
}
