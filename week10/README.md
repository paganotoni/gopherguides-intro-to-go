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
        // Subscribing a subscriber to passed categories
        week10.Subscribe(yourSubscriber{}, []string{"sports", "politics"})
        
        // Adding sources to the service.
        week10.AddSource(week10.MockSource{})
        week10.AddSource(week10.FileSource("/path/to/file"))
        
        // Starting the service
        err := week10.Start()
        if err != nil {
            panic(err)
        }

        time.Sleep(time.Second * 10)
        week10.Stop()
    }
```