package pubsub

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Id      int
	Message string
}

type Observer struct {
	ID     int
	Name   string
	Notify chan Event
	wg     *sync.WaitGroup
}

// Get the new observer
func NewObserver(id int, name string, wg *sync.WaitGroup) *Observer {
	return &Observer{
		ID:     id,
		Name:   name,
		Notify: make(chan Event, 1),
		wg:     wg,
	}
}

func (O *Observer) StartListeningForNotification() {
	O.wg.Add(1)
	go func() {
		defer O.wg.Done()
		for event := range O.Notify {
			fmt.Printf("[Observer: %d]| Received Event: event id: %d | message: %s \n", O.ID, event.Id, event.Message)
		}
	}()
}

type Subject interface {
	Register(o *Observer)
	Deregister(o *Observer)
	NotifyAll(event *Event)
}

type NewsAgency struct {
	observers map[int]*Observer
	lock      sync.Mutex
	name      string
}

func GetNewNewsAgency(name string) Subject {
	return &NewsAgency{
		observers: make(map[int]*Observer),
		lock:      sync.Mutex{},
		name:      name,
	}
}

func (na *NewsAgency) Register(o *Observer) {
	na.observers[o.ID] = o
	fmt.Printf("Registered observer id: %v | observer name: %v \n", o.ID, o.Name)
}

func (na *NewsAgency) Deregister(o *Observer) {
	if observer, exist := na.observers[o.ID]; exist {
		delete(na.observers, o.ID)
		close(observer.Notify)
		fmt.Printf("De-registered observer id: %v | observer name: %v \n", observer.ID, observer.Name)
	}
}

func (na *NewsAgency) NotifyAll(event *Event) {
	for _, observer := range na.observers {
		/*
			select {
			case observer.Notify <- *event:
				fmt.Printf("Added into the channel: %v \n", event.Id)
			default:
				fmt.Printf("Observer %s is not ready for event\n", observer.Name)
			}
		*/
		/* observer.Notify <- *event */ // NOT GOOD FOR PRODUCTION CODE BASE

		go func(obs *Observer) {
			select {
			case obs.Notify <- *event:
				// Message sent successfully
			case <-time.After(100 * time.Millisecond): // adjust timeout as needed
				fmt.Printf("Timeout: Observer %s is not responding. Skipping event ID %d\n", obs.Name, event.Id)
			}
		}(observer)
	}
}
