package musicbrainzgo

type MusicBrainzAnnotationParameters struct {
	Entity string
	Id     int
	Name   string
	Text   string
	Type   string
}

type MusicBrainzAreaParameters struct {
	Aid        string
	Alias      string
	Area       string
	Areaaccent string
	Begin      string
	Comment    string
	End        string
	Ended      string
	Iso        string
	Iso1       string
	Iso2       string
	Iso3       string
	SortName   string
	Tag        string
	Type       string
}
type MusicBrainzArtistParameters struct {
	Alias        string
	PrimaryAlias string
	Area         string
	Arid         string
	Artist       string
	ArtistAccent string
	Begin        string
	BeginArea    string
	Comment      string
	Country      string
	End          string
	EndArea      string
	Ended        string
	Gender       string
	IPI          string
	ISNI         string
	SortName     string
	Tag          string
	Type         string
}

type MusicBrainzCDStubsParameters struct {
	Added   string
	Artist  string
	Barcode string
	Comment string
	DiscID  string
	Title   string
	Tracks  string
}

type MusicBrainzEventParameters struct {
	Alias       string
	Aid         string
	Area        string
	Arid        string
	Artist      string
	Begin       string
	Comment     string
	End         string
	Ended       string
	Eid         string
	Event       string
	EventAccent string
	Pid         string
	Place       string
	Tag         string
	Type        string
}

type MusicBrainzInstrumentParameters struct {
	Alias            string
	Comment          string
	Description      string
	Iid              string
	Instrument       string
	InstrumentAccent string
	Tag              string
	Type             string
}

type MusicBrainzLabelParameters struct {
	Alias        string
	Area         string
	Begin        string
	Code         string
	Comment      string
	Country      string
	End          string
	Ended        string
	IPI          string
	ISNI         string
	Label        string
	LabelAccent  string
	Laid         string
	ReleaseCount int
	SortName     string
	Tag          string
	Type         string
}

type MusicBrainzPlaceParameters struct {
	Address     string
	Alias       string
	Area        string
	Begin       string
	Comment     string
	End         string
	Ended       string
	Latitude    string
	Longitude   string
	Place       string
	PlaceAccent string
	Pid         string
	Type        string
}

type MusicBrainzRecordingParameters struct {
	Alias            string
	Arid             string
	Artist           string
	ArtistName       string
	Comment          string
	Country          string
	CreditName       string
	Date             string
	Duration         int
	FirstReleaseDate string
	Format           string
	ISRC             string
	Number           string
	Position         int
	PrimaryType      string
	QuantizedDur     int
	Recording        string
	RecordingAccent  string
	Reid             string
	Release          string
	Rgid             string
	Rid              string
	SecondaryType    string
	Status           string
	Tag              string
	Tid              string
	TrackNum         int
	Tracks           int
	TracksRelease    int
	Type             string
	Video            bool
}

type MusicBrainzReleaseGroupParameters struct {
	Alias              string
	Arid               string
	Artist             string
	ArtistName         string
	Comment            string
	CreditName         string
	FirstReleaseDate   string
	PrimaryType        string
	Reid               string
	Release            string
	ReleaseGroup       string
	ReleaseGroupAccent string
	Releases           int
	Rgid               string
	SecondaryType      string
	Status             string
	Tag                string
	Type               string
}

type MusicBrainzReleaseParameters struct {
	Alias         string
	Arid          string
	Artist        string
	ArtistName    string
	ASIN          string
	Barcode       string
	CatNo         string
	Comment       string
	Country       string
	CreditName    string
	Date          string
	DiscIDs       int
	DiscIDsMedium int
	Format        string
	Laid          string
	Label         string
	Lang          string
	Mediums       int
	Packaging     string
	PrimaryType   string
	Quality       int
	Reid          string
	Release       string
	ReleaseAccent string
	Rgid          string
	Script        string
	SecondaryType string
	Status        string
	Tag           string
	Tracks        int
	TracksMedium  int
	Type          string
}

type MusicBrainzSeriesParameters struct {
	Alias        string
	Comment      string
	Series       string
	SeriesAccent string
	Sid          string
	Tag          string
	Type         string
}

type MusicBrainzTagParameters struct {
	Tag string
}

type MusicBrainzUrlParameters struct {
	RelationType  string
	TargetId      string
	TargetType    string
	Uid           string
	Url           string
	UrlAncestor   string
	UrlDescendent string
}

type MusicBrainzWorkParameters struct {
	Alias          string
	Arid           string
	Artist         string
	Comment        string
	ISWC           string
	Lang           string
	Recording      string
	RecordingCount int
	Rid            string
	Tag            string
	Type           string
	Wid            string
	Work           string
	WorkAccent     string
}

type MusicBrainzAnnotationResponse struct {
	Created     string `json:"created"`
	Count       int    `json:"count"`
	Offset      int    `json:"offset"`
	Annotations []struct {
		Type   string `json:"type"`
		Score  int    `json:"score"`
		Entity string `json:"entity"`
		Name   string `json:"name"`
		Text   string `json:"text"`
	} `json:"annotations"`
}

type MusicBrainzAreaResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	Areas   []struct {
		ID            string   `json:"id"`
		Type          string   `json:"type"`
		Score         int      `json:"score"`
		Name          string   `json:"name"`
		SortName      string   `json:"sort-name"`
		ISO31662Codes []string `json:"iso-3166-2-codes"`
		LifeSpan      struct {
			Ended interface{} `json:"ended"`
		} `json:"life-span"`
		Aliases []struct {
			SortName  string      `json:"sort-name"`
			Name      string      `json:"name"`
			Locale    string      `json:"locale"`
			Type      string      `json:"type"`
			Primary   bool        `json:"primary"`
			BeginDate interface{} `json:"begin-date"`
			EndDate   interface{} `json:"end-date"`
		} `json:"aliases"`
		RelationList []struct {
			Relations []struct {
				Type      string `json:"type"`
				TypeID    string `json:"type-id"`
				Target    string `json:"target"`
				Direction string `json:"direction"`
				Area      struct {
					ID       string `json:"id"`
					Type     string `json:"type"`
					Name     string `json:"name"`
					SortName string `json:"sort-name"`
					LifeSpan struct {
						Ended interface{} `json:"ended"`
					} `json:"life-span"`
				} `json:"area"`
			} `json:"relations"`
		} `json:"relation-list"`
	} `json:"areas"`
}

type MusicBrainzArtistResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	Artists []struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		Score    int    `json:"score"`
		Name     string `json:"name"`
		SortName string `json:"sort-name"`
		Country  string `json:"country"`
		Area     struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			SortName string `json:"sort-name"`
		} `json:"area"`
		BeginArea struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			SortName string `json:"sort-name"`
		} `json:"begin-area,omitempty"`
		Disambiguation string `json:"disambiguation,omitempty"`
		LifeSpan       struct {
			Begin string      `json:"begin,omitempty"`
			End   string      `json:"end,omitempty"`
			Ended interface{} `json:"ended"`
		} `json:"life-span"`
		Aliases []struct {
			SortName  string      `json:"sort-name"`
			Name      string      `json:"name"`
			Locale    interface{} `json:"locale"`
			Type      interface{} `json:"type"`
			Primary   interface{} `json:"primary"`
			BeginDate interface{} `json:"begin-date"`
			EndDate   interface{} `json:"end-date"`
		} `json:"aliases,omitempty"`
		Tags []struct {
			Count int    `json:"count"`
			Name  string `json:"name"`
		} `json:"tags,omitempty"`
	} `json:"artists"`
}

type MusicBrainzCDStubsResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	CDStubs []struct {
		ID      string `json:"id"`
		Score   int    `json:"score"`
		Count   int    `json:"count"`
		Title   string `json:"title"`
		Artist  string `json:"artist"`
		Barcode string `json:"barcode"`
		Comment string `json:"comment"`
	} `json:"cdstubs"`
}

type MusicBrainzEventResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	Events  []struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		Score    int    `json:"score"`
		Name     string `json:"name"`
		LifeSpan struct {
			Begin string `json:"begin"`
			End   string `json:"end"`
		} `json:"life-span"`
		Time      string `json:"time"`
		Relations []struct {
			Type      string `json:"type"`
			Direction string `json:"direction"`
			Artist    struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				SortName string `json:"sort-name"`
			} `json:"artist,omitempty"`
			Place struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"place,omitempty"`
		} `json:"relations"`
	} `json:"events"`
}

type MusicBrainzInstrumentResponse struct {
	Created     string `json:"created"`
	Count       int    `json:"count"`
	Offset      int    `json:"offset"`
	Instruments []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		Score       int    `json:"score"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Aliases     []struct {
			SortName  string      `json:"sort-name"`
			Name      string      `json:"name"`
			Locale    interface{} `json:"locale"`
			Type      interface{} `json:"type"`
			Primary   interface{} `json:"primary"`
			BeginDate interface{} `json:"begin-date"`
			EndDate   interface{} `json:"end-date"`
		} `json:"aliases"`
	} `json:"instruments"`
}

type MusicBrainzLabelResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	Labels  []struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		Score    int    `json:"score"`
		Name     string `json:"name"`
		SortName string `json:"sort-name"`
		Country  string `json:"country"`
		Area     struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			SortName string `json:"sort-name"`
		} `json:"area"`
		LifeSpan struct {
			Ended interface{} `json:"ended"`
		} `json:"life-span"`
		Aliases []struct {
			SortName  string      `json:"sort-name"`
			Name      string      `json:"name"`
			Locale    interface{} `json:"locale"`
			Type      interface{} `json:"type"`
			Primary   interface{} `json:"primary"`
			BeginDate interface{} `json:"begin-date"`
			EndDate   interface{} `json:"end-date"`
		} `json:"aliases"`
	} `json:"labels"`
}

type MusicBrainzPlaceResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	Places  []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		Score       int    `json:"score"`
		Name        string `json:"name"`
		Address     string `json:"address"`
		Coordinates struct {
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"coordinates"`
		Area struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			SortName string `json:"sort-name"`
		} `json:"area"`
		LifeSpan struct {
			Begin string `json:"begin"`
			End   string `json:"end"`
			Ended bool   `json:"ended"`
		} `json:"life-span"`
	} `json:"places"`
}

type MusicBrainzRecordingResponse struct {
	Created    string `json:"created"`
	Count      int    `json:"count"`
	Offset     int    `json:"offset"`
	Recordings []struct {
		ID           string      `json:"id"`
		Score        int         `json:"score"`
		Title        string      `json:"title"`
		Length       int         `json:"length"`
		Video        interface{} `json:"video"` // Might Have to fix something here
		ArtistCredit []struct {
			Artist struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				SortName string `json:"sort-name"`
			} `json:"artist"`
		} `json:"artist-credit"`
		FirstReleaseDate string `json:"first-release-date"`
		Releases         []struct {
			ID           string `json:"id"`
			Title        string `json:"title"`
			StatusID     string `json:"status-id"`
			Status       string `json:"status"`
			ReleaseGroup struct {
				ID             string   `json:"id"`
				PrimaryType    string   `json:"primary-type"`
				SecondaryTypes []string `json:"secondary-types,omitempty"`
			} `json:"release-group"`
			Date          string `json:"date"`
			Country       string `json:"country"`
			ReleaseEvents []struct {
				Date string `json:"date"`
				Area struct {
					ID            string   `json:"id"`
					Name          string   `json:"name"`
					SortName      string   `json:"sort-name"`
					ISO31661Codes []string `json:"iso-3166-1-codes"`
				} `json:"area"`
			} `json:"release-events"`
			TrackCount int `json:"track-count"`
			Media      []struct {
				Position    int    `json:"position"`
				Format      string `json:"format"`
				TrackCount  int    `json:"track-count"`
				TrackOffset int    `json:"track-offset"`
				Track       []struct {
					ID     string `json:"id"`
					Number string `json:"number"`
					Title  string `json:"title"`
					Length int    `json:"length"`
				} `json:"track"`
			} `json:"media"`
		} `json:"releases"`
		ISRCs []struct {
			ID string `json:"id"`
		} `json:"isrcs"`
		Tags []struct {
			Count int    `json:"count"`
			Name  string `json:"name"`
		} `json:"tags"`
	} `json:"recordings"`
}

type MusicBrainzReleaseGroupResponse struct {
	Created       string `json:"created"`
	Count         int    `json:"count"`
	Offset        int    `json:"offset"`
	ReleaseGroups []struct {
		ID               string `json:"id"`
		Score            int    `json:"score"`
		Count            int    `json:"count"`
		Title            string `json:"title"`
		FirstReleaseDate string `json:"first-release-date"`
		PrimaryType      string `json:"primary-type"`
		ArtistCredit     []struct {
			Artist struct {
				ID             string `json:"id"`
				Name           string `json:"name"`
				SortName       string `json:"sort-name"`
				Disambiguation string `json:"disambiguation"`
				Aliases        []struct {
					SortName  string      `json:"sort-name"`
					Name      string      `json:"name"`
					Locale    interface{} `json:"locale"`
					Type      interface{} `json:"type"`
					Primary   interface{} `json:"primary"`
					BeginDate interface{} `json:"begin-date"`
					EndDate   interface{} `json:"end-date"`
				} `json:"aliases"`
			} `json:"artist"`
		} `json:"artist-credit"`
		Releases []struct {
			ID     string `json:"id"`
			Title  string `json:"title"`
			Status string `json:"status"`
		} `json:"releases"`
	} `json:"release-groups"`
}

type MusicBrainzReleaseResponse struct {
	Created  string `json:"created"`
	Count    int    `json:"count"`
	Offset   int    `json:"offset"`
	Releases []struct {
		ID                 string `json:"id"`
		Score              int    `json:"score"`
		Count              int    `json:"count"`
		Title              string `json:"title"`
		StatusID           string `json:"status-id"`
		Status             string `json:"status"`
		Packaging          string `json:"packaging"`
		TextRepresentation struct {
			Language string `json:"language"`
			Script   string `json:"script"`
		} `json:"text-representation"`
		ArtistCredit []struct {
			Artist struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				SortName string `json:"sort-name"`
				Aliases  []struct {
					SortName  string      `json:"sort-name"`
					Name      string      `json:"name"`
					Locale    interface{} `json:"locale"`
					Type      interface{} `json:"type"`
					Primary   interface{} `json:"primary"`
					BeginDate interface{} `json:"begin-date"`
					EndDate   interface{} `json:"end-date"`
				} `json:"aliases"`
			} `json:"artist"`
		} `json:"artist-credit"`
		ReleaseGroup struct {
			ID          string `json:"id"`
			PrimaryType string `json:"primary-type"`
		} `json:"release-group"`
		Date          string `json:"date"`
		Country       string `json:"country"`
		ReleaseEvents []struct {
			Date string `json:"date"`
			Area struct {
				ID            string   `json:"id"`
				Name          string   `json:"name"`
				SortName      string   `json:"sort-name"`
				ISO31661Codes []string `json:"iso-3166-1-codes"`
			} `json:"area"`
		} `json:"release-events"`
		Barcode   string `json:"barcode"`
		LabelInfo []struct {
			CatalogNumber string `json:"catalog-number"`
			Label         struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"label"`
		} `json:"label-info"`
		TrackCount int `json:"track-count"`
		Media      []struct {
			Format     string `json:"format"`
			DiscCount  int    `json:"disc-count"`
			TrackCount int    `json:"track-count"`
		} `json:"media"`
	} `json:"releases"`
}

type MusicBrainzSeriesResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	Series  []struct {
		ID             string `json:"id"`
		Type           string `json:"type"`
		Score          int    `json:"score"`
		Name           string `json:"name"`
		Disambiguation string `json:"disambiguation"`
		Tags           []struct {
			Count int    `json:"count"`
			Name  string `json:"name"`
		} `json:"tags,omitempty"`
	} `json:"series"`
}

type MusicBrainzTagResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	Tags    []struct {
		Score int    `json:"score"`
		Name  string `json:"name"`
	} `json:"tags"`
}

type MusicBrainzUrlResponse struct {
	Created string     `json:"created"`
	Count   int        `json:"count"`
	Offset  int        `json:"offset"`
	Urls    []struct{} `json:"urls"`
}

type MusicBrainzWorkResponse struct {
	Created string `json:"created"`
	Count   int    `json:"count"`
	Offset  int    `json:"offset"`
	Works   []struct {
		ID        string `json:"id"`
		Score     int    `json:"score"`
		Title     string `json:"title"`
		Relations []struct {
			Type      string `json:"type"`
			Direction string `json:"direction"`
			Artist    struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				SortName string `json:"sort-name"`
			} `json:"artist,omitempty"`
			Recording struct {
				ID    string      `json:"id"`
				Title string      `json:"title"`
				Video interface{} `json:"video"`
			} `json:"recording,omitempty"`
		} `json:"relations"`
	} `json:"works"`
}
