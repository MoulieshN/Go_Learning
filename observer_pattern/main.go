package main

import (
	pubsub "observer_pattern/pub_sub"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	indianPublishers := pubsub.GetNewNewsAgency("IndianPublishers")

	obs1 := pubsub.NewObserver(1, "The Hindu", &wg)
	obs2 := pubsub.NewObserver(2, "Newyork times", &wg)

	obs1.StartListeningForNotification()
	obs2.StartListeningForNotification()

	indianPublishers.Register(obs1)
	indianPublishers.Register(obs2)

	event1 := &pubsub.Event{
		Id:      1,
		Message: "Breaking news!!!!",
	}

	event2 := &pubsub.Event{
		Id:      2,
		Message: "Hurrah!!",
	}

	indianPublishers.NotifyAll(event1)
	indianPublishers.NotifyAll(event2)

	time.Sleep(500 * time.Millisecond)

	indianPublishers.Deregister(obs2)

	event3 := &pubsub.Event{
		Id:      3,
		Message: "Very confidential!!",
	}

	indianPublishers.NotifyAll(event3)
	close(obs1.Notify)
	wg.Wait()
}
