package main

type SearchResult struct {
	Page    int `json:"page"`
	Results []struct {
		PosterPath       string   `json:"poster_path"`
		Popularity       float64  `json:"popularity"`
		ID               int      `json:"id"`
		BackdropPath     string   `json:"backdrop_path"`
		VoteAverage      float64  `json:"vote_average"`
		Overview         string   `json:"overview"`
		FirstAirDate     string   `json:"first_air_date"`
		OriginCountry    []string `json:"origin_country"`
		GenreIds         []int    `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		VoteCount        int      `json:"vote_count"`
		Name             string   `json:"name"`
		OriginalName     string   `json:"original_name"`
	} `json:"results"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
}

type ShowDetails struct {
	BackdropPath     string        `json:"backdrop_path"`
	CreatedBy        []interface{} `json:"created_by"`
	EpisodeRunTime   []int         `json:"episode_run_time"`
	FirstAirDate     string        `json:"first_air_date"`
	Homepage         string        `json:"homepage"`
	ID               int           `json:"id"`
	InProduction     bool          `json:"in_production"`
	Languages        []string      `json:"languages"`
	LastAirDate      string        `json:"last_air_date"`
	Name             string        `json:"name"`
	NextEpisodeToAir interface{}   `json:"next_episode_to_air"`
	Networks         []struct {
		Name          string `json:"name"`
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes int      `json:"number_of_episodes"`
	NumberOfSeasons  int      `json:"number_of_seasons"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	Seasons          []struct {
		AirDate      string `json:"air_date"`
		EpisodeCount int    `json:"episode_count"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		PosterPath   string `json:"poster_path"`
		SeasonNumber int    `json:"season_number"`
	} `json:"seasons"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Type        string  `json:"type"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

type SeasonDetails struct {
	StringID       string `json:"_id"`
	AirDate  string `json:"air_date"`
	Episodes []struct {
		AirDate       string `json:"air_date"`
		EpisodeNumber int    `json:"episode_number"`
		Crew          []struct {
			Department         string  `json:"department"`
			Job                string  `json:"job"`
			CreditID           string  `json:"credit_id"`
			Adult              bool    `json:"adult"`
			Gender             int     `json:"gender"`
			ID                 int     `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float64 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
		} `json:"crew"`
		GuestStars []struct {
			CreditID           string  `json:"credit_id"`
			Order              int     `json:"order"`
			Character          string  `json:"character"`
			Adult              bool    `json:"adult"`
			Gender             int     `json:"gender"`
			ID                 int     `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float64 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
		} `json:"guest_stars"`
		ID             int     `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int     `json:"vote_count"`
	} `json:"episodes"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	ID           int    `json:"id"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
}

type EpisodeDetails struct {
	AirDate string `json:"air_date"`
	Crew    []struct {
		ID          int    `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Department  string `json:"department"`
		Job         string `json:"job"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
	EpisodeNumber int `json:"episode_number"`
	GuestStars    []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		CreditID    string `json:"credit_id"`
		Character   string `json:"character"`
		Order       int    `json:"order"`
		ProfilePath string `json:"profile_path"`
	} `json:"guest_stars"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ID             int     `json:"id"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}
