package song

import "testing"

func TestSongOfTheDay(t *testing.T) {
	song := getSongOfTheDay()
	if song == "" {
		t.Fail()
	}
}
