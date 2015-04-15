package broadcaster

// Broadcaster is a one-to-many publisher.
//
// All published messages will be delivered to all
// subscribers.
type Broadcaster struct {
	subscribec   chan chan<- string
	unsubscribec chan chan<- string
	publishc     chan string

	subscribers map[chan<- string]struct{}
}

// New returns a new Broadcaster.
func New() *Broadcaster {
	b := &Broadcaster{
		subscribec:   make(chan chan<- string),
		unsubscribec: make(chan chan<- string),
		publishc:     make(chan string),
		subscribers:  map[chan<- string]struct{}{},
	}

	go b.run()

	return b
}

// Subscribe causes published messages to be sent to c.
//
// It is invalid usage to subscribe without expecting
// to consume all messages sent in a timely manner.
func (b *Broadcaster) Subscribe(c chan<- string) {
	b.subscribec <- c
}

// Unsubscribe causes messages to no longer be send to c.
func (b *Broadcaster) Unsubscribe(c chan<- string) {
	b.unsubscribec <- c
}

// Publish sends m to all subscribers.
func (b *Broadcaster) Publish(m string) {
	b.publishc <- m
}

func (b *Broadcaster) run() {
	for {
		select {
		case c := <-b.subscribec:
			b.subscribers[c] = struct{}{}

		case c := <-b.unsubscribec:
			delete(b.subscribers, c)

		case m := <-b.publishc:
			b.broadcast(m)
		}
	}
}

func (b *Broadcaster) broadcast(m string) {
	for c := range b.subscribers {
		b.notify(c, m)
	}
}

// notify delivers m to c, unless unsubscribe is called for c.
func (b *Broadcaster) notify(c chan<- string, m string) {
	for {
		select {
		case c <- m:
			return

		case unsub := <-b.unsubscribec:
			delete(b.subscribers, unsub)

			if c == unsub {
				return
			}
		}
	}
}
