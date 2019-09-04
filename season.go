package jikan

import (
	"fmt"
)

// Season struct defines a season
type Season struct {
	Year   int    // Year
	Season string // Season
}

// GetSeason returns a map of a season as specified in the Season struct
// Calls responses through the /season/ endpoint
func GetSeason(season Season) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl(fmt.Sprintf("/season/%v/%v", season.Year, season.Season)), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}

// GetSeasonArchive returns a map of season archives
// Calls responses through the /season/archive endpoint
func GetSeasonArchive() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl("/season/archive"), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}

// GetSeasonLater returns a map of a list of anime from seasons later
// Calls responses through the /season/archive endpoint
func GetSeasonLater() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl("/season/later"), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
