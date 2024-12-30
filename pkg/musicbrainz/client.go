/*
Package musicbrainz-go provides a client for interacting with the MusicBrainz API.

This package allows querying the MusicBrainz database for music recordings and related metadata.
It provides a simple interface to search MusicBrainz API.

Authored by: Shubhang Sharma
Github: https://github.com/rustedshader
*/

package musicbrainzgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_ROOT_URL = "https://musicbrainz.org/ws/2"

type MusicBrainzClient struct {
	HTTPClient *http.Client
}

// Creates a http client
// Used pointer because its efficient helps maintain same state and modification to be done to the orignial instance not to copy instance
func Client() *MusicBrainzClient {
	return &MusicBrainzClient{
		HTTPClient: &http.Client{},
	}
}

func (c *MusicBrainzClient) searchEntity(query string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, query, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("performing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("parsing JSON response: %w", err)
	}

	return nil
}

func (c *MusicBrainzClient) SearchAnnotation(opts MusicBrainzAnnotationParameters) (MusicBrainzAnnotationResponse, error) {
	apiUrl, err := buildAnnotationSearchQuery(opts)
	if err != nil {
		return MusicBrainzAnnotationResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzAnnotationResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchArea(opts MusicBrainzAreaParameters) (MusicBrainzAreaResponse, error) {
	apiUrl, err := buildAreaSearchQuery(opts)
	if err != nil {
		return MusicBrainzAreaResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzAreaResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchArtist(opts MusicBrainzArtistParameters) (MusicBrainzArtistResponse, error) {
	apiUrl, err := buildArtistSearchQuery(opts)
	if err != nil {
		return MusicBrainzArtistResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzArtistResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchCDStub(opts MusicBrainzCDStubsParameters) (MusicBrainzCDStubsResponse, error) {
	apiUrl, err := buildCDStubSearchQuery(opts)
	if err != nil {
		return MusicBrainzCDStubsResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzCDStubsResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchEvent(opts MusicBrainzEventParameters) (MusicBrainzEventResponse, error) {
	apiUrl, err := buildEventSearchQuery(opts)
	if err != nil {
		return MusicBrainzEventResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzEventResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchInstrument(opts MusicBrainzInstrumentParameters) (MusicBrainzInstrumentResponse, error) {
	apiUrl, err := buildInstrumentSearchQuery(opts)
	if err != nil {
		return MusicBrainzInstrumentResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzInstrumentResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchLabel(opts MusicBrainzLabelParameters) (MusicBrainzLabelResponse, error) {
	apiUrl, err := buildLabelSearchQuery(opts)
	if err != nil {
		return MusicBrainzLabelResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzLabelResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchPlace(opts MusicBrainzPlaceParameters) (MusicBrainzPlaceResponse, error) {
	apiUrl, err := buildPlaceSearchQuery(opts)
	if err != nil {
		return MusicBrainzPlaceResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzPlaceResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchRecording(opts MusicBrainzRecordingParameters) (MusicBrainzRecordingResponse, error) {
	apiUrl, err := buildRecordingSearchQuery(opts)
	if err != nil {
		return MusicBrainzRecordingResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzRecordingResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchReleaseGroup(opts MusicBrainzReleaseGroupParameters) (MusicBrainzReleaseGroupResponse, error) {
	apiUrl, err := buildReleaseGroupSearchQuery(opts)
	if err != nil {
		return MusicBrainzReleaseGroupResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzReleaseGroupResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchRelease(opts MusicBrainzReleaseParameters) (MusicBrainzReleaseResponse, error) {
	apiUrl, err := buildReleaseSearchQuery(opts)
	if err != nil {
		return MusicBrainzReleaseResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzReleaseResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchSeries(opts MusicBrainzSeriesParameters) (MusicBrainzSeriesResponse, error) {
	apiUrl, err := buildSeriesSearchQuery(opts)
	if err != nil {
		return MusicBrainzSeriesResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzSeriesResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchTag(opts MusicBrainzTagParameters) (MusicBrainzTagResponse, error) {
	apiUrl, err := buildTagSearchQuery(opts)
	if err != nil {
		return MusicBrainzTagResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzTagResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchUrl(opts MusicBrainzUrlParameters) (MusicBrainzUrlResponse, error) {
	apiUrl, err := buildUrlSearchQuery(opts)
	if err != nil {
		return MusicBrainzUrlResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzUrlResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}

func (c *MusicBrainzClient) SearchWork(opts MusicBrainzWorkParameters) (MusicBrainzWorkResponse, error) {
	apiUrl, err := buildWorkSearchQuery(opts)
	if err != nil {
		return MusicBrainzWorkResponse{}, fmt.Errorf("building query: %w", err)
	}

	var result MusicBrainzWorkResponse
	err = c.searchEntity(apiUrl, &result)
	return result, err
}
