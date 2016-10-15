package trakt

import (
	"fmt"
	"time"
)

var (
	ShowSeasonEpisodesURL      = Hyperlink("shows/{showTraktID}/seasons/{seasonNumber}/episodes")
	ShowSeasonEpisodeNumberURL = Hyperlink("shows/{showTraktID}/seasons/{seasonNumber}/episodes/{episodeNumber}")
)

// Create a ShowsService with the base url.URL
func (c *Client) Episodes() (episodes *EpisodesService) {
	episodes = &EpisodesService{client: c}
	return
}

type EpisodesService struct {
	client *Client
}

// AllBySeason returns all the episodes of a particular Season number.
func (r *EpisodesService) AllBySeason(showTraktID int, seasonNumber int) (episodes []Episode, result *Result) {
	url, _ := ShowSeasonEpisodesURL.Expand(M{
		"showTraktID":  fmt.Sprintf("%d", showTraktID),
		"seasonNumber": fmt.Sprintf("%d", seasonNumber),
	})
	result = r.client.get(url, &episodes)
	return
}

// OneBySeasonByNumber returns one episode of a specific season of a show.
func (r *EpisodesService) OneBySeasonByNumber(showTraktID int, seasonNumber int, episodeNumber int) (episode Episode, result *Result) {
	url, _ := ShowSeasonEpisodesURL.Expand(M{
		"showTraktID":   fmt.Sprintf("%d", showTraktID),
		"seasonNumber":  fmt.Sprintf("%d", seasonNumber),
		"episodeNumber": fmt.Sprintf("%d", episodeNumber),
	})

	result = r.client.get(url, &episode)
	return
}

// Episode struct for the Trakt v2 API
type Episode struct {
	AvailableTranslations []string   `json:"available_translations"`
	FirstAired            *time.Time `json:"first_aired"`
	IDs                   struct {
		Imdb   string `json:"imdb"`
		Tmdb   int    `json:"tmdb"`
		Trakt  int    `json:"trakt"`
		Tvdb   int    `json:"tvdb"`
		Tvrage int    `json:"tvrage"`
	} `json:"ids"`
	Images struct {
		Screenshot struct {
			Full   string `json:"full"`
			Medium string `json:"medium"`
			Thumb  string `json:"thumb"`
		} `json:"screenshot"`
	} `json:"images"`
	Number    int     `json:"number"`
	Overview  string  `json:"overview"`
	Rating    float64 `json:"rating"`
	Season    int     `json:"season"`
	Title     string  `json:"title"`
	UpdatedAt string  `json:"updated_at"`
	Votes     int     `json:"votes"`
}
