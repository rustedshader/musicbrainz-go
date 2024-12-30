# MusicBrainz-Go

**MusicBrainz-Go** is a Go library for interacting with the [MusicBrainz API](https://musicbrainz.org/doc/MusicBrainz_API). This package provides an easy-to-use client for querying the MusicBrainz database to fetch metadata related to music recordings, artists, releases, and more.

## Features

- Query MusicBrainz API for:
  - Artists
  - Releases
  - Recordings
  - Labels
  - Areas
  - Events
  - Instruments
  - Tags
  - Places
  - Series
  - Work
  - URLs
- Simple and efficient HTTP client implementation.
- Handles API response parsing and error management.

---

## Installation

To use this library in your project, install it using `go get`:

```bash
go get github.com/rustedshader/musicbrainz-go
```

## Usage

### Initialize the Client
To start using the library, create a new MusicBrainzClient instance:
```go
package main

import (
	"fmt"
	"musicbrainzgo"
)

func main() {
	client := musicbrainzgo.Client()
	fmt.Println("Client initialized:", client)
}
```

### Example: Search for an Artist
```go
package main

import (
	"fmt"
	"musicbrainzgo"
)

func main() {
	client := musicbrainzgo.Client()

	// Define search parameters for the artist
	opts := musicbrainzgo.MusicBrainzArtistParameters{
		Query: "The Beatles",
	}

	// Search for the artist
	response, err := client.SearchArtist(opts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print results
	fmt.Println("Artist Search Results:", response)
}
```

## API Reference

### Search Methods

The following search methods are available in the client:

```go
// Search for annotations
SearchAnnotation(opts MusicBrainzAnnotationParameters)

// Search for areas
SearchArea(opts MusicBrainzAreaParameters)

// Search for artists
SearchArtist(opts MusicBrainzArtistParameters)

// Search for CD stubs
SearchCDStub(opts MusicBrainzCDStubsParameters)

// Search for events
SearchEvent(opts MusicBrainzEventParameters)

// Search for instruments
SearchInstrument(opts MusicBrainzInstrumentParameters)

// Search for labels
SearchLabel(opts MusicBrainzLabelParameters)

// Search for places
SearchPlace(opts MusicBrainzPlaceParameters)

// Search for recordings
SearchRecording(opts MusicBrainzRecordingParameters)

// Search for release groups
SearchReleaseGroup(opts MusicBrainzReleaseGroupParameters)

// Search for releases
SearchRelease(opts MusicBrainzReleaseParameters)

// Search for series
SearchSeries(opts MusicBrainzSeriesParameters)

// Search for tags
SearchTag(opts MusicBrainzTagParameters)

// Search for URLs
SearchUrl(opts MusicBrainzUrlParameters)

// Search for works
SearchWork(opts MusicBrainzWorkParameters)
```

## Author

- **Shubhang Sharma** - [GitHub](https://github.com/rustedshader)
