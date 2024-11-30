package response

import "gopkg.in/guregu/null.v4"

type MostViewedMovieAndGenre struct {
	Movie null.String `json:"movie"`
	Genre null.String `json:"genre"`
}
