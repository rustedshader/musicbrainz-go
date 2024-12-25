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