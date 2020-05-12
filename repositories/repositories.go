package repositories

import "github.com/jinzhu/gorm"

// NewDBMoviewRepository return new MoviewRepository
func NewDBMovieRepository(dbConn *gorm.DB) MovieRepositoryContract {
	return &DBMovieRepository{
		db: dbConn,
	}
}
