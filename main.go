/*
Package musicbrainz-go provides a client for interacting with the MusicBrainz API.

This package allows querying the MusicBrainz database for music recordings and related metadata.
It provides a simple interface to search recordings by title and duration through the MusicBrainz web service API v2.

I am not so fluent in go take this package with grain of salt.
I am learning while creating this package


Basic usage:

	client := musicbrainzgo.NewClient()
	opts := musicbrainzgo.SearchOptions{
		Query: "song title",
		MusicLength: 180000 // duration in milliseconds
	}
	recordings, err := client.SearchRecordings(opts)


For more information about the MusicBrainz API, visit: https://musicbrainz.org/doc/Development/XML_Web_Service/Version_2

Authored by: Shubhang Sharma
Github: https://github.com/rustedshader
*/

package musicbrainzgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const API_ROOT_URL = "https://musicbrainz.org/ws/2"

type MusicBrainz struct {
	Created    time.Time   `json:"created"`
	Count      int         `json:"count"`
	Offset     int         `json:"offset"`
	Recordings []Recording `json:"recordings"`
}

type Recording struct {
	ID             string         `json:"id"`
	Score          int            `json:"score"`
	Title          string         `json:"title"`
	Length         int            `json:"length"`
	Disambiguation string         `json:"disambiguation,omitempty"`
	ArtistCredit   []ArtistCredit `json:"artist-credit"`
}

type ArtistCredit struct {
	Name   string `json:"name"`
	Artist struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		SortName string `json:"sort-name"`
	} `json:"artist"`
}

type MusicBrainzClient struct {
	HTTPClient *http.Client
}

// Search Parameter :)
type SearchOptions struct {
	Query       string
	MusicLength int
}

// Creates a http client
// Used pointer because its efficient helps maintain same state and modification to be done to the orignial instance not to copy instance
func NewClient() *MusicBrainzClient {
	return &MusicBrainzClient{
		HTTPClient: &http.Client{},
	}
}

func (c *MusicBrainzClient) SearchRecordings(opts SearchOptions) ([]Recording, error) {
	lengthFilter := ""
	if opts.MusicLength > 0 {
		lengthFilter = fmt.Sprintf(" AND dur:[%d TO %d]", opts.MusicLength-500, opts.MusicLength+500)
	}

	queryString := fmt.Sprintf("/recording/?query=%s%s", url.QueryEscape(opts.Query), url.QueryEscape(lengthFilter))
	apiURL := API_ROOT_URL + queryString

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var musicBrainzResponse MusicBrainz
	if err := json.Unmarshal(body, &musicBrainzResponse); err != nil {
		return nil, fmt.Errorf("parsing JSON response: %w", err)
	}

	return musicBrainzResponse.Recordings, nil
}
