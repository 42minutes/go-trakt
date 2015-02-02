package trakt

// ExtendedInfo for different information levels for calls
type ExtendedInfo string

// Options for ExtendedInfo
const (
	Min            ExtendedInfo = "min"         // Default Returns enough info to match locally.
	Images                      = "images"      // Minimal info and all images.
	Full                        = "full"        // Complete info for an item.
	FullWithImages              = "full,images" // Complete info and all images.
	Metadata                    = "metadata"    // Collection only. Additional video and audio info.
)

// Movie struct for the Trakt v2 API
type Movie struct {
	IDs struct {
		Imdb  int `json:"imdb"`
		Slug  int `json:"slug"`
		Tmdb  int `json:"tmdb"`
		Trakt int `json:"trakt"`
	} `json:"ids"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

// Person struct for the Trakt v2 API
type Person struct {
	IDs struct {
		Imdb   int `json:"imdb"`
		Slug   int `json:"slug"`
		Tmdb   int `json:"tmdb"`
		Trakt  int `json:"trakt"`
		Tvrage int `json:"tvrage"`
	} `json:"ids"`
	Name string `json:"name"`
}
