package week09

// A subscription for a subscriber and the topics related.
type Subscription struct {
	Subscriber Subscriber
	Categories []string
}

func (s Subscription) HasAny(categories ...string) bool {
	for _, cat := range s.Categories {
		for _, c := range categories {
			if cat != c {
				continue
			}

			return true
		}
	}

	return false
}

type Subscriptions []Subscription

func (s Subscriptions) Contains(subscriber Subscriber) bool {
	for _, sub := range s {
		if sub.Subscriber.Identifier() != subscriber.Identifier() {
			continue
		}

		return true
	}

	return false
}

func (s Subscriptions) Subscribers() []Subscriber {
	result := []Subscriber{}
	for _, sub := range s {
		result = append(result, sub.Subscriber)
	}

	return result
}

func (s Subscriptions) SubscribersFor(categories ...string) []Subscriber {
	result := []Subscriber{}
	for _, sub := range s {
		if !sub.HasAny(categories...) {
			continue
		}

		result = append(result, sub.Subscriber)
	}

	return result
}
