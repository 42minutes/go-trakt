package trakt

import "fmt"

var (
	ShowURL         = Hyperlink("shows/{traktId}")
	ShowsPopularURL = Hyperlink("shows/popular")
	ShowsSearchURL  = Hyperlink("search?query={query}&type=show")
)

// Create a ShowsService with the base url.URL
func (c *Client) Shows() (shows *ShowsService) {
	shows = &ShowsService{client: c}
	return
}

type ShowsService struct {
	client *Client
}

func (r *ShowsService) One(traktId int) (show *Show, result *Result) {
	url, _ := ShowURL.Expand(M{"traktId": fmt.Sprintf("%d", traktId)})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) AllPopular() (shows []Show, result *Result) {
	url, _ := ShowsPopularURL.Expand(M{})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) Search(query string) (shows []ShowResult, result *Result) {
	url, _ := ShowsSearchURL.Expand(M{"query": query})
	result = r.client.get(url, &shows)
	return
}

// Show struct for the Trakt v2 API
type Show struct {
	AiredEpisodes int `json:"aired_episodes"`
	Airs          struct {
		Day      string `json:"day"`
		Time     string `json:"time"`
		Timezone string `json:"timezone"`
	} `json:"airs"`
	AvailableTranslations []string `json:"available_translations"`
	Certification         string   `json:"certification"`
	Country               string   `json:"country"`
	FirstAired            string   `json:"first_aired"`
	Genres                []string `json:"genres"`
	Homepage              string   `json:"homepage"`
	Ids                   struct {
		Imdb   string `json:"imdb"`
		Slug   string `json:"slug"`
		Tmdb   int    `json:"tmdb"`
		Trakt  int    `json:"trakt"`
		Tvdb   int    `json:"tvdb"`
		Tvrage int    `json:"tvrage"`
	} `json:"ids"`
	Images struct {
		Banner struct {
			Full string `json:"full"`
		} `json:"banner"`
		Clearart struct {
			Full string `json:"full"`
		} `json:"clearart"`
		Fanart struct {
			Full   string `json:"full"`
			Medium string `json:"medium"`
			Thumb  string `json:"thumb"`
		} `json:"fanart"`
		Logo struct {
			Full string `json:"full"`
		} `json:"logo"`
		Poster struct {
			Full   string `json:"full"`
			Medium string `json:"medium"`
			Thumb  string `json:"thumb"`
		} `json:"poster"`
		Thumb struct {
			Full string `json:"full"`
		} `json:"thumb"`
	} `json:"images"`
	Language  string  `json:"language"`
	Network   string  `json:"network"`
	Overview  string  `json:"overview"`
	Rating    float64 `json:"rating"`
	Runtime   float64 `json:"runtime"`
	Status    string  `json:"status"`
	Title     string  `json:"title"`
	Trailer   string  `json:"trailer"`
	UpdatedAt string  `json:"updated_at"`
	Votes     int     `json:"votes"`
	Year      int     `json:"year"`
}

type ShowResult struct {
	Score float64 `json:"score"`
	Show  struct {
		Ids struct {
			Imdb   string `json:"imdb"`
			Slug   string `json:"slug"`
			Tmdb   int    `json:"tmdb"`
			Trakt  int    `json:"trakt"`
			Tvdb   int    `json:"tvdb"`
			Tvrage int    `json:"tvrage"`
		} `json:"ids"`
		Images struct {
			Fanart struct {
				Full   string `json:"full"`
				Medium string `json:"medium"`
				Thumb  string `json:"thumb"`
			} `json:"fanart"`
			Poster struct {
				Full   string `json:"full"`
				Medium string `json:"medium"`
				Thumb  string `json:"thumb"`
			} `json:"poster"`
		} `json:"images"`
		Overview string `json:"overview"`
		Title    string `json:"title"`
		Year     int    `json:"year"`
	} `json:"show"`
	Type string `json:"type"`
}
