package trakt

import "fmt"

var (
	ShowSeasonEpisodesURL = Hyperlink("shows/{showTraktId}/seasons/{seasonNumber}/episodes")
)

// Create a ShowsService with the base url.URL
func (c *Client) Episodes() (episodes *EpisodesService) {
	episodes = &EpisodesService{client: c}
	return
}

type EpisodesService struct {
	client *Client
}

func (r *EpisodesService) AllBySeason(showTraktId int, seasonNumber int) (episodes []Episode, result *Result) {
	url, _ := ShowSeasonEpisodesURL.Expand(M{"showTraktId": fmt.Sprintf("%d", showTraktId), "seasonNumber": fmt.Sprintf("%d", seasonNumber)})
	result = r.client.get(url, &episodes)
	return
}

// Episode struct for the Trakt v2 API
type Episode struct {
	AvailableTranslations []interface{} `json:"available_translations"`
	FirstAired            string        `json:"first_aired"`
	Ids                   struct {
		Imdb   string      `json:"imdb"`
		Tmdb   int         `json:"tmdb"`
		Trakt  int         `json:"trakt"`
		Tvdb   int         `json:"tvdb"`
		Tvrage interface{} `json:"tvrage"`
	} `json:"ids"`
	Number    int         `json:"number"`
	NumberAbs interface{} `json:"number_abs"`
	Overview  string      `json:"overview"`
	Rating    float64     `json:"rating"`
	Season    int         `json:"season"`
	Title     string      `json:"title"`
	UpdatedAt string      `json:"updated_at"`
	Votes     int         `json:"votes"`
}
