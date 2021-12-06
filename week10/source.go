package week10

// Sources should be able to publish stories for any category, or categories, they wish to define.
// Sources should be able to self-cancel.
// Sources should be stopped by the news service when it is stopped.
// Sources should not be effected by the removal of a subscriber.
// Sources should not be effected by the removal of another news source.
// Sources should be free to deliver stories as frequently, or as infrequently, as they wish.

type Source interface {
	//
	PublishTo(chan News)

	//
	Stop()
}
