package week04

type Entertainer interface {
	Name() string
	Perform(v Venue)
}
