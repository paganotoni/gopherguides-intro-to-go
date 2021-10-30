package week03

import (
	"errors"
	"fmt"
	"gopherguides-intro-to-go/internal/verify"
	"testing"
)

// ✅ Define a Movie struct and export two public fields; Length of type int and Name of type string. You must NOT export any more public fields on Movie. You are allowed, however, to use non-exported, private fields as needed.
// ✅ Define a method on the pointer receiver for Movie named Rate. Rate should take a float32 rating and return an error. Calling Rate should track this rating. If the number of plays is 0 return the following error: fmt.Errorf()
func TestMovieRate(t *testing.T) {
	movie := &Movie{}

	t.Run("NoPlaysReturnError", func(t *testing.T) {
		movie.plays = 0
		err := movie.Rate(0.5)

		expErr := errors.New("can't review a movie without watching it first")
		verify.Equals(t, expErr, err)
	})

	t.Run("NoTrackRatings", func(t *testing.T) {
		movie.plays = 1
		err := movie.Rate(0.5)

		verify.Equalsf(t, err, nil, "expected error to be %v, got %v")
		verify.Equalsf(t, 1, len(movie.ratings), "expected length of ratings to be %v got %v")
	})
}

// ✅  Define a method on the pointer receiver for Movie named Play. Play should take the number (int) of viewers watching the movie. Calling Play should increase both the number of viewers, as well as the number of plays, for the movie.
func TestMoviePlay(t *testing.T) {
	movie := &Movie{}

	t.Run("PlayPositive", func(t *testing.T) {
		movie.plays = 0

		err := movie.Play(100)
		verify.Equalsf(t, err, nil, "expected error to be %v, got %v")
		verify.Equalsf(t, 1, movie.plays, "expected movieplays to be %v got %v")
		verify.Equalsf(t, 100, movie.viewers, "expected movie viewers to be %v got %v")
	})

	t.Run("PlayNegative", func(t *testing.T) {
		movie.plays = 1
		err := movie.Play(-10)
		expErr := fmt.Errorf("invalid number of viewers, must be greather than 0")
		verify.Equalsf(t, err, expErr, "expected error to be %v, got %v")
	})
}

// ✅ Define a method on the value receiver for Movie named Rating. Rating takes no arguments and returns the rating (float64) of the movie. This can be calculated by the total ratings for the movie divided by the number of times the movie has been played.
func TestMovieRating(t *testing.T) {
	tcases := []struct {
		name   string
		values []float32
		rating float64
	}{
		{"NoRating", []float32{}, 0.0},
		{"OneRating", []float32{3.5}, 3.5},
		{"SomeMore", []float32{3.5, 3.0, 2.5}, 3.0},
		{"SomeMoreTwo", []float32{0.0, 5.0, 0.5, 0.0}, 1.375},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			movie := &Movie{}
			movie.ratings = tcase.values

			rating := movie.Rating()
			verify.Equalsf(t, tcase.rating, rating, "expected rating to be %v got %v")
		})

	}
}

// ✅ Define a method on the value receiver for Movie named Plays. Plays takes no arguments and returns the number (int) of times the movie has been played.
func TestMoviePlays(t *testing.T) {
	mo := Movie{}
	verify.Equalsf(t, 0, mo.Plays(), "expected plays to be %v got %v")

	mo.plays = 100
	verify.Equalsf(t, 100, mo.Plays(), "expected plays to be %v got %v")
}

// ✅ Define a method on the value receiver for Movie named Viewers. Viewers takes no arguments and returns the number (int) of people who have viewed the movie.
func TestMovieViewers(t *testing.T) {
	mo := Movie{}
	verify.Equalsf(t, 0, mo.Viewers(), "expected viewers to be %v got %v")

	mo.viewers = 100
	verify.Equalsf(t, 100, mo.Viewers(), "expected viewers to be %v got %v")
}

// ✅ Define a method on the value receiver for Movie named String. String should return a string that that includes the name, length, and rating of the film. Ex. Wizard of Oz (102m) 99.0%
func TestMovieString(t *testing.T) {
	tcases := []struct {
		name  string
		movie Movie
		str   string
	}{
		{"NoRating", Movie{Name: "AAA"}, "AAA (0m) 0.0%"},
		{"OneRating", Movie{Name: "The Matrix", Length: 102, ratings: []float32{3.5}}, "The Matrix (102m) 3.5%"},
		{"MultiRating", Movie{Name: "The Matrix Reloaded", Length: 102, ratings: []float32{3.5, 100.0}}, "The Matrix Reloaded (102m) 51.8%"},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			verify.Equalsf(t, tcase.str, tcase.movie.String(), "expected string to be %v got %v")
		})
	}
}

// ✅ Define a Theatre struct. You must not export any public fields on Theatre. You are allowed, however, to use non-exported, private fields as needed.
// ✅ Define a method on the pointer receiver for Theatre named Play.
// ✅ Play should take the number (int) of viewers, a variadic list of pointer of type Movie, and return an error.
// ✅ Calling Play will call each movie’s Play method  with the number of viewers.
// ✅ If no movies are passed in return the following error: fmt.Errorf("no movies to play").
func TestTeatherPlay(t *testing.T) {
	theatre := &Theatre{}

	tcases := []struct {
		name          string
		viewers       int
		expectedError error
		movies        []*Movie
	}{
		{"NoMovies", 100, fmt.Errorf("no movies to play"), []*Movie{}},
		{"OneMovie", 100, nil, []*Movie{{Name: "The Mask", Length: 102}}},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			err := theatre.Play(tcase.viewers, tcase.movies...)
			verify.Equalsf(t, tcase.expectedError, err, "expected error to be %v got %v")

			for _, mo := range tcase.movies {
				if mo.plays < 1 {
					t.Errorf("expected plays to be greater than 0 got %v", mo.plays)
				}

				verify.Equalsf(t, tcase.viewers, mo.viewers, "expected viewers to be %v got %v")
			}
		})
	}
}

// ✅ Define a new type named CritiqueFn that is a function type that takes a pointer of type Movie and returns a float32 and an error. Note: This is not a func definition, but rather a type definition.
// ✅ Define a method on the value receiver for Theatre named Critique.
// ✅ Critique should take a slice of pointers to type Movie, a CritiqueFn, and return an error.
// ✅ Calling Critique will iterate over the movies and for each movie.
// ✅ First, each movie’s Play method should be called with a value of 1.
// ✅ Next, the CritiqueFn should be called with the movie, the return values should be error checked.
// ✅ If there is no error the movie’s Rate method should be called with the float32 value that was returned from the CritiqueFn.
// ✅ Again, this call should be error checked.
func TestTeatherCritique(t *testing.T) {
	movies := []*Movie{
		{Name: "The Mask", Length: 102},
		{Name: "Ace Ventura", Length: 102},
		{Name: "The Matrix", Length: 102},
		{Name: "The Matrix Reloaded", Length: 102},
	}

	lazyCritique := func(mo Movie) (float32, error) {
		return 5.0, nil
	}

	theatre := &Theatre{}
	err := theatre.Critique(movies, lazyCritique)
	verify.Equalsf(t, nil, err, "expected err to be %v got %v")

	for _, mo := range movies {
		verify.Equalsf(t, 1, mo.plays, "expected plays to be %v got %v")
		verify.Equalsf(t, 1, mo.viewers, "expected viewers to be %v got %v")
		verify.Equalsf(t, 1, len(mo.ratings), "expected to have %v rating got %v")
		verify.Equalsf(t, 0.5, mo.ratings[0], "expected first rating to be %v got %v")
	}
}
