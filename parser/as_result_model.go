package parser

type AsResultModel struct {
	Artwork struct {
		Width                int    `json:"width"`
		URL                  string `json:"url"`
		Height               int    `json:"height"`
		TextColor3           string `json:"textColor3"`
		TextColor2           string `json:"textColor2"`
		TextColor4           string `json:"textColor4"`
		HasAlpha             bool   `json:"hasAlpha"`
		TextColor1           string `json:"textColor1"`
		BgColor              string `json:"bgColor"`
		HasP3                bool   `json:"hasP3"`
		SupportsLayeredImage bool   `json:"supportsLayeredImage"`
	} `json:"artwork"`
	ArtistName                             string `json:"artistName"`
	FirstVersionSupportingInAppPurchaseAPI string `json:"firstVersionSupportingInAppPurchaseApi"`
	URL                                    string `json:"url"`
	ShortURL                               string `json:"shortUrl"`
	VideoPreviewsByType                    struct {
		Iphone6 []struct {
			PreviewFrame struct {
				Width                int    `json:"width"`
				URL                  string `json:"url"`
				Height               int    `json:"height"`
				TextColor3           string `json:"textColor3"`
				TextColor2           string `json:"textColor2"`
				TextColor4           string `json:"textColor4"`
				HasAlpha             bool   `json:"hasAlpha"`
				TextColor1           string `json:"textColor1"`
				BgColor              string `json:"bgColor"`
				Checksum             string `json:"checksum"`
				HasP3                bool   `json:"hasP3"`
				SupportsLayeredImage bool   `json:"supportsLayeredImage"`
			} `json:"previewFrame"`
			Video string `json:"video"`
		} `json:"iphone6+"`
		Iphone65 []struct {
			PreviewFrame struct {
				Width                int    `json:"width"`
				URL                  string `json:"url"`
				Height               int    `json:"height"`
				TextColor3           string `json:"textColor3"`
				TextColor2           string `json:"textColor2"`
				TextColor4           string `json:"textColor4"`
				HasAlpha             bool   `json:"hasAlpha"`
				TextColor1           string `json:"textColor1"`
				BgColor              string `json:"bgColor"`
				Checksum             string `json:"checksum"`
				HasP3                bool   `json:"hasP3"`
				SupportsLayeredImage bool   `json:"supportsLayeredImage"`
			} `json:"previewFrame"`
			Video string `json:"video"`
		} `json:"iphone_6_5"`
		IpadPro2018 []struct {
			PreviewFrame struct {
				Width                int    `json:"width"`
				URL                  string `json:"url"`
				Height               int    `json:"height"`
				TextColor3           string `json:"textColor3"`
				TextColor2           string `json:"textColor2"`
				TextColor4           string `json:"textColor4"`
				HasAlpha             bool   `json:"hasAlpha"`
				TextColor1           string `json:"textColor1"`
				BgColor              string `json:"bgColor"`
				Checksum             string `json:"checksum"`
				HasP3                bool   `json:"hasP3"`
				SupportsLayeredImage bool   `json:"supportsLayeredImage"`
			} `json:"previewFrame"`
			Video string `json:"video"`
		} `json:"ipadPro_2018"`
		IpadPro []struct {
			PreviewFrame struct {
				Width                int    `json:"width"`
				URL                  string `json:"url"`
				Height               int    `json:"height"`
				TextColor3           string `json:"textColor3"`
				TextColor2           string `json:"textColor2"`
				TextColor4           string `json:"textColor4"`
				HasAlpha             bool   `json:"hasAlpha"`
				TextColor1           string `json:"textColor1"`
				BgColor              string `json:"bgColor"`
				Checksum             string `json:"checksum"`
				HasP3                bool   `json:"hasP3"`
				SupportsLayeredImage bool   `json:"supportsLayeredImage"`
			} `json:"previewFrame"`
			Video string `json:"video"`
		} `json:"ipadPro"`
	} `json:"videoPreviewsByType"`
	DeviceFamilies []string `json:"deviceFamilies"`
	GenreNames     []string `json:"genreNames"`
	NameSortValue  string   `json:"nameSortValue"`
	ID             string   `json:"id"`
	ReleaseDate    string   `json:"releaseDate"`
	UserRating     struct {
		Value               float64 `json:"value"`
		RatingCount         int     `json:"ratingCount"`
		RatingCountList     []int   `json:"ratingCountList"`
		AriaLabelForRatings string  `json:"ariaLabelForRatings"`
	} `json:"userRating"`
	ContentRatingsBySystem struct {
		AppsApple struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
			Rank  int    `json:"rank"`
		} `json:"appsApple"`
	} `json:"contentRatingsBySystem"`
	Name              string      `json:"name"`
	ArtistURL         string      `json:"artistUrl"`
	ScreenshotsByType interface{} `json:"screenshotsByType"`
	NameRaw           string      `json:"nameRaw"`
	EditorialArtwork  struct {
		BrandLogo struct {
			Width                int    `json:"width"`
			URL                  string `json:"url"`
			Height               int    `json:"height"`
			TextColor3           string `json:"textColor3"`
			TextColor2           string `json:"textColor2"`
			TextColor4           string `json:"textColor4"`
			HasAlpha             bool   `json:"hasAlpha"`
			TextColor1           string `json:"textColor1"`
			BgColor              string `json:"bgColor"`
			HasP3                bool   `json:"hasP3"`
			SupportsLayeredImage bool   `json:"supportsLayeredImage"`
		} `json:"brandLogo"`
	} `json:"editorialArtwork"`
	Subtitle              string `json:"subtitle"`
	BundleID              string `json:"bundleId"`
	HasInAppPurchases     bool   `json:"hasInAppPurchases"`
	Kind                  string `json:"kind"`
	Copyright             string `json:"copyright"`
	ChartPositionForStore struct {
		AppStore struct {
			Position  int    `json:"position"`
			GenreName string `json:"genreName"`
			ChartURL  string `json:"chartUrl"`
		} `json:"appStore"`
	} `json:"chartPositionForStore"`
	ArtistID string `json:"artistId"`
	Genres   []struct {
		GenreID       string `json:"genreId"`
		Name          string `json:"name"`
		URL           string `json:"url"`
		MediaType     string `json:"mediaType"`
		ParentGenreID string `json:"parentGenreId"`
	} `json:"genres"`
	MinimumOSVersion    string `json:"minimumOSVersion"`
	MessagesScreenshots struct {
	} `json:"messagesScreenshots"`
	Offers []struct {
		ActionText struct {
			Short       string `json:"short"`
			Medium      string `json:"medium"`
			Long        string `json:"long"`
			Downloaded  string `json:"downloaded"`
			Downloading string `json:"downloading"`
		} `json:"actionText"`
		Type           string `json:"type"`
		PriceFormatted string `json:"priceFormatted"`
		Price          int    `json:"price"`
		BuyParams      string `json:"buyParams"`
		Version        struct {
			Display    string `json:"display"`
			ExternalID int    `json:"externalId"`
		} `json:"version"`
		Assets []struct {
			Flavor string `json:"flavor"`
			Size   int    `json:"size"`
		} `json:"assets"`
	} `json:"offers"`
}
