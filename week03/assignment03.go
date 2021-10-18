package week03

import (
	"fmt"
)

var (
	ErrRatingNoPlays       = fmt.Errorf("can't review a movie without watching it first")
	ErrInvalidViewersValue = fmt.Errorf("invalid number of viewers, must be greather than 0")
	ErrTeatherPlayNoMovies = fmt.Errorf("no movies to play")
)

type (
	CritiqueFn func(Movie) (float32, error)

	Movie struct {
		Length int
		Name   string

		plays   int
		viewers int
		ratings []float32
	}

	Theatre struct{}
)

func (mo *Movie) Rate(rate float32) error {
	if mo.plays == 0 {
		return ErrRatingNoPlays
	}

	mo.ratings = append(mo.ratings, rate)

	return nil
}

func (mo *Movie) Play(viewers int) error {
	if viewers <= 0 {
		return ErrInvalidViewersValue
	}

	mo.plays++
	mo.viewers += viewers

	return nil
}

func (mo Movie) Viewers() int {
	return mo.viewers
}

func (mo Movie) Plays() int {
	return mo.plays
}

func (mo Movie) String() string {
	return fmt.Sprintf("%s (%dm) %.1f%%", mo.Name, mo.Length, mo.Rating())
}

func (mo Movie) Rating() float64 {
	if len(mo.ratings) == 0 {
		return 0.0
	}

	var total float64
	for _, rate := range mo.ratings {
		total += float64(rate)
	}

	return total / float64(len(mo.ratings))
}

func (mo *Theatre) Play(viewers int, movies ...*Movie) error {
	if len(movies) == 0 {
		return ErrTeatherPlayNoMovies
	}

	for _, mo := range movies {
		err := mo.Play(viewers)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mo *Theatre) Critique(movies []*Movie, critique CritiqueFn) error {
	for _, mo := range movies {
		err := mo.Play(1)
		if err != nil {
			return fmt.Errorf("error playing on critique: %w", err)
		}

		rating, err := critique(*mo)
		if err != nil {
			return fmt.Errorf("error running critique: %w", err)
		}

		err = mo.Rate(rating)
		if err != nil {
			return fmt.Errorf("error rating on critique: %w", err)
		}
	}

	return nil
}
