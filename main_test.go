package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var titleName = "how I met"
var page = 1
var showId = 1100
var seasonId = 1
var episodeId = 1

func TestCheckInput(t *testing.T) {
	content := []byte(titleName)
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	os.Stdin = tmpfile
	actual := checkInput()
	if actual != titleName {
		t.Errorf("expected '%s', got '%s'", titleName , actual)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

}

func TestGetSearchResult(t *testing.T) {
	result , err := getSearchResult(titleName, page)
	if err != nil {
		t.Errorf("expected 'nil', got '%s'", err)
	}
	if result.TotalResults <= 0 {
		t.Errorf("expected result count above 0 but got %d", result.TotalResults)
	}
}

func TestGetSeasonDetails(t *testing.T) {
	seasonDetails, err := getSeasonDetails(showId,seasonId)
	if err != nil {
		t.Errorf("expected 'nil', got '%s'", err)
	}
	if seasonDetails.SeasonNumber < 0 {
		t.Errorf("expected more then 0 Seasons")
	}
}

func TestGetEpisodeDetails(t *testing.T) {
	episodeDetails, err := getEpisodeDetails(showId,seasonId, episodeId)
	if err != nil {
		t.Errorf("expected 'nil', got '%s'", err)
	}
	if episodeDetails.Overview == "" {
		t.Errorf("expected a none empty string")
	}
}
