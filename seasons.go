package trakt

import "fmt"

var (
	ShowSeasonsURL = Hyperlink("shows/{showTraktID}/seasons")
)

// Create a ShowsService with the base url.URL
func (c *Client) Seasons() (seasons *SeasonsService) {
	seasons = &SeasonsService{client: c}
	return
}

type SeasonsService struct {
	client *Client
}

// All returns all the seasons of a particular Show. The seasons do not include
// the episodes.
func (r *SeasonsService) All(showTraktID int) (seasons []Season, result *Result) {
	url, _ := ShowSeasonsURL.Expand(M{"showTraktID": fmt.Sprintf("%d", showTraktID)})
	result = r.client.get(url, &seasons)
	return
}

// Season struct for the Trakt v2 API
type Season struct {
	EpisodeCount int `json:"episode_count"`
	IDs          struct {
		Tmdb   int `json:"tmdb"`
		Trakt  int `json:"trakt"`
		Tvdb   int `json:"tvdb"`
		Tvrage int `json:"tvrage"`
	} `json:"ids"`
	Images struct {
		Poster struct {
			Full   string `json:"full"`
			Medium string `json:"medium"`
			Thumb  string `json:"thumb"`
		} `json:"poster"`
		Thumb struct {
			Full string `json:"full"`
		} `json:"thumb"`
	} `json:"images"`
	Number   int     `json:"number"`
	Overview string  `json:"overview"`
	Rating   float64 `json:"rating"`
	Votes    int     `json:"votes"`
}
