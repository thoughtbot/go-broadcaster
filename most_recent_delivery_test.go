package broadcaster

import (
	"testing"
	"time"
)

func TestDeliverMostRecent(t *testing.T) {
	var (
		b    = New()
		c    = make(chan string)
		stop = make(chan struct{})
	)

	defer close(stop)

	go DeliverMostRecent(b, c, stop)

	// wait for goroutine to subscribe
	<-time.After(time.Millisecond)

	b.Publish("hello")

	if expected, got := "hello", <-c; expected != got {
		t.Fatalf("expected %q, got %q", expected, got)
	}

	b.Publish("hello")
	b.Publish("world")

	if expected, got := "world", <-c; expected != got {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
