package parser

import "time"

type AsAppsModel struct {
	StorePlatformData struct {
		NativeSearchLockup struct {
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
		} `json:"native-search-lockup"`
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
				SearchTerm string `json:"searchTerm"`
			} `json:"fields"`
		} `json:"metrics"`
		PageType string `json:"pageType"`
		Term     string `json:"term"`
		Bubbles  []struct {
			Results []struct {
				Type   int    `json:"type"`
				ID     string `json:"id"`
				Entity string `json:"entity"`
			} `json:"results"`
			Name       string `json:"name"`
			TotalCount int    `json:"totalCount"`
		} `json:"bubbles"`
		Sf6ResourceImagePath string `json:"sf6ResourceImagePath"`
	} `json:"pageData"`
	Properties struct {
		RevNum    string    `json:"revNum"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"properties"`
}
