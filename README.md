# go-broadcaster

`go-broadcaster` is an example one-to-many publisher,
built around channel primitives.

I would recommend studying it as a resource
for patterns of concurrent code and API design.

The code for the broadcaster is generally
quite simple. The one exception is its notify
method. It might be valuable to play around
with other implementations, to see how it
changes behavior, and to reason about
what cases the existing version covers.

For example, what's wrong with the most
obvious implementation?

```go
func (b *Broadcaster) notify(c chan<-string, m string) {
  c <- m
}
```

The tests should cover any changes you might
make to this method.

The contract of Broadcaster is that all
subscribers must consume all messages
delivered to them. That is, there is no
buffering or dropping of messages in the
case of a slow consumer. That means, for
example, that it would be problematic to
have a client connected over a websocket
subscribed to the broadcaster.

But that doesn't mean the Broadcaster is
just a toy! If you look at the
DeliverMostRecent function, you'll see
a method of subscribing to Broadcaster
in a way that deals with slow consumers,
in this case, by only sending the most
recent value.

There are many, many other ways to build
new semantics on top of the Broadcaster.
They are possible because of intentional
choices in API design, most importantly
that of leveraging channels.

I've not implemented a way to shut down the
broadcaster, as it should be a natural extension
of what's here.

Contributing
------------

See the [CONTRIBUTING] document.
Thank you, [contributors]!

[CONTRIBUTING]: CONTRIBUTING.md
[contributors]: https://github.com/thoughtbot/go-broadcaster/graphs/contributors

Need Help?
----------

We offer 1-on-1 coaching. [Get in touch] if you'd
like to learn more about building
awesome software in Go.

[Get in touch]: http://coaching.thoughtbot.com/go/?utm_source=github

License
-------

go-broadcaster is Copyright (c) 2015 thoughtbot, inc. It is free software,
and may be redistributed under the terms specified in the [LICENSE] file.

[LICENSE]: /LICENSE

About
-----

![thoughtbot](https://thoughtbot.com/logo.png)

go-broadcaster is maintained and funded by thoughtbot, inc.
The names and logos for thoughtbot are trademarks of thoughtbot, inc.

We love open source software!
See [our other projects][community]
or [hire us][hire] to help build your product.

[community]: https://thoughtbot.com/community?utm_source=github
[hire]: https://thoughtbot.com/hire-us?utm_source=github
