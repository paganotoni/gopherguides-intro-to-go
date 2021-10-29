Inside of the gopherguides-intro-to-go folder create a new folder named week04.

```
gopherguides-intro-to-go
└── week04
```
In your week04 folder/module create a new git branch named, week04. Use the week04 folder for this week's assignment. Create as many files as needed to complete the assignment.

### Instructions
Given the following Venue type and interfaces implement an `Entertain` method on *Venue that takes the number of audience members, int, and a list of `Entertainer`. For each `Entertainer` call its Perform method passing in the Venue. The Venue should check each `Entertainer` to see if it implements the `Setuper` or `Teardowner` interfaces and call them accordingly. The Venue should log all `Setup`, `Perform` and `Teardown` calls. Logging should be written to the Venue.Log field and use the following formats:

- Setup - `%s has completed setup.\n`
- Perform - `%s has performed for %d people.\n`
- Teardown - `%s has completed teardown.\n`

Write test cases, including error cases, for the provided interfaces by implementing the appropriate interfaces, calling the Venue.Entertain method, and checking the logged messages.

You will need to create at least two implementations of the Entertainer interfaces. No type should implement more than two interfaces. No type should implement all of the interfaces.

### Submission
Open a PR named `Week04` to merge your new changes into your main branch. Your PR's title and comments should answer the following questions (800 words minimum):

- What is included in the PR?
- Explain the changes being made.
- What difficulties you ran into, and the resulting architectural choices you made, while doing the assignment?
- While surprises did you find while doing the assignment?