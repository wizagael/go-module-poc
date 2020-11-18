package song

import "testing"

func TestSongOfTheDay(t *testing.T) {
	song := GetSongOfTheDay()
	if song == "" {
		t.Fail()
	}
}
