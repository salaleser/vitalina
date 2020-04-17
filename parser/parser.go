package parser

import (
	"encoding/json"
	"log"
	"strings"
)

// App contains application
type App struct {
	Title string
	AppID string
}

// Metadata contains application's metadata
type Metadata struct { // TODO add more fields
	Title       string
	Subtitle    string
	Description string
	Screenshot1 string // TODO add array
	Logo        string
}

// ParseAsIDsBody parses applications by keyword
func ParseAsIDsBody(body []byte, count int) []AsResultModel {
	var data AsAppsModel
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal (234): %q", err.Error())
		return []AsResultModel{} // TODO handle error
	}

	results := data.StorePlatformData.NativeSearchLockup.Results.(map[string]interface{})

	apps := make([]AsResultModel, 0)
	for k, v := range results {
		result := v.(map[string]interface{})

		if result["kind"].(string) == "iosSoftware" {
			apps = append(apps, AsResultModel{
				Name: result["name"].(string),
				ID:   k,
			})
		}
	}

	return apps
}

// ParseAsAsoBody parses response body and returns ASO
func ParseAsAsoBody(body []byte) Metadata {
	var data AsAppModel
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error while trying to unmarshal (1): %q", err.Error())
		return Metadata{} // TODO handle error
	}

	results := data.StorePlatformData.ProductDv.Results.(map[string]interface{})

	var metadata Metadata
	for _, v := range results {
		result := v.(map[string]interface{})

		screenshotsByType := result["screenshotsByType"].(map[string]interface{})
		var screenshot1 string
		for _, screenshots := range screenshotsByType {
			screenshot1 = strings.Replace(screenshots.([]interface{})[0].(map[string]interface{})["url"].(string), "{w}x{h}bb.{f}", "512x512bb.png", -1)
			break
		}

		var subtitle string
		if result["subtitle"] != nil {
			subtitle = result["subtitle"].(string)
		}

		metadata = Metadata{
			Title:       result["name"].(string),
			Subtitle:    subtitle,
			Screenshot1: screenshot1,
			Logo:        strings.Replace(result["artwork"].(map[string]interface{})["url"].(string), "{w}x{h}bb.{f}", "128x128bb.png", -1),
		}
	}

	return metadata
}

// ParseGpIDsBody parses applications by keyword
func ParseGpIDsBody(body []byte, count int) []App {
	var data1 [][]interface{}
	if err := json.Unmarshal(body, &data1); err != nil {
		log.Printf("Error while trying to unmarshal (1): %q", err.Error())
		return []App{} // TODO handle error
	}

	d := data1[0]

	if d[0] != "wrb.fr" {
		log.Printf("The first ASO section element isn't \"wrb.fr\" (%q).", d[0])
		return []App{} // TODO handle error
	}

	if d[1] != "lGYRle" {
		log.Printf("The second ASO section element isn't \"lGYRle\" (%q).", d[0])
		return []App{} // TODO handle error
	}

	if d[2] == nil {
		log.Printf("Error while parsing (386).")
		return []App{} // TODO handle error
	}

	var data2 []interface{}
	if err := json.Unmarshal([]byte(d[2].(string)), &data2); err != nil {
		log.Printf("Error while trying to unmarshal (2): %q", err.Error())
		return []App{} // TODO handle error
	}

	// FIXME
	if len(data2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[0].([]interface{})) < 2 {
		return []App{}
	}

	appIDs := make([]App, count)
	for i := 0; i < count; i++ {
		appIDs[i] = App{
			Title: data2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[0].([]interface{})[i].([]interface{})[2].(string),
			AppID: data2[0].([]interface{})[1].([]interface{})[0].([]interface{})[0].([]interface{})[0].([]interface{})[i].([]interface{})[12].([]interface{})[0].(string),
		}
	}

	return appIDs
}

// ParseGpAsoBody parses response body and returns Metadata
func ParseGpAsoBody(body []byte) Metadata {
	var data1 [][]interface{}
	if err := json.Unmarshal(body, &data1); err != nil {
		log.Printf("Error while trying to unmarshal (1): %q", err.Error())
		return Metadata{} // TODO handle error
	}

	d := data1[0]

	if d[0] != "wrb.fr" {
		log.Printf("The first ASO section element isn't \"wrb.fr\" (%q).", d[0])
		return Metadata{} // TODO handle error
	}

	if d[1] != "jLZZ2e" {
		log.Printf("The second ASO section element isn't \"jLZZ2e\" (%q).", d[0])
		return Metadata{} // TODO handle error
	}

	if d[2] == nil {
		log.Printf("Error while parsing (567).")
		return Metadata{} // TODO handle error
	}

	var data2 [][][]interface{}
	if err := json.Unmarshal([]byte(d[2].(string)), &data2); err != nil {
		log.Printf("Error while trying to unmarshal (2): %q", err.Error())
		return Metadata{} // TODO handle error
	}

	return Metadata{
		Title:       data2[0][0][0].(string),
		Subtitle:    data2[0][10][1].([]interface{})[1].(string),
		Description: data2[0][10][0].([]interface{})[1].(string),
		Screenshot1: data2[0][12][0].([]interface{})[0].([]interface{})[3].([]interface{})[2].(string),
		Logo:        data2[0][12][1].([]interface{})[3].([]interface{})[2].(string),
	}
}
