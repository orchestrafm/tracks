package routing

import (
	"net/http"

	"github.com/orchestrafm/music/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func editTrack(c echo.Context) error {
	t := new(database.Track)
	if err := c.Bind(t); err != nil {
		logger.Error().
			Err(err).
			Msg("Request Data could not be binded to datastructure.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	err := t.Edit()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be updated.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "Track edited successfully."})
}
