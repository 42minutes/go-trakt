package trakt

import "fmt"

var (
	ShowSeasonsURL = Hyperlink("shows/{showTraktId}/seasons")
)

// Create a ShowsService with the base url.URL
func (c *Client) Seasons() (seasons *SeasonsService) {
	seasons = &SeasonsService{client: c}
	return
}

type SeasonsService struct {
	client *Client
}

func (r *SeasonsService) All(showTraktId int) (seasons []Season, result *Result) {
	url, _ := ShowsPopularURL.Expand(M{"showTraktId": fmt.Sprintf("%d", showTraktId)})
	result = r.client.get(url, &seasons)
	return
}

// Season struct for the Trakt v2 API
type Season struct {
	EpisodeCount int `json:"episode_count"`
	Ids          struct {
		Tmdb   int         `json:"tmdb"`
		Trakt  int         `json:"trakt"`
		Tvdb   int         `json:"tvdb"`
		Tvrage interface{} `json:"tvrage"`
	} `json:"ids"`
	Number   int         `json:"number"`
	Overview interface{} `json:"overview"`
	Rating   float64     `json:"rating"`
	Votes    int         `json:"votes"`
}
