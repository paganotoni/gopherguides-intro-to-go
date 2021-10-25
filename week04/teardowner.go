package week04

type Teardowner interface {
	Teardown(v Venue) error
}
