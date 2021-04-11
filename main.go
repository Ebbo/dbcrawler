package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const dbUrl = "https://api.themoviedb.org/3"
const apiKey = "" // put your api key here
const language = "en-US"
const query = "?api_key=" + apiKey + "&language=" + language

func getSearchResult(titleName string, page int) (SearchResult, error) {
	var jsonData []byte
	var result SearchResult
	titleName = strings.ReplaceAll(titleName, " ", "%20%")
	resp, err := http.Get(dbUrl + "/search/tv?api_key=" + apiKey + "&query=" + titleName + "&page=" + strconv.Itoa(page))
	if err != nil {
		return result, errors.New("could not handle request")
	}
	jsonData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, errors.New("could not read response")
	}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, errors.New("could not unmarshal json data")
	}
	return result, err
}

func searchTitle(nameOfTitle string, page int) (SearchResult, string, int, error) {
	var pageNumber = page
	var titleName = nameOfTitle
	var result SearchResult

	result, err := getSearchResult(titleName, page)
	if err != nil {
		return result, titleName, pageNumber, err
	}

	if len(result.Results) <= 0 {
		fmt.Println("No Result found for ", titleName)
		startSearch()
	} else {
		fmt.Printf("There are %d result(s) on %d pages\n", result.TotalResults, result.TotalPages)
		for i := range result.Results {
			fmt.Println(i, result.Results[i].Name)
		}
	}
	fmt.Println("You are on Page", pageNumber)

	selectTitle(result, titleName, pageNumber)
	return result, titleName, pageNumber, err
}

func selectTitle(result SearchResult, titleName string, pageNumber int) {
	fmt.Println("Please enter the number of your title to select " +
		"or press 'n' for the next page or 'p' for the previous page. ")

	input := checkInput()
	getTitleID(result, titleName, pageNumber, input)

}

func getTitleID(result SearchResult, titleName string, pageNumber int, input string) int {
	var selection int

	if input == "n" && pageNumber < result.TotalPages {
		pageNumber = pageNumber + 1
		_, _, _, _ = searchTitle(titleName, pageNumber)

	}
	if input == "p" && pageNumber >= 2 {
		pageNumber = pageNumber - 1
		_, _, _, _ = searchTitle(titleName, pageNumber)
	}

	selection, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("%q is not a number\n", input)
		_, _, _, _ = searchTitle(titleName, pageNumber)
	}

	if selection > len(result.Results) || selection < 0 {
		fmt.Println("The number you entered does not exist")
		_, _, _, _ = searchTitle(titleName, pageNumber)
	}

	_, _ = getTitleDetails(result.Results[selection].ID)
	return result.Results[selection].ID
}

func getTitleDetails(chosenShow int) (ShowDetails, error) {
	var show ShowDetails
	var jsonData []byte

	resp, err := http.Get(dbUrl + "/tv/" + strconv.Itoa(chosenShow) + query)
	if err != nil {
		return show, err
	}
	jsonData, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonData, &show)
	if err != nil {
		return show, err
	}
	fmt.Println("The selected Title"+
		" has", len(show.Seasons), "season(s)")

	for i := range show.Seasons {
		fmt.Println(show.Seasons[i].SeasonNumber, show.Seasons[i].Name)
	}
	err = selectSeason(show, show.ID)
	return show, err
}

func selectSeason(show ShowDetails, chosenShow int) error {
	var seasonId int

	fmt.Println("Please enter the number of your show to select.")
	input := checkInput()

	if _, err := strconv.Atoi(input); err != nil {
		fmt.Printf("%q is not a number\n", input)
		_ = selectSeason(show, chosenShow)
	}

	seasonId, _ = strconv.Atoi(input)

	if seasonId > len(show.Seasons) || seasonId < 0 {
		fmt.Println("The season number you entered does not exist")
		_ = selectSeason(show, chosenShow)
	}
	err := searchSeasonDetails(chosenShow, seasonId)
	return err
}

func searchSeasonDetails(chosenShow, seasonId int) error {
	seasonDetails, _ := getSeasonDetails(chosenShow, seasonId)

	fmt.Println("The selected Season"+
		" has", len(seasonDetails.Episodes), "Episode(s)")

	for i := range seasonDetails.Episodes {
		fmt.Println(seasonDetails.Episodes[i].EpisodeNumber, seasonDetails.Episodes[i].Name)
	}
	err := selectEpisode(chosenShow, seasonId, seasonDetails)
	return err
}

func getSeasonDetails(chosenShow int, seasonId int) (SeasonDetails, error) {
	var seasonDetails SeasonDetails
	var jsonData []byte

	resp, err := http.Get(dbUrl + "/tv/" + strconv.Itoa(chosenShow) + "/season/" + strconv.Itoa(seasonId) + query)
	if err != nil {
		return seasonDetails, errors.New("could not handle request")
	}
	jsonData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return seasonDetails, errors.New("could not read response")
	}
	err = json.Unmarshal(jsonData, &seasonDetails)
	if err != nil {
		return seasonDetails, errors.New("could not unmarshal json data")
	}
	return seasonDetails, err
}

func selectEpisode(chosenShow, seasonId int, season SeasonDetails) error {
	fmt.Println("Please enter the number of your episode to select.")
	input := checkInput()

	if _, err := strconv.Atoi(input); err != nil {
		fmt.Printf("%q is not a number\n", input)
		_ = selectEpisode(chosenShow, seasonId, season)
	}

	episodeId, _ := strconv.Atoi(input)

	if episodeId > len(season.Episodes) || episodeId < 1 {
		fmt.Println("The season number you entered does not exist")
		_ = selectEpisode(chosenShow, seasonId, season)
	}
	episodeId, _ = strconv.Atoi(input)

	episode, err := getEpisodeDetails(chosenShow, seasonId, episodeId)
	printEpisodeDetails(episode)
	return err
}

func getEpisodeDetails(chosenShow, seasonId, episodeId int) (EpisodeDetails, error) {
	var episode EpisodeDetails
	var jsonData []byte

	resp, err := http.Get(dbUrl + "/tv/" + strconv.Itoa(chosenShow) + "/season/" + strconv.Itoa(seasonId) + "/episode/" + strconv.Itoa(episodeId) + query)
	if err != nil {
		return episode, errors.New("could not handle request")
	}
	jsonData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return episode, errors.New("could not read response")
	}
	err = json.Unmarshal(jsonData, &episode)
	if err != nil {
		return episode, errors.New("could not unmarshal json data")
	}
	return episode, nil
}

func printEpisodeDetails(episode EpisodeDetails) {
	// TODO make better or more details
	fmt.Println("Details:")
	fmt.Println("Season:", episode.SeasonNumber, "Episode Number:", episode.EpisodeNumber)
	fmt.Println("Title:", episode.Name)
	fmt.Println("Plot:\n", episode.Overview)
	nextSearch()
}

func checkInput() string {
	var input = ""
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() && scanner.Err() == nil {
		input = scanner.Text()
	} else {
		fmt.Println("Received EOF via CTRL-D - Closing App.")
		os.Exit(1)
	}
	return input
}

func startSearch() {
	fmt.Println("Please enter a name of series you want to get information about.")
	input := checkInput()
	_, _, _, _ = searchTitle(input, 1)
}

func nextSearch() {
	fmt.Println("Do want to search again? (y/N)")
	input := checkInput()

	if input == "y" || input == "Y" {
		startSearch()
	}
}

func main() {
	startSearch()
}
