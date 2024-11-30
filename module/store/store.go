package store

import (
	"context"
	"movie-festival-app/entity"
	"time"

	"github.com/lib/pq"
	"go.uber.org/zap"
)

func (c *Client) UpsertMovies(ctx context.Context, req entity.Movie) (err error) {
	qs := `INSERT INTO movie_festival.movies (
	id,
	title,
	duration,
	actors,
	watch_url,
	created_at,
	description,
	poster_url
	) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	on conflict (id) 
	do update set
	title=$2,
	duration=$3,
	actors=$4,
	watch_url=$5,
	updated_at=$6,
	description=$7,
	poster_url=$8
	`

	if _, err = c.postgre.ExecContext(ctx, qs,
		req.Id,
		req.Title,
		req.Duration,
		pq.Array(req.Actors),
		req.WatchUrl,
		time.Now().Unix(),
		req.Description,
		req.PosterUrl,
	); err != nil {
		zap.S().Info(err)
	}
	return
}

func (c *Client) InsertMovieGenres(ctx context.Context, req []entity.MovieGenres) (err error) {
	qs := `INSERT INTO movie_festival.movie_genres (id, movie_id, genre_id, created_at) 
VALUES(uuid_generate_v4(), $1, $2, $3) on conflict (movie_id, genre_id) do nothing`

	for _, v := range req {
		if _, err = c.postgre.ExecContext(ctx, qs, v.MovieId, v.GenreId, time.Now().Unix()); err != nil {
			zap.S().Info(err)
		}
	}

	return
}

func (c *Client) GetMostViewedMovieAndGenre(ctx context.Context) (res entity.MostViewedMovieAndGenre, err error) {
	qs := `select 
(select m.title as result
from movie_festival.movie_viewerships mv 
left join movie_festival.movies m on m.id = mv.movie_id 
group by m.title 
order by count(mv.id) desc limit 1) as most_viewed_movie,
(select g."name" as result
from movie_festival.movie_viewerships mv 
left join movie_festival.movie_genres mg on mv.movie_id = mg.movie_id
left join movie_festival.genres g on g.id = mg.genre_id  
group by g."name"  
order by count(mv.id) desc limit 1) as most_viewed_genre`

	row := c.postgre.QueryRowContext(ctx, qs)
	if err = row.Scan(&res.Movie, &res.Genre); err != nil {
		zap.S().Info(err)
	}
	return
}

func (c *Client) UpsertMovieViewerships(ctx context.Context, req entity.MovieViewerships) (err error) {
	qs := `INSERT INTO movie_festival.movie_viewerships (id, movie_id, watch_duration, created_at)
VALUES ($1, $2, $3, $4) on conflict do update set movie_id = $2, watch_duration = $3, updated_at $4`

	if _, err = c.postgre.ExecContext(ctx, qs, req.Id, req.MovieId, req.WatchDuration, time.Now().Unix()); err != nil {
		zap.S().Info(err)
	}
	return
}
