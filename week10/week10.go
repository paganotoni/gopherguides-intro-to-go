package week10

var (
	// sharedService will be the only instance of the service
	// that this package will wrap.
	sharedService = &service{}
)

// Subscribe a subscriber to the passed categories
func Subscribe(sub Subscriber, categories []string) error {
	return sharedService.Subscribe(sub, categories)
}

// Unsubscribe a subscriber to the passed categories
func Unsubscribe(sub Subscriber) {
	sharedService.Unsubscribe(sub)
}

func Start() error {
	return sharedService.Start()
}

func Stop() {
	sharedService.Stop()
}

func Subscribers() []Subscriber {
	return sharedService.Subscribers()
}

func AddSource(source Source) {
	source.PublishTo(sharedService.NewsCh())
	sharedService.AddSource(source)
}

func Clean() {
	sharedService.Clean()
}
