package week09

import (
	"sync"
)

// Service in charge of the orchestration for the flow of
// news stories from sources into subscribers.
type Service struct {
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

	sync.RWMutex
}

func (s *Service) Start() {
	// The news service should be able to be stopped and started multiple times.
	// The news service should load any saved state from the backup file when it is started.
	// The news service should periodically save the state of the news service to a backup file, in JSON format.
}

func (s *Service) Stop() {
	// The news service should be able to be stopped by the end user.
	// The news service should be able to be stopped by the news service itself.
	// The news service should not be able to be stopped by the news sources.
	// The news service should not be able to be stopped by the subscribers.

	// The news service should be able to be stopped and started multiple times.
	// The news service should save the state of the news service to the backup file, in JSON format, when it is stopped.
	// The news service should stop all sources and subscribers when it is stopped.
}

func (s *Service) Receive(n News) {
	// Receive the news implies
	// 1. Add the news to the historical archive

	// 2. Notify subscribers that the news listening on news categories.
	sn := s.subscriptions.SubscribersFor(n.Categories...)
	for _, v := range sn {
		v.Receive(n)
	}
}

func (s *Service) Subscribe(sc Subscriber, categories []string) error {
	s.RLock()
	if s.subscriptions.Contains(sc) {
		s.RUnlock()
		return SubscriberExists
	}
	s.RUnlock()

	s.Lock()
	s.subscriptions = append(s.subscriptions, Subscription{
		Subscriber: sc,
		Categories: categories,
	})
	s.Unlock()

	return nil
}

func (s *Service) Unsubscribe(identifier string) {
	s.Lock()
	new := Subscriptions{}
	for _, v := range s.subscriptions {
		if v.Subscriber.Identifier() == identifier {
			continue
		}

		new = append(new, v)
	}

	s.subscriptions = new
	s.Unlock()
}

// Subscribers returns the list of subscribers, which is useful
// for testing purposes and other operations
func (s *Service) Subscribers() []Subscriber {
	s.Lock()
	defer s.Unlock()

	return s.subscriptions.Subscribers()
}

func (s *Service) backup() {
	// Save the state of the news service to the backup file
}

func (s *Service) FindBy(id ...int) []News {
	// The news service should provide access to historical news stories by ID number, or range of ID numbers.
	return []News{}
}

func (s *Service) Register(source Source) {
	// Add to the sources registry so it can be stopped when the service is stopped.
	// The news service should be able to receive news stories from sources.
	// and publish them to subscribers by checking the listensTo method of the subscribers.
}
