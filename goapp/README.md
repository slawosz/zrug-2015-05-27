Queues setup

By default, goworker uses namespace `resque` and prepends `queue` to given queue name.
In order to use it with default `sidekiq` queue `queue:default`, goworker needs to
be runned like this:

```
bin/worker -queues="default" -namespace=""
```

Usage:

Compilation:

```
export GOPATH=`pwd`
make deps
make
```

Run:

```
bin/worker
bin/http
```
