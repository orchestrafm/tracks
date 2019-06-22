package database

import (
	"time"

	"github.com/spidernest-go/logger"
)

func Remove(id uint64) error {
	err := db.Collection("music").Find(id).Delete()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Entry did not exist or could not be deleted.")
	}
	return err
}