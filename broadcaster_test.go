package broadcaster

import (
	"testing"
	"time"
)

func TestBroadcaster(t *testing.T) {
	var (
		b        = New()
		messages = make(chan string, 1)
	)

	b.Subscribe(messages)
	b.Publish("hello!")

	if expected, got := "hello!", <-messages; expected != got {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

func TestBroadcasterUnsubscribe(t *testing.T) {
	var (
		b        = New()
		messages = make(chan string, 1)
	)

	b.Subscribe(messages)
	b.Unsubscribe(messages)

	b.Publish("hello!")

	select {
	case <-messages:
		t.Fatal("expected not to receive message")
	default:
	}
}

func TestBroadcasterWaitsForSubscribers(t *testing.T) {
	var (
		b        = New()
		messages = make(chan string)
	)

	b.Subscribe(messages)

	go b.Publish("hello!")

	// do some work
	<-time.After(time.Millisecond)

	select {
	case <-messages:
	case <-time.After(time.Millisecond):
		t.Fatal("expected message to be delivered")
	}
}

func TestBroadcasterUnsubscribeDuringPublish(t *testing.T) {
	var (
		b           = New()
		messages    = make(chan string)
		done        = make(chan struct{})
		unsubscribe = make(chan struct{})
	)

	// Start a subscriber
	go func() {
		defer close(done)
		b.Subscribe(messages)

		<-unsubscribe
		b.Unsubscribe(messages)
	}()

	// goroutine 2 publishes a message
	go b.Publish("hello!")

	// receive a message
	<-messages

	// goroutine 2 publishes another message
	go b.Publish("hello!")

	// goroutine 1 unsubscribes
	unsubscribe <- struct{}{}

	// goroutine 1 should terminate
	select {
	case <-done:
	case <-time.After(10 * time.Millisecond):
		t.Fatal("unsubscribe failed to terminate")
	}
}
