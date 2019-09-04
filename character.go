package jikan

import "fmt"

// Character struct defines a character
type Character struct {
	ID      int    // MyAnimeList Character ID
	Request string // Request type (Optional)
}

// GetCharacter returns a map of a character as specified in the Character struct
// Calls responses through the /character/ endpoint
func GetCharacter(character Character) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl(fmt.Sprintf("/character/%v/%v", character.ID, character.Request)), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
