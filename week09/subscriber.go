package week09

// Subscriber is the interface that any subscriber will implement to
// receive news stories from the news service.
type Subscriber interface {
	// Identifier for the subscriber, used for unsubscription
	Identifier() string

	// Receive a news story, typically these stories will be filtered
	// by the ListensTo method before being passed to the subscriber.
	Receive(news News)
}

// Subscribers should be able to receive news stories for the categories they are subscribed to.
// Subscribers should be able to subscribe to the news service and receive news stories for one or more categories.
// Subscribers should be able to unsubscribe from the news service. Other subscribers should not be affected.
// Subscribers should be cancelled by the news service when the news service is stopped.
// Subscribers should not be aware of each other, nor should they have any direct contact with the news sources.
// Subscribers should not be effected by the removal of a news source.
