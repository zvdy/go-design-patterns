package main

import "fmt"

// Subject interface
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

// Observer interface
type Observer interface {
	Update(message string)
}

// YouTubeChannel implements Subject
type YouTubeChannel struct {
	observers []Observer
	message   string
}

func (y *YouTubeChannel) RegisterObserver(o Observer) {
	y.observers = append(y.observers, o)
}

func (y *YouTubeChannel) RemoveObserver(o Observer) {
	y.observers = removeFromSlice(y.observers, o)
}

func (y *YouTubeChannel) NotifyObservers() {
	for _, observer := range y.observers {
		observer.Update(y.message)
	}
}

func (y *YouTubeChannel) PostMessage(message string) {
	fmt.Println("YouTubeChannel: Posting message to subscribers...")
	y.message = message
	y.NotifyObservers()
}

// Subscriber implements Observer
type Subscriber struct {
	name string
}

func (s *Subscriber) Update(message string) {
	fmt.Printf("Subscriber %s: Received message: %s\n", s.name, message)
}

// Helper function to remove an observer from the slice
func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	for i, observer := range observerList {
		if observer == observerToRemove {
			return append(observerList[:i], observerList[i+1:]...)
		}
	}
	return observerList
}

func main() {
	youtube := &YouTubeChannel{}

	subscriber1 := &Subscriber{name: "Subscriber 1"}
	subscriber2 := &Subscriber{name: "Subscriber 2"}

	youtube.RegisterObserver(subscriber1)
	youtube.RegisterObserver(subscriber2)

	youtube.PostMessage("Hello, World!")
}
