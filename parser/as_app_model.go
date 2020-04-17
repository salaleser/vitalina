package parser

import "time"

type AsAppModel struct {
	StorePlatformData struct {
		ProductDv struct {
			Results         interface{} `json:"results"`
			Version         int         `json:"version"`
			IsAuthenticated bool        `json:"isAuthenticated"`
			Meta            struct {
				Storefront struct {
					ID string `json:"id"`
					Cc string `json:"cc"`
				} `json:"storefront"`
				Language struct {
					Tag string `json:"tag"`
				} `json:"language"`
			} `json:"meta"`
		} `json:"product-dv"`
	} `json:"storePlatformData"`
	PageData struct {
		ComponentName string `json:"componentName"`
		MetricsBase   struct {
			PageType              string `json:"pageType"`
			PageID                string `json:"pageId"`
			PageDetails           string `json:"pageDetails"`
			Page                  string `json:"page"`
			ServerInstance        string `json:"serverInstance"`
			StoreFrontHeader      string `json:"storeFrontHeader"`
			Language              string `json:"language"`
			PlatformID            string `json:"platformId"`
			PlatformName          string `json:"platformName"`
			StoreFront            string `json:"storeFront"`
			EnvironmentDataCenter string `json:"environmentDataCenter"`
		} `json:"metricsBase"`
		Metrics struct {
			Config struct {
			} `json:"config"`
			Fields struct {
			} `json:"fields"`
		} `json:"metrics"`
		TopApps struct {
			Iphone struct {
				Ids   []string `json:"ids"`
				Title string   `json:"title"`
			} `json:"iphone"`
		} `json:"topApps"`
		RatingAndAdvisories struct {
			RatingText string `json:"rating-text"`
		} `json:"rating-and-advisories"`
		KindExtID              string `json:"kindExtId"`
		UserReviewsSortOptions []struct {
			SortID int    `json:"sortId"`
			Name   string `json:"name"`
		} `json:"userReviewsSortOptions"`
		CustomerReviewsURL string `json:"customerReviewsUrl"`
		IsFatBinary        int    `json:"isFatBinary"`
		UserReviewList     []struct {
			UserReviewID             string    `json:"userReviewId"`
			Body                     string    `json:"body"`
			Date                     time.Time `json:"date"`
			Name                     string    `json:"name"`
			Rating                   int       `json:"rating"`
			Title                    string    `json:"title"`
			VoteCount                int       `json:"voteCount"`
			VoteSum                  int       `json:"voteSum"`
			IsEdited                 bool      `json:"isEdited"`
			ViewUsersUserReviewsURL  string    `json:"viewUsersUserReviewsUrl"`
			VoteURL                  string    `json:"voteUrl"`
			ReportConcernURL         string    `json:"reportConcernUrl"`
			ReportConcernExplanation string    `json:"reportConcernExplanation"`
			CustomerType             string    `json:"customerType"`
			DeveloperResponse        struct {
				ID       int       `json:"id"`
				Body     string    `json:"body"`
				Modified time.Time `json:"modified"`
			} `json:"developerResponse"`
			ReportConcernReasons []struct {
				ReasonID      string `json:"reasonId"`
				Name          string `json:"name"`
				UpperCaseName string `json:"upperCaseName"`
			} `json:"reportConcernReasons"`
		} `json:"userReviewList"`
		TotalNumberOfReviews int `json:"totalNumberOfReviews"`
		VersionHistory       []struct {
			ReleaseNotes  string    `json:"releaseNotes"`
			VersionString string    `json:"versionString"`
			ReleaseDate   time.Time `json:"releaseDate"`
		} `json:"versionHistory"`
		CustomersAlsoBoughtApps []string `json:"customersAlsoBoughtApps"`
		KindName                string   `json:"kindName"`
		ID                      string   `json:"id"`
		AppRatingsLearnMoreURL  string   `json:"appRatingsLearnMoreUrl"`
		SellerLabel             string   `json:"sellerLabel"`
		MoreByThisDeveloper     []string `json:"moreByThisDeveloper"`
		KindID                  int      `json:"kindId"`
		Sf6ResourceImagePath    string   `json:"sf6ResourceImagePath"`
	} `json:"pageData"`
	Properties struct {
		RevNum    string    `json:"revNum"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"properties"`
}
