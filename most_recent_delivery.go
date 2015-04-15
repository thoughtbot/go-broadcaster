package broadcaster

// DeliverMostRecent delivers messages
// published to b to c, dropping old
// messages if c falls behind.
//
// Close stop to stop delivery.
func DeliverMostRecent(b *Broadcaster, c chan string, stop chan struct{}) {
	var (
		message   string
		deliveryc chan string
		messagec  = make(chan string)
	)

	b.Subscribe(messagec)
	defer b.Unsubscribe(messagec)

	for {
		select {
		case message = <-messagec:
			deliveryc = c

		case deliveryc <- message:
			deliveryc = nil

		case <-stop:
			return
		}
	}
}
