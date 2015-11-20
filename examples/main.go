package main

import (
	"fmt"

	"github.com/42minutes/go-trakt"
)

func main() {
	client := trakt.NewClient(
		"CLIENT_ID",
		trakt.TokenAuth{AccessToken: "USER_TOKEN"},
	)

	shows, err := client.Shows().AllPopular()
	fmt.Println(shows)
	fmt.Println(err) // Should be nil

	show, err := client.Shows().One(1390)
	fmt.Println(show)
	fmt.Println(err) // Should be nil

	seasons, err := client.Seasons().All(1390)
	fmt.Println(seasons)
	fmt.Println(err) // Should be nil

	episodes, err := client.Episodes().AllBySeason(1390, 1)
	fmt.Println(episodes)
	fmt.Println(err) // Should be nil

	showResults, err := client.Shows().Search("game of thrones")
	for _, showResult := range showResults {
		fmt.Println(showResult.Show)
	}

	showByID, err := client.Shows().OneOfType("tt2802850", "imdb")
	fmt.Println(showByID)

	fmt.Println(err) // Should be nil
}
