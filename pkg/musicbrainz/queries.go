package musicbrainzgo

import "strconv"

func buildAnnotationSearchQuery(opts MusicBrainzAnnotationParameters) (string, error) {
	query := NewQueryBuilder("annotation").
		AddParam("entity", opts.Entity).
		AddNumericParam("id", opts.Id).
		AddParam("name", opts.Name).
		AddParam("text", opts.Text).
		AddParam("type", opts.Type)
	return query.Build()
}

func buildAreaSearchQuery(opts MusicBrainzAreaParameters) (string, error) {
	query := NewQueryBuilder("area").
		AddParam("area", opts.Area).
		AddParam("areaaccent", opts.Areaaccent).
		AddParam("begin", opts.Begin).
		AddParam("comment", opts.Comment).
		AddParam("end", opts.End).
		AddParam("ended", opts.Ended).
		AddParam("iso", opts.Iso).
		AddParam("iso1", opts.Iso1).
		AddParam("iso2", opts.Iso2).
		AddParam("iso3", opts.Iso3).
		AddParam("sortname", opts.SortName).
		AddParam("tag", opts.Tag).
		AddParam("type", opts.Type)
	return query.Build()
}
func buildArtistSearchQuery(opts MusicBrainzArtistParameters) (string, error) {
	query := NewQueryBuilder("artist").
		AddParam("alias", opts.Alias).
		AddParam("primary_alias", opts.PrimaryAlias).
		AddParam("area", opts.Area).
		AddParam("arid", opts.Arid).
		AddParam("artist", opts.Artist).
		AddParam("artistaccent", opts.ArtistAccent).
		AddParam("begin", opts.Begin).
		AddParam("beginarea", opts.BeginArea).
		AddParam("comment", opts.Comment).
		AddParam("country", opts.Country).
		AddParam("end", opts.End).
		AddParam("endarea", opts.EndArea).
		AddParam("ended", opts.Ended).
		AddParam("gender", opts.Gender).
		AddParam("ipi", opts.IPI).
		AddParam("isni", opts.ISNI).
		AddParam("sortname", opts.SortName).
		AddParam("tag", opts.Tag).
		AddParam("type", opts.Type)
	return query.Build()
}
func buildCDStubSearchQuery(opts MusicBrainzCDStubsParameters) (string, error) {
	query := NewQueryBuilder("cdstub").
		AddParam("added", opts.Added).
		AddParam("artist", opts.Artist).
		AddParam("barcode", opts.Barcode).
		AddParam("comment", opts.Comment).
		AddParam("discid", opts.DiscID).
		AddParam("title", opts.Title).
		AddParam("tracks", opts.Tracks)
	return query.Build()
}

func buildEventSearchQuery(opts MusicBrainzEventParameters) (string, error) {
	query := NewQueryBuilder("event").
		AddParam("alias", opts.Alias).
		AddParam("aid", opts.Aid).
		AddParam("area", opts.Area).
		AddParam("arid", opts.Arid).
		AddParam("artist", opts.Artist).
		AddParam("begin", opts.Begin).
		AddParam("comment", opts.Comment).
		AddParam("end", opts.End).
		AddParam("ended", opts.Ended).
		AddParam("eid", opts.Eid).
		AddParam("event", opts.Event).
		AddParam("eventaccent", opts.EventAccent).
		AddParam("pid", opts.Pid).
		AddParam("place", opts.Place).
		AddParam("tag", opts.Tag).
		AddParam("type", opts.Type)
	return query.Build()
}

func buildInstrumentSearchQuery(opts MusicBrainzInstrumentParameters) (string, error) {
	query := NewQueryBuilder("instrument").
		AddParam("alias", opts.Alias).
		AddParam("comment", opts.Comment).
		AddParam("description", opts.Description).
		AddParam("iid", opts.Iid).
		AddParam("instrument", opts.Instrument).
		AddParam("instrumentaccent", opts.InstrumentAccent).
		AddParam("tag", opts.Tag).
		AddParam("type", opts.Type)
	return query.Build()
}

func buildLabelSearchQuery(opts MusicBrainzLabelParameters) (string, error) {
	query := NewQueryBuilder("label").
		AddParam("alias", opts.Alias).
		AddParam("area", opts.Area).
		AddParam("begin", opts.Begin).
		AddParam("code", opts.Code).
		AddParam("comment", opts.Comment).
		AddParam("country", opts.Country).
		AddParam("end", opts.End).
		AddParam("ended", opts.Ended).
		AddParam("ipi", opts.IPI).
		AddParam("isni", opts.ISNI).
		AddParam("label", opts.Label).
		AddParam("labelaccent", opts.LabelAccent).
		AddParam("laid", opts.Laid).
		AddNumericParam("releasecount", opts.ReleaseCount).
		AddParam("sortname", opts.SortName).
		AddParam("tag", opts.Tag).
		AddParam("type", opts.Type)
	return query.Build()
}

func buildPlaceSearchQuery(opts MusicBrainzPlaceParameters) (string, error) {
	query := NewQueryBuilder("place").
		AddParam("address", opts.Address).
		AddParam("alias", opts.Alias).
		AddParam("area", opts.Area).
		AddParam("begin", opts.Begin).
		AddParam("comment", opts.Comment).
		AddParam("end", opts.End).
		AddParam("ended", opts.Ended).
		AddParam("lat", opts.Latitude).
		AddParam("long", opts.Longitude).
		AddParam("place", opts.Place).
		AddParam("placeaccent", opts.PlaceAccent).
		AddParam("pid", opts.Pid).
		AddParam("type", opts.Type)
	return query.Build()
}

func buildRecordingSearchQuery(opts MusicBrainzRecordingParameters) (string, error) {
	query := NewQueryBuilder("recording").AddParam("alias", opts.Alias).
		AddParam("arid", opts.Arid).
		AddParam("artist", opts.Artist).
		AddParam("artistname", opts.ArtistName).
		AddParam("comment", opts.Comment).
		AddParam("country", opts.Country).
		AddParam("creditname", opts.CreditName).
		AddParam("date", opts.Date).
		AddParam("firstReleaseDate", opts.FirstReleaseDate).
		AddParam("format", opts.Format).
		AddParam("isrc", opts.ISRC).
		AddParam("number", opts.Number).
		AddNumericParam("position", opts.Position).
		AddParam("primarytype", opts.PrimaryType).
		AddNumericParam("quantizeddur", opts.QuantizedDur).
		AddParam("recording", opts.Recording).
		AddParam("recordingaccent", opts.RecordingAccent).
		AddParam("reid", opts.Reid).
		AddParam("release", opts.Release).
		AddParam("rgid", opts.Rgid).
		AddParam("rid", opts.Rid).
		AddParam("secondarytype", opts.SecondaryType).
		AddParam("status", opts.Status).
		AddParam("tag", opts.Tag).
		AddParam("tid", opts.Tid).
		AddNumericParam("tracknumber", opts.TrackNum).
		AddNumericParam("tracks", opts.Tracks).
		AddNumericParam("tracksrelease", opts.TracksRelease).
		AddParam("type", opts.Type).
		AddParam("video", strconv.FormatBool(opts.Video))
	if opts.Duration > 0 {
		query.AddRangeParam("dur", opts.Duration-500, opts.Duration+500)
	}
	return query.Build()
}

func buildReleaseGroupSearchQuery(opts MusicBrainzReleaseGroupParameters) (string, error) {
	query := NewQueryBuilder("release-group").
		AddParam("alias", opts.Alias).
		AddParam("arid", opts.Arid).
		AddParam("artist", opts.Artist).
		AddParam("artistname", opts.ArtistName).
		AddParam("comment", opts.Comment).
		AddParam("creditname", opts.CreditName).
		AddParam("firstrelease", opts.FirstReleaseDate).
		AddParam("primarytype", opts.PrimaryType).
		AddParam("reid", opts.Reid).
		AddParam("release", opts.Release).
		AddParam("releasegroup", opts.ReleaseGroup).
		AddParam("releasegroupaccent", opts.ReleaseGroupAccent).
		AddNumericParam("releases", opts.Releases).
		AddParam("rgid", opts.Rgid).
		AddParam("secondarytype", opts.SecondaryType).
		AddParam("status", opts.Status).
		AddParam("tag", opts.Tag).
		AddParam("type", opts.Type)
	return query.Build()
}

func buildReleaseSearchQuery(opts MusicBrainzReleaseParameters) (string, error) {
	query := NewQueryBuilder("release").
		AddParam("alias", opts.Alias).
		AddParam("arid", opts.Arid).
		AddParam("artist", opts.Artist).
		AddParam("artistname", opts.ArtistName).
		AddParam("asin", opts.ASIN).
		AddParam("barcode", opts.Barcode).
		AddParam("catno", opts.CatNo).
		AddParam("comment", opts.Comment).
		AddParam("country", opts.Country).
		AddParam("creditname", opts.CreditName).
		AddParam("date", opts.Date).
		AddNumericParam("discids", opts.DiscIDs).
		AddNumericParam("discidsmedium", opts.DiscIDsMedium).
		AddParam("format", opts.Format).
		AddParam("laid", opts.Laid).
		AddParam("label", opts.Label).
		AddParam("lang", opts.Lang).
		AddNumericParam("mediums", opts.Mediums).
		AddParam("packaging", opts.Packaging).
		AddParam("primarytype", opts.PrimaryType).
		AddNumericParam("quality", opts.Quality).
		AddParam("reid", opts.Reid).
		AddParam("release", opts.Release).
		AddParam("releaseaccent", opts.ReleaseAccent).
		AddParam("rgid", opts.Rgid).
		AddParam("script", opts.Script).
		AddParam("secondarytype", opts.SecondaryType).
		AddParam("status", opts.Status).
		AddParam("tag", opts.Tag).
		AddNumericParam("tracks", opts.Tracks).
		AddNumericParam("tracksmedium", opts.TracksMedium).
		AddParam("type", opts.Type)
	return query.Build()
}

func buildSeriesSearchQuery(opts MusicBrainzSeriesParameters) (string, error) {
	query := NewQueryBuilder("series").
		AddParam("alias", opts.Alias).
		AddParam("comment", opts.Comment).
		AddParam("series", opts.Series).
		AddParam("seriesaccent", opts.SeriesAccent).
		AddParam("sid", opts.Sid).
		AddParam("tag", opts.Tag).
		AddParam("type", opts.Type)
	return query.Build()
}

func buildTagSearchQuery(opts MusicBrainzTagParameters) (string, error) {
	query := NewQueryBuilder("tag").
		AddParam("tag", opts.Tag)
	return query.Build()
}

func buildUrlSearchQuery(opts MusicBrainzUrlParameters) (string, error) {
	query := NewQueryBuilder("url").
		AddParam("relationtype", opts.RelationType).
		AddParam("targetid", opts.TargetId).
		AddParam("targettype", opts.TargetType).
		AddParam("uid", opts.Uid).
		AddParam("url", opts.Url).
		AddParam("urlancestor", opts.UrlAncestor).
		AddParam("urldescendent", opts.UrlDescendent)
	return query.Build()
}

func buildWorkSearchQuery(opts MusicBrainzWorkParameters) (string, error) {
	query := NewQueryBuilder("work").
		AddParam("alias", opts.Alias).
		AddParam("arid", opts.Arid).
		AddParam("artist", opts.Artist).
		AddParam("comment", opts.Comment).
		AddParam("iswc", opts.ISWC).
		AddParam("lang", opts.Lang).
		AddParam("recording", opts.Recording).
		AddNumericParam("recordingcount", opts.RecordingCount).
		AddParam("rid", opts.Rid).
		AddParam("tag", opts.Tag).
		AddParam("type", opts.Type).
		AddParam("wid", opts.Wid).
		AddParam("work", opts.Work).
		AddParam("workaccent", opts.WorkAccent)
	return query.Build()
}
