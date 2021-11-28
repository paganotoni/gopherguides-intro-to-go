## NewsService

NewsService is a library that allows to listen multiple sources and aggregate these into a single stream of news. It is developed as part of the GopherGuides Go training course.

### Getting started

To get started you should get the library:

```sh
go get github.com/paganotoni/gopherguides-intro-to-go@latest
```

### Usage

The starting point of the library is the Service, to start the library you should initialize an instance of it with:

```go
    import "github.com/paganotoni/gopherguides-intro-to-go/week09"

    func main() {
        service := week09.Service{}

        subscriber = yourSubscriber{}
        // This would add the subscriber to the service
        // so the service notifies it.
        service.Subscribe(subscriber, []string{"sports", "politics"})

        // This would the mockSource to the service
        // so it pushes news to the Service.
        service.AddSource(week09.MockSource)

        // This would the FileSource to the service
        // so it listens to changes on that given folder
        // and pushes news to the Service.
        service.AddSource(week09.FileSource("/path/to/file"))


        // This starts the service and waits for cancellation
        err := service.Start()
        if err != nil {
            panic(err)
        }
    }
```