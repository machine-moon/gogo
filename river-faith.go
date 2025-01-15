package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type Event struct {
	Type    string
	Object  *Object
	Faction *Faction
}

var (
	objects  = make(map[string]*Object)
	factions = make(map[string]*Faction)
	events   = make(chan Event, 10)
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

func handleEvents() {
	for event := range events {
		switch event.Type {
		case "claim":
			event.Object.Claim(event.Faction)
		case "release":
			event.Object.Release()
		case "confirm":
			event.Object.ConfirmFaction(event.Faction)
		case "leave":
			event.Object.LeaveFaction()
		case "kill":
			event.Faction.KillObject(event.Object)
		}
	}
}

func main() {
	// Start the event handler
	go handleEvents()

	// Example factions and objects
	f1 := NewFaction("Faction1")
	f2 := NewFaction("Faction2")
	factions[f1.Name] = f1
	factions[f2.Name] = f2

	o1 := NewObject("Object1")
	o2 := NewObject("Object2")
	objects[o1.ID] = o1
	objects[o2.ID] = o2

	// User input loop
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command (claim/release/confirm/leave/kill/exit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		fmt.Print("Enter object ID: ")
		objectID, _ := reader.ReadString('\n')
		objectID = strings.TrimSpace(objectID)
		object, exists := objects[objectID]
		if !exists {
			fmt.Println("Object not found")
			continue
		}

		var faction *Faction
		if input != "release" && input != "leave" {
			fmt.Print("Enter faction name: ")
			factionName, _ := reader.ReadString('\n')
			factionName = strings.TrimSpace(factionName)
			faction, exists = factions[factionName]
			if !exists {
				fmt.Println("Faction not found")
				continue
			}
		}

		events <- Event{Type: input, Object: object, Faction: faction}
		time.Sleep(1 * time.Second) // Give some time for the event to be processed
	}

	// Close the events channel and wait for the handler to finish
	close(events)
	time.Sleep(2 * time.Second)
}
