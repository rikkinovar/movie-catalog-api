package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/rikkinovar/movie-catalog-api/models"
)

type DBMovieRepository struct {
	db *gorm.DB
}

// Get returns movies
func (repository *DBMovieRepository) Get() ([]models.Movie, error) {
	movieList := make([]models.Movie, 0)

	repository.db.Debug().Find(&movieList)
	return movieList, nil
}

// Find return movie
func (repository *DBMovieRepository) Find(id int) (models.Movie, error) {
	var movie models.Movie

	repository.db.Where("id = ?", id).Find(&movie)
	return movie, nil
}

//Create saves new movie
func (repository *DBMovieRepository) Create(data models.Movie) (models.Movie, error) {
	if err := repository.db.Debug().Save(&data).Error; err != nil {
		return models.Movie{}, err
	}

	return repository.Find(data.ID)
}

//Delete selected movie
func (repository *DBMovieRepository) Delete(id int) (bool, error) {
	err := repository.db.BlockGlobalUpdate(true).Where("id = ?", id).Delete(&models.Movie{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//Update movie
func (repository *DBMovieRepository) Update(id int, data models.Movie) (models.Movie, error) {
	_, err := repository.Find(id)
	if err != nil {
		return models.Movie{}, err
	}

	err = repository.db.Model(models.Movie{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":        data.Title,
			"genre":        data.Genre,
			"imdb_rating":  data.ImdbRating,
			"poster_url":   data.PosterUrl,
			"release_date": data.ReleaseDate,
		}).Error
	if err != nil {
		return models.Movie{}, err
	}
	return repository.Find(id)
}
