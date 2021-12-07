package week10

import (
	"context"
	"sync"
)

// service in charge of the orchestration for the flow of
// news stories from sources into subscribers.
type service struct {
	// News in the historical archive that the
	// Service keeps track of. these are loaded from the backup file
	// when teh Service is started.
	news []News

	// Subscriptions in the service, these are used to know
	// which news to send to which subscriber.
	subscriptions Subscriptions

	// Sources registry that the service uses to stop the
	// sources when the service is stopped.
	sources []Source

	// sourceCh is used to listen for sources posting through
	// this channel.
	sourceCh chan News

	sync.RWMutex
}

func (s *service) Start(ctx context.Context) error {

	// TODO: Load persisted state
	// The news service should load any saved state from the backup file when it is started.
	// The news service should periodically save the state of the news service to a backup file, in JSON format.

	// The news service should be able to be stopped and started multiple times.
	go func() {
		// Start listening for sources posting news.
		for {
			select {
			case <-ctx.Done():
				return
			case n := <-s.sourceCh:
				// Receive the news implies
				// TODO: Add the news to the historical archive
				sn := s.SubscribersFor(n.Categories...)
				for _, v := range sn {
					go v.Receive(n)
				}
			}
		}
	}()

	return nil
}

func (s *service) Stop() {
	// The news service should be able to be stopped by the end user.
	// The news service should be able to be stopped by the news service itself.
	// The news service should not be able to be stopped by the news sources.
	// The news service should not be able to be stopped by the subscribers.

	// The news service should be able to be stopped and started multiple times.
	// The news service should save the state of the news service to the backup file, in JSON format, when it is stopped.
	// The news service should stop all sources and subscribers when it is stopped.
}

func (s *service) NewsCh() chan News {
	s.Lock()
	defer s.Unlock()

	if s.sourceCh != nil {
		return s.sourceCh
	}

	s.sourceCh = make(chan News)
	return s.sourceCh
}

// Subscribe a subscriber on the passed categories, this method
// register a subscription within the service with both the subscriber
// and the categories. These will be taken into account when receiving
// news.
func (s *service) Subscribe(sc Subscriber, categories []string) error {
	s.Lock()
	defer s.Unlock()

	if s.subscriptions.Contains(sc) {
		return SubscriberExists
	}

	s.subscriptions = append(s.subscriptions, Subscription{
		Subscriber: sc,
		Categories: categories,
	})

	return nil
}

// Unsubscribe removes a subscriber from the service by
// the Identifier of the subscriber.
func (s *service) Unsubscribe(sub Subscriber) {
	s.Lock()
	new := Subscriptions{}
	for _, v := range s.subscriptions {
		if v.Subscriber.Identifier() == sub.Identifier() {
			continue
		}

		new = append(new, v)
	}

	s.subscriptions = new
	s.Unlock()
}

// Subscribers returns the list of subscribers, which is useful
// for testing purposes and other operations
func (s *service) Subscribers() []Subscriber {
	s.Lock()
	defer s.Unlock()

	return s.subscriptions.Subscribers()
}

func (s *service) AddSource(source Source) {
	s.RLock()
	defer s.RUnlock()

	s.sources = append(s.sources, source)
}

func (s *service) Clean() {
	s.Lock()
	defer s.Unlock()

	s.subscriptions = Subscriptions{}
	s.sources = []Source{}
	s.news = []News{}
}

func (s *service) SubscribersFor(categories ...string) []Subscriber {
	s.RLock()
	defer s.RUnlock()

	return s.subscriptions.SubscribersFor(categories...)
}

func (s *service) backup() {
	// Save the state of the news service to the backup file
}

func (s *service) FindBy(id ...int) []News {
	// The news service should provide access to historical news stories by ID number, or range of ID numbers.
	return []News{}
}