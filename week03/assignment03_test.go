package week03

import (
	"errors"
	"testing"
)

// DONE: Define a Movie struct and export two public fields; Length of type int and Name of type string. You must NOT export any more public fields on Movie. You are allowed, however, to use non-exported, private fields as needed.
// DONE: Define a method on the pointer receiver for Movie named Rate. Rate should take a float32 rating and return an error. Calling Rate should track this rating. If the number of plays is 0 return the following error: fmt.Errorf()
// DONE: Define a method on the pointer receiver for Movie named Play. Play should take the number (int) of viewers watching the movie. Calling Play should increase both the number of viewers, as well as the number of plays, for the movie.
// DONE: Define a method on the value receiver for Movie named Viewers. Viewers takes no arguments and returns the number (int) of people who have viewed the movie.
// DONE: Define a method on the value receiver for Movie named Plays. Plays takes no arguments and returns the number (int) of times the movie has been played.
// DONE: Define a method on the value receiver for Movie named Rating. Rating takes no arguments and returns the rating (float64) of the movie. This can be calculated by the total ratings for the movie divided by the number of times the movie has been played.
// DONE: Define a method on the value receiver for Movie named String. String should return a string that that includes the name, length, and rating of the film. Ex. Wizard of Oz (102m) 99.0%
// DONE  Define a new type named CritiqueFn that is a function type that takes a pointer of type Movie and returns a float32 and an error. Note: This is not a func definition, but rather a type definition.
// DONE: Define a Theatre struct. You must not export any public fields on Theatre. You are allowed, however, to use non-exported, private fields as needed.

func TestMovieRate(t *testing.T) {
	movie := &Movie{}

	t.Run("NoPlaysReturnError", func(t *testing.T) {
		movie.plays = 0
		err := movie.Rate(0.5)

		if !errors.Is(err, ErrRatingNoPlays) {
			t.Errorf("Expected error to be %v got %v", ErrRatingNoPlays, err)
		}
	})

	t.Run("NoTrackRatings", func(t *testing.T) {
		movie.plays = 1
		err := movie.Rate(0.5)

		if err != nil {
			t.Errorf("Expected error to be nil got %v", err)
		}

		if len(movie.ratings) != 1 {
			t.Errorf("Expected length of ratings to be 1 got %v", len(movie.ratings))
		}
	})
}

func TestMoviePlay(t *testing.T) {
	movie := &Movie{}

	t.Run("PlayPositive", func(t *testing.T) {
		movie.plays = 0

		err := movie.Play(100)
		if err != nil {
			t.Errorf("Expected error to be nil got %v", err)
		}

		if movie.plays != 1 {
			t.Errorf("Expected plays to be 1 got %v", movie.plays)
		}

		if movie.viewers != 100 {
			t.Errorf("Expected viewers to be 100 got %v", movie.viewers)
		}
	})

	t.Run("PlayNegative", func(t *testing.T) {
		movie.plays = 1
		err := movie.Play(-10)

		if err != ErrInvalidViewersValue {
			t.Errorf("Expected error to be %v got %v", ErrInvalidViewersValue, err)
		}
	})
}

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
			if rating != tcase.rating {
				t.Errorf("Expected rating to be %v got %v", tcase.rating, rating)
			}
		})

	}
}

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
		{"NoMovies", 100, ErrTeatherPlayNoMovies, []*Movie{}},
		{"OneMovie", 100, nil, []*Movie{{Name: "The Mask", Length: 102}}},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			err := theatre.Play(tcase.viewers, tcase.movies...)
			if err != tcase.expectedError {
				t.Errorf("Expected error to be %v got %v", tcase.expectedError, err)
			}

			for _, mo := range tcase.movies {
				if mo.plays < 1 {
					t.Errorf("Expected plays to be greater than 0 got %v", mo.plays)
				}

				if mo.viewers != tcase.viewers {
					t.Errorf("Expected viewers to be %v got %v", tcase.viewers, mo.viewers)
				}
			}
		})
	}
}

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

	if err != nil {
		t.Errorf("Expected error to be nil got %v", err)
	}

	for _, mo := range movies {
		if mo.plays != 1 {
			t.Errorf("Expected plays to be 1 got %v", mo.plays)
		}

		if mo.viewers != 1 {
			t.Errorf("Expected viewers to be 1 got %v", mo.viewers)
		}

		if len(mo.ratings) != 1 {
			t.Errorf("Expected length of ratings to be 1 got %v", len(mo.ratings))
		}

		if mo.ratings[0] != 5.0 {
			t.Errorf("Expected rating to be 5.0 got %v", mo.ratings[0])
		}
	}
}
