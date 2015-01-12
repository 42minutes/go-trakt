# Go-Trakt

Trakt API v2 Wrapper for GO (Golang).  
Under heavy development and in dire need of work.

API structure and quite a bit of code has been shamelessly ripped off from
[go-octokit](https://github.com/octokit/go-octokit); a lovely GitHub API
wrapper, the authors of which I owe one or more beers.

## Current Features

* Authenticate via AccessToken
* Retrieve shows by show.ids.trakt
* Retrieve seasons by show.ids.trakt and season.number
* Retrieve episodes by show.ids.trakt and season.number

## Example

```go
package main

import (
  "fmt"

  "github.com/42minutes/go-trakt"
)

func main() {
  client := trakt.NewClient(
    "CLIENT_ID",
    trakt.TokenAuth{AccessToken: "ACCESS_TOKEN"},
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
  fmt.Println(err) // Should be nil
}
```
