package week09

const (
	SubscriberExists SubscriptionError = "subscriber is already listening to the service"
)

type SubscriptionError string

func (s SubscriptionError) Error() string {
	return string(s)
}
